package main

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

const ExitCodeInterrupt = 2

func listenForCancellationAndAddToContext() (ctx context.Context, done func()) {
	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	done = func() {
		signal.Stop(signalChan)
		cancel()
	}

	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			fmt.Println("called signalChan")
			cancel()
		case <-gracefulShutdown:
			fmt.Println("\nCancelling contexts...... Shutting down application gracefully")
			cancel()
			os.Exit(ExitCodeInterrupt)
		case <-ctx.Done():
			fmt.Println("called ctx.Done()")
		}

		<-signalChan // second signal, hard exit
		os.Exit(ExitCodeInterrupt)
	}()

	return ctx, done
}
