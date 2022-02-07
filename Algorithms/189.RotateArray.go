package Algorithms

func Rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	d := n - k
	if k == 0 {
		return
	}
	tmp := make([]int, 2*n)
	for i := range nums {
		tmp[i] = nums[i]
		tmp[i+n] = nums[i]
	}

	for i := range nums {
		nums[i] = tmp[i+d]
	}
}
