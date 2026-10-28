package binarysearch

func SearchInts(slice []int, key int) int {
	return binarySearch(slice, 0, len(slice)-1, key)
}

func binarySearch(slice []int, left, right, key int) int {
	if left > right {
		return -1
	}

	middleIndex := (left + right) / 2
	if slice[middleIndex] == key {
		return middleIndex
	}

	if slice[middleIndex] > key {
		return binarySearch(slice, left, middleIndex-1, key)
	}

	if slice[middleIndex] < key {
		return binarySearch(slice, middleIndex+1, right, key)
	}

	return -1
}
