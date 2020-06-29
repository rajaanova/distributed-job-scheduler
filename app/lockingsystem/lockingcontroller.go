package lockingsystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"time"
)

type AppLocker struct {
	consulApi     *api.Client
	timeInt       time.Duration
	failureRetry  int
	retryInterval time.Duration
}

type Value struct {
	State   string
	TimeVal int64
}

func NewAppLocker(timeInterval, retryInt time.Duration, retry int) *AppLocker {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic("error starting consul api")
	}
	return &AppLocker{consulApi: client, timeInt: timeInterval, retryInterval: retryInt, failureRetry: retry}

}

func (a *AppLocker) InitKey(key string) {
	kv := a.consulApi.KV()
	va := Value{"ended", time.Now().Unix()}
	data, err := json.Marshal(va)
	if err != nil {
		panic("not able to set the key" + err.Error())
	}
	kvPair := &api.KVPair{Key: key, Value: data}
	_, err = kv.Put(kvPair, nil)
	if err != nil {
		panic("not able to set the key" + err.Error())
	}
}

func (a *AppLocker) Lock() error {
	kv := a.consulApi.KV()
	kvData, _, err := kv.Get("jobStatus", nil)
	if err != nil {
		return err
	}
	oldVal := Value{}
	err = json.Unmarshal(kvData.Value, &oldVal)
	if err != nil {
		return err
	}
	if oldVal.State == "ended" && (time.Now().Unix()-oldVal.TimeVal) > int64(a.timeInt.Seconds()) {
		va := Value{"started", time.Now().Unix()}
		valToSet, err := json.Marshal(va)
		if err != nil {
			return err
		}

		isAccepted, _, err := kv.CAS(&api.KVPair{Key: "jobStatus", Value: valToSet, ModifyIndex: kvData.ModifyIndex}, nil)
		if !isAccepted {
			return fmt.Errorf("some error setting the values %v ", err)

		}
		log.Println("lock on job succeeded")
		return nil
	}

	return errors.New("some other host has access to it ")
}

func (a *AppLocker) Unlock() error {
	count := a.failureRetry
	var err error
	for count > 0 {
		count--
		kv := a.consulApi.KV()
		va, err := json.Marshal(Value{"ended", time.Now().Unix()})
		if err != nil {
			log.Println("marshalling error while sending the updated vale ", err)
			continue
		}
		_, err = kv.Put(&api.KVPair{Key: "jobStatus", Value: va}, nil)
		if err == nil {
			break
		}
		time.Sleep(a.retryInterval)
	}
	return err
}
