package main

import (
	"authentication-service/internal/channels/rest"
	"authentication-service/internal/config"

	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	if err := rest.NewRestChannel().Start(); err != nil {
		logrus.Panic()
	}
}
