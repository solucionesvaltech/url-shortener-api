package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"url-shortener/internal/dependency"
	"url-shortener/pkg/log"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	deps, err := dependency.Init()
	if err != nil {
		log.Log.Fatalf("an error occur in initialize dependencies: %v", err)
	}

	err = deps.Start(context.Background())
	if err != nil {
		log.Log.Fatalf("an error occur in dependencies setup: %v", err)
	}

	log.Log.Infof("starting %v app", deps.Conf.AppName)

	sig := <-quit
	shutdown(sig, deps)
}

func shutdown(sig os.Signal, deps *dependency.Dependency) {
	log.Log.Infof("exit app with signal: %v", sig)
	if deps != nil {
		deps.Shutdown()
	}
}
