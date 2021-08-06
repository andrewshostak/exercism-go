package binarysearch

func SearchInts(slice []int, key int) int {
	return binarySearch(slice, 0, len(slice)-1, key)
}

func binarySearch(slice []int, left, right, key int) int {
	if right >= left {
		middle := left + (right-left)/2
		if slice[middle] == key {
			return middle
		}

		if slice[middle] > key {
			return binarySearch(slice, left, middle-1, key)
		}

		return binarySearch(slice, middle+1, right, key)
	}

	return -1
}
