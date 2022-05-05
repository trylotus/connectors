package ethnft

/*

import (
	"context"
	"strings"
	"testing"

	//ethmocks "blep.ai/data/chain/ethereum/ethclient/mocks"
	"github.com/nakji-network/connector/examples/nft/erc1155"
	"github.com/nakji-network/connector/examples/nft/erc721"
	//"blep.ai/data/connectors/tests"

	kafkamocks "blep.ai/data/kafkautils/mocks"
	"github.com/nakji-network/connector/kafkautils"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/golang/mock/gomock"
)

type MockSubscription struct {
	errorChan chan error
}

func (m MockSubscription) Err() <-chan error {
	return m.errorChan
}

func (MockSubscription) Unsubscribe() {
	// noop
}

func TestStartConnector(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProducer := kafkamocks.NewMockProducerInterface(mockCtrl)
	mockETHClientPool := ethmocks.NewMockETHClientPool(mockCtrl)

	kafkautils.TopicTypeRegistry.Load(TopicTypes)

	topics := map[string]kafkautils.Topic{
		"erc1155_approvalforall": kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc1155.approvalforall.0_0_0", "test"),
		"erc1155_transferbatch":  kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc1155.transferbatch.0_0_0", "test"),
		"erc1155_transfersingle": kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc1155.transfersingle.0_0_0", "test"),
		"erc1155_uri":            kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc1155.uri.0_0_0", "test"),
		"erc721_approval":        kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc721.approval.0_0_0", "test"),
		"erc721_approvalforall":  kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc721.approvalforall.0_0_0", "test"),
		"erc721_transfer":        kafkautils.MustParseTopic(".fct.nakji.nft.0_0_0.erc721.transfer.0_0_0", "test"),
	}

	erc1155Address := ethcommon.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	erc721Address := ethcommon.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")

	addresses := map[string][]ethcommon.Address{
		"erc1155": {erc1155Address},
		"erc721":  {erc721Address},
	}

	connector := NewConnector(
		mockProducer,
		addresses,
		topics,
		mockETHClientPool,
	)

	ctx, cancel := context.WithCancel(context.Background())
	connector.setupMocks(cancel)
	connector.Start(ctx, 10)
}

func (c *NftConnector) setupMocks(cancel context.CancelFunc) {
	mockETHClientPool := c.ClientPool.(*ethmocks.MockETHClientPool)
	mockProducer := c.KP.(*kafkamocks.MockProducerInterface)

	mockETHClientPool.EXPECT().ConsumeHeaders(gomock.Any())

	mockSubscription := MockSubscription{
		errorChan: make(chan error),
	}

	mockETHClientPool.EXPECT().ChunkedFilterLogs(context.Background(), gomock.Any(), int64(100), int64(1), gomock.Any(), gomock.Any()).AnyTimes()
	mockETHClientPool.EXPECT().GetLogTimestamp(gomock.Any(), c.blockCache).AnyTimes().Return(uint64(1000), nil)

	erc1155Query := geth.FilterQuery{
		Addresses: c.addresses["erc1155"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), erc1155Query, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			erc1155Abi, err := abi.JSON(strings.NewReader(erc1155.Erc1155ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Erc1155 ABI")
			}

			tests.WriteEventLogs(eventLogs, erc1155Abi)
			return mockSubscription, nil
		},
	)

	erc721Query := geth.FilterQuery{
		Addresses: c.addresses["erc721"],
	}
	mockETHClientPool.EXPECT().SubscribeFilterLogs(gomock.Any(), erc721Query, gomock.Any()).DoAndReturn(
		func(ctx context.Context, query geth.FilterQuery, eventLogs chan<- ethtypes.Log) (event.Subscription, error) {
			erc721Abi, err := abi.JSON(strings.NewReader(erc721.Erc721ABI))
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to read Erc721 ABI")
			}

			tests.WriteEventLogs(eventLogs, erc721Abi)
			return mockSubscription, nil
		},
	)

	mockProducer.EXPECT().EnableTransactions()
	counter := 0
	mockProducer.EXPECT().WriteAndCommitSink(gomock.Any()).DoAndReturn(
		func(sink <-chan *kafkautils.Message) {
			go func() {
				for range sink {
					counter++

					if counter >= 7 {
						cancel()
					}
				}
			}()
		},
	)
}

*/
