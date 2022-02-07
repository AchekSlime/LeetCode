package Algorithms

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version iAs bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return true
}

func FirstBadVersion(n int) int {
	l, r := 1, n

	for l < r {
		m := (l + r) / 2
		if ans := checkThrees(&l, &r, m); ans != -1 {
			return ans
		}
	}

	return l
}

func checkThrees(l, r *int, m int) int {
	center := isBadVersion(m)
	var left bool
	if m == 1 {
		left = false
	} else {
		left = isBadVersion(m - 1)
	}
	right := isBadVersion(m + 1)

	if center {
		if !left {
			return m
		} else {
			*r = m - 1
			return -1
		}
	} else {
		if right {
			return m + 1
		} else {
			*l = m + 1
			return -1
		}
	}

}
