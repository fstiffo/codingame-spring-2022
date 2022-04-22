package main

// Copy a map of Entity
func CopyMap(m map[int]Entity) map[int]Entity {
	newMap := make(map[int]Entity)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
