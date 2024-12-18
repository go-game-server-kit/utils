package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func Graceful() os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	return <-quit
}
