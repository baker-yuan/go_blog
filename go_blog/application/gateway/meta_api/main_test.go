package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
)

func TestMainWrapper(t *testing.T) {
	if os.Getenv("ENV") == "test" {
		t.Skip("skipping build binary when execute unit test")
	}

	var (
		args []string
	)
	for _, arg := range os.Args {
		switch {
		case strings.HasPrefix(arg, "-test"):
		default:
			args = append(args, arg)
		}
	}
	waitCh := make(chan int, 1)
	os.Args = args
	go func() {
		main()
		close(waitCh)
	}()
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-signalCh:
		return
	case <-waitCh:
		return
	}
}
