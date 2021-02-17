package log

import (
	"testing"
)

func newGoTest(t *testing.T) Writer {
	return &goTestWriter{
		t: t,
	}
}

type goTestWriter struct {
	t *testing.T
}

func (g *goTestWriter) Write(level Level, message Message) error {
	levelString, err := level.Name()
	if err != nil {
		return err
	}
	g.t.Logf("%s\t%s", levelString, message.Explanation())
	return nil
}

func (g *goTestWriter) Rotate() error {
	return nil
}

func (g *goTestWriter) Close() error {
	return nil
}
