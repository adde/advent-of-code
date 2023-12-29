package utils

import "sort"

type OrderedMap struct {
	Keys   []string
	Values map[string]int
}

func (om *OrderedMap) Set(key string, value int) {
	if _, exists := om.Values[key]; !exists {
		om.Keys = append(om.Keys, key)
	}
	om.Values[key] = value
}

func (om *OrderedMap) Get(key string) (int, bool) {
	value, exists := om.Values[key]
	return value, exists
}

func (om *OrderedMap) Delete(key string) {
	for i, k := range om.Keys {
		if k == key {
			om.Keys = append(om.Keys[:i], om.Keys[i+1:]...)
			break
		}
	}
	delete(om.Values, key)
}

func (om *OrderedMap) Sort(isDesc bool) {
	if isDesc {
		sort.SliceStable(om.Keys, func(i, j int) bool {
			if om.Values[om.Keys[i]] > om.Values[om.Keys[j]] {
				return true
			}
			return false
		})
	} else {
		sort.SliceStable(om.Keys, func(i, j int) bool {
			if om.Values[om.Keys[i]] < om.Values[om.Keys[j]] {
				return true
			}
			return false
		})
	}
}
