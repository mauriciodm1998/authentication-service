package main

import (
	"authentication-service/internal/channels/rest"
	"authentication-service/internal/config"

	"github.com/rs/zerolog/log"
)

func main() {
	config.ParseFromFlags()

	if err := rest.NewRestChannel().Start(); err != nil {
		log.Panic().Err(err).Msg("an error occurred when run rest service")
	}
}
