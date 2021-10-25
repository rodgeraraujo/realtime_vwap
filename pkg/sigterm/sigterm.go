package sigterm

import (
	"os"
	"os/signal"
	"syscall"
)

func HnderSigterm() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}
