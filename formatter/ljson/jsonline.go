package ljson

import (
	"github.com/containerssh/log"
)

type JsonLine struct {
	Time    string          `json:"timestamp"`
	Level   log.LevelString `json:"level"`
	Message string          `json:"message,omitempty"`
	Details interface{}     `json:"details,omitempty"`
}






