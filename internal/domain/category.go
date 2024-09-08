package domain

import (
	"strings"
)

type CategoriesMap struct {
	mp map[string]DifficultiesMap
}

func NewCategoriesMap(categories map[string]map[string][]Word) CategoriesMap {
	mp := make(map[string]DifficultiesMap)

	for k, v := range categories {
		newVal := NewDifficultiesMap(v)
		if newVal.Len() != 0 {
			mp[strings.ToLower(k)] = newVal
		}
	}

	return CategoriesMap{
		mp: mp,
	}
}

func (cm *CategoriesMap) Len() int {
	return len(cm.mp)
}

func (cm *CategoriesMap) Get(category string) (DifficultiesMap, bool) {
	val, ok := cm.mp[category]
	return val, ok
}
