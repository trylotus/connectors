package near

import (
	"context"
	"embed"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	reflect "reflect"
	sync "sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"github.com/trylotus/connector/kafkautils"
	"google.golang.org/protobuf/proto"
)

//go:embed target/x86_64-unknown-linux-musl/release/lotus_near_client
var f embed.FS

var endpointMap = map[string]proto.Message{
	"/block":   &Block{},
	"/tx":      &Transaction{},
	"/receipt": &Receipt{},
	"/outcome": &ExecutionOutcome{},
}

type Subscription struct {
	hostname             string
	port                 string
	backfillPort         string
	maxRetries           int
	fromBlock            uint64
	numBlocks            uint64
	events               []string
	fctMsgTypes          []proto.Message
	bfMsgTypes           []proto.Message
	done                 chan struct{}
	ctx                  context.Context
	blockChan            chan *kafkautils.Message
	transactionChan      chan *kafkautils.Message
	receiptChan          chan *kafkautils.Message
	executionOutcomeChan chan *kafkautils.Message
	errChan              chan error
}

func NewSubscription(ctx context.Context, config *Config, events []string) (*Subscription, error) {
	// Read binary from embed
	bin, err := f.ReadFile("target/x86_64-unknown-linux-musl/release/lotus_near_client")
	if err != nil {
		log.Fatal().Err(err).Msg("failed read embedded Lotus Near Client binary")
	}

	// Write binary to file so we can execute it
	err = os.WriteFile("/nearclient/lotus_near_client", bin, 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to write embedded Lotus Near Client binary to local fs")
	}

	sub := &Subscription{
		hostname:             config.Hostname,
		port:                 config.Port,
		backfillPort:         config.BackfillPort,
		maxRetries:           config.MaxRetries,
		events:               events,
		fctMsgTypes:          config.FctMsgTypes,
		bfMsgTypes:           config.BfMsgTypes,
		fromBlock:            config.FromBlock,
		numBlocks:            config.NumBlocks,
		done:                 make(chan struct{}),
		ctx:                  ctx,
		blockChan:            make(chan *kafkautils.Message, config.ChannelSize),
		transactionChan:      make(chan *kafkautils.Message, config.ChannelSize),
		receiptChan:          make(chan *kafkautils.Message, config.ChannelSize),
		executionOutcomeChan: make(chan *kafkautils.Message, config.ChannelSize),
		errChan:              make(chan error, config.ChannelSize),
	}
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		select {
		case <-interrupt:
			sub.Unsubscribe()
		case <-ctx.Done():
			sub.Unsubscribe()
		}
	}()

	go sub.subscribe()
	return sub, nil
}

func (s *Subscription) Unsubscribe() {
	log.Info().Str("hostname", s.hostname).Msg("shutting down subscription")
	close(s.done)
}

func (s *Subscription) Done() <-chan struct{} {
	return s.done
}

func (s *Subscription) Blocks() <-chan *kafkautils.Message {
	return s.blockChan
}

func (s *Subscription) Transactions() <-chan *kafkautils.Message {
	return s.transactionChan
}

func (s *Subscription) Receipts() <-chan *kafkautils.Message {
	return s.receiptChan
}

func (s *Subscription) ExecutionOutcomes() <-chan *kafkautils.Message {
	return s.executionOutcomeChan
}

func (s *Subscription) Err() <-chan error {
	return s.errChan
}

func (s *Subscription) subscribe() {
	// Start Lake Stream from current block
	s.startLakeStream(s.port, 0, 0)
	go s.initWsStreams(false)
	// Backfill if needed
	if s.fromBlock > 0 || s.numBlocks > 0 {
		go s.backfill()
	}
}

func (s *Subscription) backfill() {
	log.Info().Msg("starting backfill")
	defer log.Info().Msg("finished backfill")
	if s.fromBlock > 0 {
		s.startLakeStream(s.backfillPort, s.fromBlock, 0)
		s.initWsStreams(true)
	} else if s.numBlocks > 0 {
		s.startLakeStream(s.backfillPort, 0, s.numBlocks)
		s.initWsStreams(true)
	}
}

func (s *Subscription) startLakeStream(port string, fromBlock uint64, numBlocks uint64) {
	// Setup command for Lake stream
	cmd := []string{
		"/nearclient/lotus_near_client",
	}

	cmd = append(cmd, fmt.Sprintf("--port=%s", port))

	// Arguments for binary are detrmined by backfill config
	if fromBlock > 0 {
		cmd = append(cmd, "from-block")
		cmd = append(cmd, fmt.Sprintf("%d", fromBlock))
	} else if numBlocks > 0 {
		cmd = append(cmd, "num-blocks")
		cmd = append(cmd, fmt.Sprintf("%d", numBlocks))
	} else {
		cmd = append(cmd, "from-latest")
	}

	execCmd := exec.Command(cmd[0], cmd[1:]...)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	// Run binary in goroutine
	go func() {
		if err := execCmd.Run(); err != nil {
			log.Fatal().Err(err).Msg(fmt.Sprintf("%s Error running Lotus NEAR Client (Rust binary): ", err))
		}
	}()
}

func (s *Subscription) initWsStreams(isBackfill bool) {
	var wg sync.WaitGroup
	var msgTypes []proto.Message
	// Connect to ws server to receive data from Lake stream
	if isBackfill {
		msgTypes = s.bfMsgTypes
	} else {
		msgTypes = s.fctMsgTypes
	}
	for endpoint, msgType := range endpointMap {
		// Check if msgType is in sub's msgTypes (based on bf or fct msgTypes)
		for _, subType := range msgTypes {
			if reflect.TypeOf(msgType) == reflect.TypeOf(subType) {
				wg.Add(1)
				go s.startWsStream(&wg, endpoint, msgType, isBackfill)
			}
		}
	}
	wg.Wait()
}

func (s *Subscription) startWsStream(wg *sync.WaitGroup, endpoint string, protoMsgType proto.Message, isBackfill bool) {
	defer wg.Done()

	// Get port from config
	var port string
	if isBackfill {
		port = s.backfillPort
	} else {
		port = s.port
	}

	// Construct ws url
	host := fmt.Sprintf("%s%s%s", s.hostname, ":", port)
	u := url.URL{Scheme: "ws", Host: host, Path: endpoint}
	log.Info().Str("host", host).Str("endpoint", endpoint).Msg("connecting to ws server")

	// Output chan determined by protoMsgType
	ch := s.getMsgChan(protoMsgType)
	if ch == nil {
		log.Error().Msg("could not determine output channel")
		return
	}

	// kafkaMsgType determined by backfill
	var kafkaMsgType kafkautils.MsgType
	if isBackfill {
		kafkaMsgType = kafkautils.MsgTypeBf
	} else {
		kafkaMsgType = kafkautils.MsgTypeFct
	}

	// Connect to ws server with retry
	var conn *websocket.Conn
	for retry := 0; ; retry++ {
		tryConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			if retry < s.maxRetries {
				// Do exponential backoff with 10% jitter
				backoff := float64(int(1) << retry)
				backoff += backoff * (0.1 * rand.Float64())
				select {
				case <-s.done:
					return
				case <-time.After(time.Second * time.Duration(backoff)):
					continue
				}
			}
			log.Fatal().Err(err).Msg(fmt.Sprintf("%s Unable to connect to NEAR WS Server", err))
		}
		conn = tryConn
		break
	}

	defer conn.Close()

	// Read messages from connection
	for {
		select {
		case <-s.done:
			return
		default:
			_, wsMessage, err := conn.ReadMessage()
			if err != nil && isBackfill {
				log.Debug().Msg(fmt.Sprintf("%s Backfill ws connection closed", err))
				return
			} else if err != nil {
				s.errChan <- err
				continue
			}

			protoMsg := proto.Clone(protoMsgType)

			if err := proto.Unmarshal(wsMessage, protoMsg); err != nil {
				s.errChan <- err
				continue
			}

			ch <- &kafkautils.Message{
				MsgType:  kafkaMsgType,
				ProtoMsg: protoMsg,
			}

		}
	}
}

func (s *Subscription) getMsgChan(msgType proto.Message) chan *kafkautils.Message {
	switch reflect.TypeOf(msgType) {
	case reflect.TypeOf(&Block{}):
		return s.blockChan
	case reflect.TypeOf(&Transaction{}):
		return s.transactionChan
	case reflect.TypeOf(&Receipt{}):
		return s.receiptChan
	case reflect.TypeOf(&ExecutionOutcome{}):
		return s.executionOutcomeChan
	default:
		return nil
	}
}
