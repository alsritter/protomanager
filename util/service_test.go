package util

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/alsritter/protomanager/util/logger"
)

func TestStartServiceAsync(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	StartServiceAsync(ctx, logger.NewDefault("test"), cancel, func() error {
		// Here is the initialization project
		log.Println("TestServer Starting...")
		return nil
	}, func() error {
		// Call if cancel is closed
		log.Println("TestServer Closed...")
		return nil
	})

	time.Sleep(time.Second * 1)

	// close.
	cancel()

	time.Sleep(time.Second * 2)
}
