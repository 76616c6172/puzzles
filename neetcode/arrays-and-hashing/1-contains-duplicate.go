package neetcode

func Contains_duplicate(nums []int) bool {
	numIsPresent := make(map[int]bool)
	for _, n := range nums {
		if numIsPresent[n] {
			return true
		}
		numIsPresent[n] = true
	}

	return false
}
