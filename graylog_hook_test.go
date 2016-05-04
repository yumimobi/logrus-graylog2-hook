package graylog

import (
	"errors"
	"io/ioutil"
	"reflect"
	"testing"
	"time"

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
	if _, err := New("127.0.0.1", "", logrus.DebugLevel); err == nil {
		t.Errorf("expected error not nil from New but get nil")
	}

	hostname = func() (string, error) {
		return "", errors.New("cannot get hostname")
	}
	if _, err := New("127.0.0.1:123", "", logrus.DebugLevel); err.Error() != "cannot get hostname" {
		t.Errorf("expected error cannot get hostname but get %v", err)
	}

	hostname = func() (string, error) {
		return "hostname", nil
	}
	now = func() time.Time {
		return time.Unix(111, 0)
	}
	r, _ := gelf.NewReader("127.0.0.1:0")
	hook, _ := New(r.Addr(), "facility", logrus.InfoLevel)
	logrus.AddHook(hook)
	logrus.SetOutput(ioutil.Discard)

	msgData := "test message\nsecond line"
	logrus.WithField("field", "1").Info(msgData)

	actualMsg, _ := r.ReadMessage()
	expectedMsg := &gelf.Message{
		Version:  "1.1",
		Host:     "hostname",
		Short:    "test message",
		Full:     msgData,
		TimeUnix: 111,
		Level:    int32(logrus.InfoLevel),
		Extra: map[string]interface{}{
			"_field":    "1",
			"_facility": "facility",
		},
	}

	if !reflect.DeepEqual(actualMsg, expectedMsg) {
		t.Errorf("expected message %#v but get %#v", expectedMsg, actualMsg)
	}
}
