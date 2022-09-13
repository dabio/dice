package dice

import (
	"math/rand"
	"time"
)

func Roll(count int) []string {
	rand.Seed(time.Now().UnixNano())
	selected := []string{}
	numWords := len(words)

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
