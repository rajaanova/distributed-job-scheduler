package main

import (
	"github.com/rajaanova/distributed-job-scheduler/app/bootstrap"
	"github.com/rajaanova/distributed-job-scheduler/app/controller"
	"github.com/rajaanova/distributed-job-scheduler/app/lockingsystem"
	"log"
)

func main() {
	conf := &bootstrap.Config{}
	if conf.Boot() != nil {
		panic("error bootstrapping the configuration value")
	}
	log.Println(conf)
	lockingSys := lockingsystem.NewAppLocker(conf.MinSchedulerInterval, conf.FailureRetryInterval, conf.FailureRetry)
	lockingSys.InitKey("jobStatus")
	scheduler := controller.NewNewJobScheduler(conf.SchedulerInterval, lockingSys)
	scheduler.InitJob()
}
