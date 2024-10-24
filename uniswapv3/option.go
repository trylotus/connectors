package uniswapv3

import (
	"github.com/trylotus/go-connector/common"
	"github.com/trylotus/go-connector/log"
)

type Option func(*Source)

func WithDefaultOptions() Option {
	return func(c *Source) {
		opts := []Option{
			WithQueryPageSize(common.EnvInt64("QUERY_PAGE_SIZE", defaultQueryPageSize)),
			WithSubscriptionPageSize(common.EnvInt64("SUB_PAGE_SIZE", defaultSubscriptionPageSize)),
		}
		for _, opt := range opts {
			opt(c)
		}
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
