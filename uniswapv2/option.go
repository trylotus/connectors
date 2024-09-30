package uniswapv2

import "github.com/trylotus/go-connector/log"

type Option func(*Source)

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
