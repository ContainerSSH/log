package log

import (
	"encoding/json"
	"fmt"
	"io"
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
	line, err := f.createLine(levelString, message)
	if err != nil {
		return WrapError(err, ELogWriteFailed, "failed to write log message")
	}
	if _, err := f.fh.Write(line); err != nil {
		return WrapError(err, ELogWriteFailed, "failed to write log message")
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
		line, err = json.Marshal(
			jsonLine{
				Time:    time.Now().Format(time.RFC3339),
				Level:   string(levelString),
				Message: message,
			},
		)
		if err != nil {
			return nil, err
		}
	case FormatText:
		line = []byte(fmt.Sprintf(
			"%s\t%s\t%s\n",
			time.Now().Format(time.RFC3339),
			levelString,
			message.Explanation(),
		))
	default:
		return nil, fmt.Errorf("log format not supported: %s", f.format)
	}
	return line, nil
}

type jsonLine struct {
	Time  string
	Level string

	Message
}

func (j jsonLine) MarshalJSON() ([]byte, error) {
	data := map[string]interface{}{}
	data["timestamp"] = j.Time
	data["level"] = j.Level
	data["message"] = j.Explanation()
	details := map[string]interface{}{}
	for label, value := range j.Labels() {
		details[string(label)] = value
	}
	data["details"] = details
	return json.Marshal(data)
}
