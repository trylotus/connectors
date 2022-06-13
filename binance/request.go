package binance

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

// http get and read body as bytes
func Get(baseURL string, params url.Values) ([]uint8, error) {
	resp, err := http.Get(baseURL + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// http get and read body as json
func GetJSON(baseURL string, params url.Values) (*gjson.Result, error) {
	body, err := Get(baseURL, params)
	if err != nil {
		return nil, err
	}

	result := gjson.ParseBytes(body)
	return &result, nil
}

// Interval, Margin, Delay are in ms
// `decay` is the muliple at which the delay decreases.
type TimedFetchOptions struct {
	Interval,
	Margin, // starts the requests a bit earlier in case their emission schedule is shorter than advertised
	Delay,
	Decay int // decay factor for each round. Affects how fast it syncs to remote emission at the cost of more requests.
}

// timedFetch makes requests to API endpoint that updates every expected
// `interval`, resistant to clock drift by 1 `margin` ahead and n `delay`
// behind.
// Algo: 1m 1m 1m 1m. If lastVal != val, 4m. 15s 15s 15s 15s. If lastVal != val, 4m.59s. 1s 1s 1s 1s. If lastVal != val, 4m.59.5s
func TimedFetch(baseURL string, params url.Values, options TimedFetchOptions, callback func(*gjson.Result, error)) {
	delay := options.Delay
	currentDuration := 0
	ticker := time.NewTicker(time.Duration(options.Delay) * time.Millisecond)
	defer ticker.Stop()

	var lastVal string
	for ; true; <-ticker.C {
		jsonData, err := GetJSON(baseURL, params)
		if err != nil {
			callback(nil, err)
		}
		result := jsonData.Array()
		val := result[len(result)-1]

		if val.Raw == lastVal {
			if currentDuration != delay {
				ticker.Stop()
				ticker = time.NewTicker(time.Duration(delay) * time.Millisecond)
				currentDuration = delay
			}
			log.Info().Str("time", time.Now().String()).Int("delay", delay)
			continue
		}
		// data updated, so it must have updated within the last ticker shortDuration
		ticker.Stop()
		ticker = time.NewTicker(time.Duration(options.Interval-delay-options.Margin) * time.Millisecond) // -550 gives us 550 ms of margin of clock skew (speedup)
		currentDuration = options.Interval - delay - options.Margin
		log.Info().Str("time", time.Now().String()).Int("delay", options.Interval-delay-options.Margin)
		delay = max(delay/options.Decay, 100) // minimum of 100ms between requests
		lastVal = val.Raw

		callback(&val, nil)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
