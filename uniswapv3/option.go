package uniswapv3

import (
	"os"
	"strconv"

	"github.com/trylotus/go-connector/log"
)

type Option func(*Source)

func WithDefaultOptions() Option {
	return func(c *Source) {
		opts := []Option{
			WithConcurrencyLimit(getEnvInt64("CONCURRENCY_LIMIT", defaultConcurrenyLimit)),
			WithBlockRangeLimit(getEnvInt64("BLOCK_RANGE_LIMIT", defaultBlockRangeLimit)),
			WithQueryPageSize(getEnvInt64("QUERY_PAGE_SIZE", defaultQueryPageSize)),
			WithSubscriptionPageSize(getEnvInt64("SUB_PAGE_SIZE", defaultSubscriptionPageSize)),
		}
		for _, opt := range opts {
			opt(c)
		}
	}
}

func WithConcurrencyLimit(concurrencyLimit int64) Option {
	if concurrencyLimit <= 0 {
		log.Fatal().Msg("concurrencyLimit must be positive")
	}
	return func(s *Source) {
		s.concurrecyLimit = concurrencyLimit
	}
}

func WithBlockRangeLimit(blockRangeLimit int64) Option {
	if blockRangeLimit <= 0 {
		log.Fatal().Msg("blockRangeLimit must be positive")
	}
	return func(s *Source) {
		s.blockRangeLimit = blockRangeLimit
	}
}

func WithQueryPageSize(queryPageSize int64) Option {
	if queryPageSize <= 0 {
		log.Fatal().Msg("queryPageSize must be positive")
	}
	return func(s *Source) {
		s.queryPageSize = queryPageSize
	}
}

func WithSubscriptionPageSize(subscriptionPageSize int64) Option {
	if subscriptionPageSize <= 0 {
		log.Fatal().Msg("subscriptionPageSize must be positive")
	}
	return func(s *Source) {
		s.subscriptionPageSize = subscriptionPageSize
	}
}

func getEnvInt64(key string, defaultValue int64) int64 {
	valueTxt := os.Getenv(key)
	if valueTxt == "" {
		return defaultValue
	}

	value, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		log.Fatal().Msgf("Invalid env: %s", key)
	}

	return value
}
