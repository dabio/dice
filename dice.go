package dice

import (
	"math/rand"
	"time"
)

const (
	MaxCount     = 100
	DefaultCount = 3
)

func Roll(count int) []string {
	rand.Seed(time.Now().UnixNano())
	selected := []string{}
	numWords := len(words)

	if count > MaxCount {
		count = MaxCount
	}
	if count < 1 {
		count = 1
	}

	for {
		word := words[rand.Intn(numWords)]
		// check if word is already in our result list
		found := false
		for _, w := range selected {
			if w == word {
				found = true
				break
			}
		}
		if !found {
			selected = append(selected, word)
		}

		if len(selected) >= count {
			break
		}
	}

	return selected
}
