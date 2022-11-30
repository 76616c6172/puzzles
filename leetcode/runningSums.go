package main

func runningSums(nums []int) []int {
	runningSum := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		numberToAppend := runningSum[i-1] + nums[i]
		runningSum = append(runningSum, numberToAppend)
	}

	return runningSum
}
