package application

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	sigkill chan os.Signal
	sighup  chan os.Signal
}

func NewApplication() *Application {
	return &Application{
		sigkill: make(chan os.Signal, 1),
		sighup:  make(chan os.Signal, 1),
	}
}

//func (a *Application) Start(ctx context.Context) error {
//	a.startTray(ctx)
//	return nil
//}

func (a *Application) start(ctx context.Context) error {
	signal.Notify(a.sigkill, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	signal.Notify(a.sighup, syscall.SIGHUP)

	slog.Info("Starting application")

	//var err error
	for {
		select {
		case sigc := <-a.sigkill:
			slog.Info("Signal %v received", sigc)
			return a.stop(ctx)
		case sigc := <-a.sighup:
			slog.Info("reload Signal %v received", sigc)

		}

	}
}

func (a *Application) Stop() {
	return
}
