package Algorithms

func TwoSum(numbers []int, target int) []int {
	for i := 0; numbers[i] <= target/2 && i < len(numbers); i++ {
		sIndex, sValue := binSearch(numbers[i+1:], target-numbers[i])
		sIndex += i + 1

		if numbers[i]+sValue == target {
			return []int{i + 1, sIndex + 1}
		}
	}
	return []int{0, 0}
}

func binSearch(nums []int, target int) (index int, value int) {
	l, r := 0, len(nums)-1
	for l < r {
		pivot := (l + r) / 2
		if nums[pivot] < target {
			l = pivot + 1
		} else if nums[pivot] > target {
			r = pivot
		} else {
			return pivot, nums[pivot]
		}
	}

	return l, nums[l]
}
