package Algorithms

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1
	i := r
	leftValue := nums[l] * nums[l]
	rightValue := nums[r] * nums[r]
	ansSlice := make([]int, r+1)

	for l < r {
		if leftValue > rightValue {
			ansSlice[i] = leftValue
			l++
			leftValue = nums[l] * nums[l]
		} else {
			ansSlice[i] = rightValue
			r--
			rightValue = nums[r] * nums[r]
		}
		i--
	}
	ansSlice[i] = leftValue

	return ansSlice
}
