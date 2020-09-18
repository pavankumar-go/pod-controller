package main

import (
	"os"
	"os/signal"

	"github.com/pod-controller/controller"
	"github.com/sirupsen/logrus"
)

func main() {
	stopCh := make(chan struct{})
	defer func() {
		logrus.Info("Stopping Informer...")
		close(stopCh)
	}()

	informer := controller.GetInformer()
	logrus.Info("Informer is up!")
	go informer.Run(stopCh)

	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Kill, os.Interrupt)
	<-sigCh
}
