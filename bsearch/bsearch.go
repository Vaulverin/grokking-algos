package bsearch

func Search(arr []int, searchVal int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (high + low) / 2
		if arr[mid] < searchVal {
			low = mid + 1
		} else if arr[mid] > searchVal {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
