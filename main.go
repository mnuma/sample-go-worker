package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	os.Exit(run())
}

func run() int {
	ret := loop()
	return ret
}

func loop() int {

	log.Println("RUN")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	code := 0
	select {
	case s := <-sigChan:
		switch s {
		// kill -SIGHUP XXXX
		case syscall.SIGHUP:
			log.Println("SIGHUP")
		// kill -SIGINT XXXX or Ctrl+c
		case syscall.SIGINT:
			log.Println("SIGINT")
		// kill -SIGTERM XXXX
		case syscall.SIGTERM:
			log.Println("SIGTERM")
		// kill -SIGQUIT XXXX
		case syscall.SIGQUIT:
			log.Println("SIGQUIT")
		default:
			log.Printf("Unknown signal. %v", s)
			code = 1
		}
	}
	return code
}
