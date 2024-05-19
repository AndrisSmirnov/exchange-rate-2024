package main

import (
	"context"
	"exchange_rate/pkg/app"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	app, err := app.New(cancelFunc, ctx)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		logrus.Panic(err)
	}

	logrus.Info("Exchange Rate launched (/◔ ◡ ◔)/")
	go syscallWait(cancelFunc)

	<-ctx.Done()
	logrus.Info("Exchange Rate stopping \t(◑ _ ◑)")

	app.Stop()
	logrus.Info("Exchange Rate stopped \t(✖ _ ✖)")
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}
