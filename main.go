package main

import (
	"os"

	"github.com/Azaijah/Syllogimind/syllogism"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	syllo := syllogism.Syllogism{}

	syllo.Generate()
	syllo.Show()

}
