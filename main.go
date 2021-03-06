package main

import (
	"github.com/fhmq/hmq/broker"
	"log"
	"os"
	"os/signal"
)

func main() {
	mqttServerStart()
	s := waitForSignal()
	log.Println("signal received, broker closed.", s)
}

func mqttServerStart() {
	config, err := broker.ConfigureConfig(os.Args[1:])
	if err != nil {
		log.Fatal("configure broker config error: ", err)
	}

	b, err := broker.NewBroker(config)
	if err != nil {
		log.Fatal("New Broker error: ", err)
	}

	b.Start()
}


func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
