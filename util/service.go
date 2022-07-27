package util

import (
	"context"

	"github.com/alsritter/protomanager/util/logger"
)

// StartServiceAsync is used to start service async
func StartServiceAsync(ctx context.Context, logger logger.Logger, cancelFunc context.CancelFunc, serveFn func() error, stopFn func() error) {
	if serveFn == nil {
		return
	}
	go func() {
		logger.Info(nil, "starting service")
		go func() {
			if err := serveFn(); err != nil {
				logger.Info(nil, "error serving service: %s \n", err)
			}
			if cancelFunc != nil {
				cancelFunc()
			}
		}()

		<-ctx.Done()
		logger.Info(nil, "stopping service")

		if stopFn() != nil {
			logger.Info(nil, "stopping service gracefully")
			if err := stopFn(); err != nil {
				logger.Info(nil, "error occurred while stopping service: %s \n", err)
			}
		}
		logger.Info(nil, "exiting service")
	}()
}
