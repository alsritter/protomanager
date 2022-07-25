package util

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
)

func TestRegisterExitHandlers(t *testing.T) {
	_, cancel := context.WithCancel(context.Background())
	stop := RegisterExitHandlers(cancel)

	// do something...
	time.Sleep(2 * time.Second)
	sendInterruptSignal()

	<-stop

	log.Println("server closed")
}

// ctrl + c
func sendInterruptSignal() error {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}
	return p.Signal(os.Interrupt)
}
