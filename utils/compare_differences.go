package utils

/*
CompareDifferences takes in two slices of uints: oldArray and newArray and returns two slices of
uints: addArr and deleteArray. addArr represents the new data passed in newArray but not in oldArray
whereas deleteArr represents the outstanding data in the oldArray which are not present in the newArray.
*/
func CompareDifferences(oldArray []uint, newArray []uint) ([]uint, []uint) {
	// Define two empty maps with uint key and bool value
	oldSet := make(map[uint]bool)
	newSet := make(map[uint]bool)

	// Convert the oldArray and newArray into maps for faster element lookup
	for _, val := range oldArray {
		oldSet[val] = true
	}

	for _, val := range newArray {
		newSet[val] = true
	}

	// Define two empty maps of length of newArray and oldArray
	addArr := make([]uint, 0, len(newArray))
	deleteArr := make([]uint, 0, len(oldArray))

	// Iterate over the elements in newArray to check if they are present in
	// oldArray or not. If not, it is added to addArr.
	for _, val := range newArray {
		if !oldSet[val] {
			addArr = append(addArr, val)
		}
	}

	// Iterate over the elements in oldArray to check if they are present in
	// newArray or not. If not, it is added to deleteArr.
	for _, val := range oldArray {
		if !newSet[val] {
			deleteArr = append(deleteArr, val)
		}
	}

	return addArr, deleteArr
}
