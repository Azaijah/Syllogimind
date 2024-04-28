package main

import (
	"os"

	"github.com/Azaijah/Syllogimind/syllogism"
	"github.com/Azaijah/Syllogimind/syllogism/gtlt"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

// Define a struct that matches the TOML file structure

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	syllogism.Init()

	gtlt.GenerateGTLT()

	//syllo.Generate()
	//syllo.Show()

}
