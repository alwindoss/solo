package main

import (
	"flag"
	"log"
	"os"
)

var (
	logLocation string
	logger      *log.Logger
)

func main() {
	setupFlags()
	if logLocation != "" {
		f, err := os.Open(logLocation)
		if err != nil {
			log.Fatal(err)
		}
		logger = log.New(f, "solo ", log.Ldate)
	} else {
		logger = log.New(os.Stdout, "solo ", log.LstdFlags)
	}
	logger.Printf("Welcome to Solo")
}

func setupFlags() {
	flag.StringVar(&logLocation, "log-loc", "", "-log-loc console.log")
	flag.Parse()
}
