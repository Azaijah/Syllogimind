package syllogism

import (
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"
)

func (s *Syllogism) generatePremises(entities map[int]*entity) {
	var premises []premise

	var prevRel relationship
	for {

		e1, e2, relationType := selectTwoEntities(entities)

		curRel := relationship{
			from: e1,
			to:   e2,
			text: relationType}

		if prevRel == curRel {
			log.Debug().Msg("premise relationships are equal... generating new")

			continue
		}

		premises = append(premises, premise{
			entity1:    *e1,
			entity2:    *e2,
			relation:   curRel,
			isInverted: false,
		})

		if len(premises) == 2 {
			break
		}

		prevRel = curRel

	}

	s.premises = premises
}

func (s *Syllogism) generateConclusion(entities map[int]*entity) {

	var statement statementType
	var statements []*relationship

	var prevRel relationship
	for {
		e1, e2, relationType := selectTwoEntities(entities)

		curRel := relationship{from: e1, to: e2, text: relationType}

		if prevRel == curRel {
			log.Debug().Msg("conclusion relationships are equal... generating new")
			continue
		}
		statements = append(statements, &curRel)

		if len(statements) == 2 {
			break
		}

		prevRel = curRel

	}

	if rand.Intn(2) == 0 { // There's a 50/50 chance since Intn(2) returns either 0 or 1
		statement = sameR
	} else {
		statement = diffR
	}

	var outcome bool
	if statements[0].text == statements[1].text && statement == sameR {
		outcome = true
	} else if statements[0].text != statements[1].text && statement == sameR {
		outcome = false
	} else if statements[0].text == statements[1].text && statement == diffR {
		outcome = false
	} else if statements[0].text != statements[1].text && statement == diffR {
		outcome = true
	}

	s.conclusion = conclusion{
		statement1:    statements[0],
		statement2:    statements[1],
		statementType: statement,
		isInverted:    false,
		outcome:       outcome,
	}

}

func (s *Syllogism) Generate() {
	log.Info().Msg("Generating syllogism")
	entities := generateEntities()
	s.generatePremises(entities)
	s.generateConclusion(entities)
}

func (s *Syllogism) Show() {

	fmt.Println(s.premises[0].entity1.text + " " + string(s.premises[0].relation.text) + " " + s.premises[0].entity2.text)
	fmt.Println(s.premises[1].entity1.text + " " + string(s.premises[0].relation.text) + " " + s.premises[1].entity2.text)
	fmt.Println()
	fmt.Println(s.conclusion.statement1.from.text, "to", s.conclusion.statement1.to.text)
	fmt.Println(s.conclusion.statementType)
	fmt.Println(s.conclusion.statement2.from.text, "to", s.conclusion.statement2.to.text)

	fmt.Println(s.conclusion.outcome)

}
