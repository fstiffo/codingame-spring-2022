package main

// Copy a map of Entity
func CopyMap(m map[int]Common) map[int]Common {
	newMap := make(map[int]Common)
	for k, v := range m {
		newMap[k] = v
	}
	return newMap
}
