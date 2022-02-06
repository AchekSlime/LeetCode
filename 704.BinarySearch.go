package LeetCode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}

	if nums[l] == target {
		return l
	} else {
		return -1
	}
}
