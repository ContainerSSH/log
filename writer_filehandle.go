package log

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

func newFileHandleWriter(fh io.Writer, format Format, lock *sync.Mutex) *fileHandleWriter {
	return &fileHandleWriter{
		fh:     fh,
		lock:   lock,
		format: format,
	}
}

type fileHandleWriter struct {
	lock   *sync.Mutex
	fh     io.Writer
	format Format
}

func (f *fileHandleWriter) Write(level Level, message Message) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	levelString, err := level.Name()
	if err != nil {
		return err
	}
	line, err := f.createLine(levelString, message)
	if err != nil {
		return Wrap(err, ELogWriteFailed, "failed to write log message")
	}
	if _, err := f.fh.Write(append(line, '\n')); err != nil {
		return Wrap(err, ELogWriteFailed, "failed to write log message")
	}
	return nil
}

func (f *fileHandleWriter) Rotate() error {
	return nil
}

func (f *fileHandleWriter) Close() error {
	return nil
}

func (f *fileHandleWriter) createLine(levelString LevelString, message Message) (line []byte, err error) {
	switch f.format {
	case FormatLJSON:
		details := map[string]interface{}{}
		for label, value := range message.Labels() {
			details[string(label)] = value
		}
		line, err = json.Marshal(
			jsonLine{
				Time:    time.Now().Format(time.RFC3339),
				Code:    message.Code(),
				Level:   string(levelString),
				Message: message.Explanation(),
				Details: details,
			},
		)
		if err != nil {
			return nil, err
		}
	case FormatText:
		msg := message.Explanation()
		var labels []string
		for labelName, labelValue := range message.Labels() {
			labels = append(labels, fmt.Sprintf("%s=%s", labelName, labelValue))
		}
		if len(labels) > 0 {
			msg += fmt.Sprintf(" (%s)", strings.Join(labels, " "))
		}
		line = []byte(fmt.Sprintf(
			"%s\t%s\t%s\n",
			time.Now().Format(time.RFC3339),
			levelString,
			msg,
		))
	default:
		return nil, fmt.Errorf("log format not supported: %s", f.format)
	}
	return line, nil
}

type jsonLine struct {
	Time    string                 `json:"timestamp"`
	Level   string                 `json:"level"`
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details"`
}
