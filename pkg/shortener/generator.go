package shortener

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	adjectives = []string{
		"adorable",
		"brave",
		"calm",
		"delightful",
		"eager",
		"friendly",
		"gentle",
		"happy",
		"inventive",
		"jolly",
		"kind",
		"lively",
		"merry",
		"nice",
		"optimistic",
		"polite",
		"quiet",
		"radiant",
		"smart",
		"thoughtful",
		"upbeat",
		"vibrant",
		"witty",
		"xenial",
		"youthful",
		"zesty",
	}

	animals = []string{
		"ant",
		"bat",
		"cat",
		"dog",
		"emu",
		"fox",
		"goat",
		"hare",
		"ibis",
		"jaguar",
		"kiwi",
		"lynx",
		"mole",
		"newt",
		"owl",
		"puma",
		"quokka",
		"rat",
		"seal",
		"tiger",
		"urchin",
		"viper",
		"wolf",
		"xerus",
		"yak",
		"zebra",
	}

	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func generateFriendlyID() string {
	adj := adjectives[rng.Intn(len(adjectives))]
	animal := animals[rng.Intn(len(animals))]
	number := rng.Intn(100) // 0–99
	return fmt.Sprintf("%s-%s-%02d", adj, animal, number)
}
