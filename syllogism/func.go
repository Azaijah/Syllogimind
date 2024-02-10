package syllogism

import (
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

func generateEntities() map[int]*entity {
	entities := []entity{
		{text: "BAP"},
		{text: "ZIG"},
		{text: "QOX"},
	}

	entityMap := make(map[int]*entity)

	numbers := make([]int, len(entities))
	for i := range numbers {
		numbers[i] = i + 1
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	rnd.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	for i, num := range numbers {
		entityMap[num] = &entities[i]
	}
	log.Debug().Any("Entity map", entityMap[1].text).Msg("Entity map")
	return entityMap
}

func selectTwoEntities(entities map[int]*entity) (e1, e2 *entity, relationType relationType) {

	keys := make([]int, 0, len(entities))
	for k := range entities {
		keys = append(keys, k)
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	rnd.Shuffle(len(keys), func(i, j int) {
		keys[i], keys[j] = keys[j], keys[i]
	})

	if len(keys) >= 2 {
		e1 = entities[keys[0]]
		e2 = entities[keys[1]]

		log.Trace().Msg("Random Entity 1: " + e1.text)
		log.Trace().Msg("Random Entity 2: " + e2.text)
	} else {
		log.Error().Msg("Not enough entities to select two randomly.")
	}

	if keys[0] > keys[1] {
		relationType = moreThan
	} else {
		relationType = lessThan
	}

	return e1, e2, relationType

}
