package Algorithms

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := (l + r) / 2
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	if target > nums[l] {
		return l + 1
	} else {
		return l
	}
}
