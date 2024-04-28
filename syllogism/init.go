package syllogism

import (
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

var system syllogismSystem
var initialized bool

func Init() {
	log.Debug().Msg("Initializing syllogism package")
	data, err := os.ReadFile("words.yaml")
	if err != nil {
		log.Fatal().Msgf("error: %v", err.Error())
	}

	err = yaml.Unmarshal(data, &system)
	if err != nil {
		log.Fatal().Msgf("error: %v", err.Error())
	}

	log.Trace().Any("system", system).Msg("")

	initialized = true
}

func IsInitialized() bool {
	return initialized
}
