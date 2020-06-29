package controller

import (
	"log"
	"os"
	"time"
)

type lockSystem interface {
	Lock() error
	Unlock() error
}
type JobScheduler struct {
	interval time.Duration
	locking  lockSystem
}

func NewNewJobScheduler(intVal time.Duration, lockSys lockSystem) *JobScheduler {
	return &JobScheduler{intVal, lockSys}
}

func (a *JobScheduler) InitJob() {
	t := time.NewTicker(a.interval)
	defer t.Stop()
	for range t.C {
		a.run()
	}
}
func (a *JobScheduler) run() {
	if err := a.locking.Lock(); err == nil {
		OneTimeJob()
		err := a.unlock()
		if err != nil {
			log.Println("error updating the key to finished state", err)
		}
	} else {
		hostName, _ := os.Hostname()
		log.Printf("%s  didn't get the lock to execute", hostName)
	}

}

func (a *JobScheduler) unlock() error {
	return a.locking.Unlock()
}

func OneTimeJob() {
	//suppose some one time job has run
	time.Sleep(5 * time.Second)
	log.Println("the job is completed")
}
