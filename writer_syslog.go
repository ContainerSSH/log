package log

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func newSyslogWriter(config SyslogConfig) (Writer, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &syslogWriter{
		lock:       &sync.Mutex{},
		connection: config.connection,
		config:     config,
	}, nil
}

type syslogWriter struct {
	connection net.Conn
	config     SyslogConfig
	lock       *sync.Mutex
}

func (s *syslogWriter) Write(level Level, message Message) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	facilityNumber, err := s.config.Facility.Number()
	if err != nil {
		return err
	}
	pri := int64(facilityNumber)*8 + int64(level)
	t := time.Now()
	timestamp := fmt.Sprintf(
		"%s %2d %02d:%02d:%02d",
		t.Format("Feb"),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	tag := s.config.tag
	if s.config.Pid {
		tag += fmt.Sprintf("[%d]", os.Getpid())
	}
	msg := message.Explanation()
	line := fmt.Sprintf("<%d>%s %s %s: %s\n", pri, timestamp, s.config.hostname, tag, msg)
	if _, err = s.connection.Write([]byte(line)); err != nil {
		return Wrap(err, ELogWriteFailed, "failed to write to syslog socket")
	}
	return nil
}

func (s *syslogWriter) Rotate() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if err := s.config.Validate(); err != nil {
		return err
	}
	if err := s.connection.Close(); err != nil {
		return Wrap(err, ELogRotateFailed, "failed to close old syslog connection")
	}
	s.connection = s.config.connection
	return nil
}

func (s *syslogWriter) Close() error {
	return s.connection.Close()
}
