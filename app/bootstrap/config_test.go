package bootstrap

import (
	"os"
	"testing"
)

func TestConfig_Boot(t *testing.T) {
	os.Setenv("FAILURE_RETRY","19")
	a := &Config{}
	err := a.Boot()
	if err!= nil {
		t.Errorf("error processign the config value")
	}
	if a.FailureRetry!=19{
		t.Errorf("wrong value for failure retry")
	}
}
