package main

import (
	"log"

	"github.com/Pratam-Kalligudda/payment-service-go/config"
)

func main() {
	_, err := config.SetupConfig()
	if err != nil {
		log.Fatal("error while loading config : " + err.Error())
	}
}
