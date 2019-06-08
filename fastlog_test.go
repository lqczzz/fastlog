package fastlog

import (
	"testing"
	"time"
)

func TestBase(t *testing.T) {
	Init()

	Error("test Error")
	time.Sleep(2 * time.Second)

	Info("test Debug")
	time.Sleep(2 * time.Second)
}
