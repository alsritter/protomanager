package util

import (
	"context"
	"log"
)

// StartServiceAsync is used to start service async
func StartServiceAsync(ctx context.Context, logger *log.Logger, cancelFunc context.CancelFunc, serveFn func() error, stopFn func() error) {
	if serveFn == nil {
		return
	}
	go func() {
		logger.Println("starting service")
		go func() {
			if err := serveFn(); err != nil {
				logger.Printf("error serving service: %s \n", err)
			}
			if cancelFunc != nil {
				cancelFunc()
			}
		}()

		<-ctx.Done()
		logger.Println("stopping service")

		if stopFn() != nil {
			logger.Println("stopping service gracefully")
			if err := stopFn(); err != nil {
				logger.Printf("error occurred while stopping service: %s \n", err)
			}
		}
		logger.Println("exiting service")
	}()
}
