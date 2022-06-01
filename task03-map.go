/*
Task 3: Maps
function that returns map values sorted in order of increasing keys.
Input -> {2: "a", 0: "b", 1: "c"}
Output -> ["b", "c", "a"]
Input -> {10: "aa", 0: "bb", 500: "cc"}
Output -> ["bb", "aa", "cc"]
*/

package homework

import (
	"sort"
)

// type ByInt map[int]string

// func (a ByInt) Len() int           { return len(a) }
// func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a ByInt) Less(i, j int) bool { return i > j }

func sortMapValues(input map[int]string) (result []string) {

	// sort.Sort(ByInt(input))
	// return
	var keys = make([]int, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, value := range keys {
		result = append(result, input[value])
	}
	return result
}
