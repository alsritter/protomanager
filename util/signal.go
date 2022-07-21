package util

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// RegisterExitHandlers is used to register exit handlers
func RegisterExitHandlers(cancelFunc func()) (stop chan struct{}) {
	var exitSignals = []os.Signal{os.Interrupt, syscall.SIGTERM} // SIGTERM is POSIX specific

	stop = make(chan struct{})
	s := make(chan os.Signal, len(exitSignals))
	signal.Notify(s, exitSignals...)

	go func() {
		// Wait for a signal from the OS before dispatching
		// a stop signal to all other goroutines observing this channel.
		<-s
		log.Println("exit signal received")
		if cancelFunc != nil {
			cancelFunc()
		}
		close(stop)
	}()

	return stop
}
