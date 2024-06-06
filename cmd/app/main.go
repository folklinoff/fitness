package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/folklinoff/fitness-app/cmd/app/processor"
)

func main() {
	errors := make(chan error)
	go func() {
		errors <- processor.Run()
	}()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errors:
		log.Println(err)
	case <-stop:
		processor.Shutdown(context.Background())
	}
}
