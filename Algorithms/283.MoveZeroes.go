package Algorithms

func moveZeroes(nums []int) {
	last := 0
	for _, v := range nums {
		if v != 0 {
			nums[last] = v
			last++
		}
	}
	for last := last; last < len(nums); last++ {
		nums[last] = 0
	}
}
