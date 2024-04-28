package syllogism

import (
	"math/rand"
	"time"
)

func GetRandomWord() string {
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)

	return system.WordList[localRand.Intn(len(system.WordList))]

}
