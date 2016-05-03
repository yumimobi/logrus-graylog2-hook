package graylog

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/Sirupsen/logrus"
)

func TestLevelThreshold(t *testing.T) {
	levels := levelThreshold(logrus.DebugLevel + 1)
	if !reflect.DeepEqual(levels, logrus.AllLevels) {
		t.Errorf("levels should resemble %v", logrus.AllLevels)
	}
}

func TestHook(t *testing.T) {
	_, err := New("127.0.0.1", logrus.DebugLevel)
	if err == nil {
		t.Errorf("expected error not nil from New but get nil")
	}

	r, _ := gelf.NewReader("127.0.0.1:0")
	hook, _ := New(r.Addr(), logrus.InfoLevel)
	logrus.AddHook(hook)
	logrus.SetOutput(ioutil.Discard)

	msgData := "test message\nsecond line"
	logrus.WithField("field", "1").Info(msgData)

	msg, _ := r.ReadMessage()

	if expected := "test message"; msg.Short != expected {
		t.Errorf("msg.Short expected %s but get %s", expected, msg.Short)
	}

	if expected := msgData; msg.Full != expected {
		t.Errorf("msg.Full expected %s but get %s", expected, msg.Full)
	}
}
