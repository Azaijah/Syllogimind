package gtlt

import (
	"fmt"

	"github.com/Azaijah/Syllogimind/syllogism"
	"github.com/rs/zerolog/log"
)

func GenerateGTLT() {

	if !syllogism.IsInitialized() {
		log.Fatal().Msg("syllogism system has not been initialized")
	}

	syllo := syllogism.Syllogism{}
	syllo.Entities = make([]syllogism.Entity, 3)

	for i := 0; i < 3; i++ {
		syllo.Entities[i].Text = syllogism.GetRandomWord()
	}

	log.Debug().Any("entities", syllo.Entities).Msg("")

	fmt.Printf(statement_1, syllo.Entities[0].Text, syllo.Entities[1].Text, syllo.Entities[2].Text, syllo.Entities[1].Text, syllo.Entities[0].Text, syllo.Entities[2].Text, syllo.Entities[2].Text, syllo.Entities[1].Text)

}
