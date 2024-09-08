package domain

import "strings"

type DifficultiesMap struct {
	mp map[string][]Word
}

func NewDifficultiesMap(difficulties map[string][]Word) DifficultiesMap {
	mp := make(map[string][]Word)
	for k, words := range difficulties {
		newKey := strings.ToLower(k)
		for _, word := range words {
			mp[newKey] = append(mp[newKey], Word{
				Value:    strings.ToLower(word.Value),
				Attempts: word.Attempts,
			})
		}
	}

	return DifficultiesMap{
		mp: mp,
	}
}

func (dm *DifficultiesMap) Len() int {
	return len(dm.mp)
}

func (dm *DifficultiesMap) Get(difficulty string) ([]Word, bool) {
	val, ok := dm.mp[difficulty]
	return val, ok
}
