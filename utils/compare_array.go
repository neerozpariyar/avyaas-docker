package utils

// this function returns bool if all the value of one array exist in another irrevelant of the succession they are in
func CompareArrayElements(arr1, arr2 []uint) bool {
	elementMap := make(map[uint]bool)

	// Add all elements of arr2 to the map
	for _, num := range arr2 {
		elementMap[num] = true
	}

	// Check if every element of arr1 is in the map
	for _, num := range arr1 {
		if !elementMap[num] {
			return false // If any element is missing, return false
		}
	}

	return true // All elements are present
}
