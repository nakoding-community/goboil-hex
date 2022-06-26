package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nakoding-community/golang-boilerplate/internal/config"
	"github.com/nakoding-community/golang-boilerplate/internal/factory"
	"github.com/nakoding-community/golang-boilerplate/pkg/util"

	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("ENV") == "" {
		env := util.NewEnv()
		env.Load()
	}

	err := config.Load("")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func main() {
	//initialize all needs (db conn, adapter outbound)
	f := initFactory()
	starterEcho, stopperEcho := f.Adapter.InBound.Rest.PrepareEcho()

	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		startProcessAtBackground(starterEcho)
		gracefullStopProcessAtBackground(time.Second*10, stopperEcho)
		wg.Done()
	}()

	//sample multiple process
	go func() {
		time.Sleep(time.Second * 20)
		logrus.Info("testing log")
		wg.Done()
	}()
	wg.Wait()
}

func initFactory() *factory.Factory {
	return factory.NewFactory()
}

func startProcessAtBackground(ps ...func() error) {
	for _, p := range ps {
		if p != nil {
			go func(_p func() error) {
				_ = _p()
			}(p)
		}
	}
}

func gracefullStopProcessAtBackground(duration time.Duration, ps ...func(ctx context.Context) error) {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	for _, p := range ps {
		if p == nil {
			continue
		}
		ctx, stop := context.WithTimeout(context.Background(), duration)
		defer stop()
		_ = p(ctx)

	}
}
