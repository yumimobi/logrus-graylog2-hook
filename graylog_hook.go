package graylog

import (
	"bytes"

	"github.com/Graylog2/go-gelf/gelf"
	"github.com/Sirupsen/logrus"
)

// Hook send logs to a logging service compatible with the Graylog API and the GELF format.
type Hook struct {
	w      *gelf.Writer
	levels []logrus.Level
}

var _ logrus.Hook = &Hook{}

// New creates a graylog2 hook
func New(address string, level logrus.Level) (*Hook, error) {
	w, err := gelf.NewWriter(address)

	return &Hook{
		levels: levelThreshold(level),
		w:      w,
	}, err
}

func levelThreshold(l logrus.Level) []logrus.Level {
	for i := range logrus.AllLevels {
		if logrus.AllLevels[i] == l {
			return logrus.AllLevels[:i+1]
		}
	}
	return logrus.AllLevels
}

// Levels implements logrus.Hook interface
func (h *Hook) Levels() []logrus.Level {
	return h.levels
}

// Fire implements logrus.Hook interface
func (h *Hook) Fire(entry *logrus.Entry) error {
	_, err := h.w.Write(bytes.TrimSpace([]byte(entry.Message)))
	return err
}
