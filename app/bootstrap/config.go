package bootstrap

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	FailureRetry         int           `envconfig:"FAILURE_RETRY"`
	MinSchedulerInterval time.Duration `envconfig:"MIN_SCHEDULER_INTERVAL"`
	SchedulerInterval    time.Duration `envconfig:"SCHEDULER_INTERVAL"`
	FailureRetryInterval time.Duration `envconfig:"FAILURE_RETRY_INTERVAL"`
}

func (a *Config) Boot() error {
	return envconfig.Process("", a)
}
