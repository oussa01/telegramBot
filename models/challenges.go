package models

import (
	"math/rand"
	"strings"
	"time"
)

type Challenge struct {
	Question string
	Answer   string
	Type     string
	Hint     string
}

var Challenges = []Challenge{
	{"What is 2+2?", "4", "Quiz", "It’s a single-digit number."},
	{"Decode: 'Khoor'", "hello", "Decode", "Caesar shift +3."},
	{"What was the first gift you gave me?", "notebook", "Memory", "You can write in it."},
	// Add more challenges as needed
}

func FilterByType(challenges []Challenge, answered map[string]bool, mode string) []Challenge {
	mode = strings.ToLower(mode)
	var result []Challenge
	for _, c := range challenges {
		if strings.ToLower(c.Type) == mode {
			if answered == nil || !answered[c.Question] {
				result = append(result, c)
			}
		}
	}
	return result
}

func RandomChallenge(challenges []Challenge, answered map[string]bool) Challenge {
	var pool []Challenge
	for _, c := range challenges {
		if answered == nil || !answered[c.Question] {
			pool = append(pool, c)
		}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return pool[r.Intn(len(pool))]
}

func MatchAnswer(given, expected string) bool {
	if expected == "" {
		return true
	}
	return strings.TrimSpace(strings.ToLower(given)) == strings.TrimSpace(strings.ToLower(expected))
}
