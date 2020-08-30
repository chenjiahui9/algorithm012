func judgeSquareSum(c int) bool {
	for a := 0; a * a <= c; a++ {
		b := c - a*a
		if binary_search(0,b,b) {
			return true
		}
	}
	return false
}

func binary_search(s,e int,n int) bool {
	if s > e {
		return false
	}
	mid := (s + e)/2
	if mid * mid  == n {
		return true
	}
	if mid*mid > n {
		return binary_search(s,mid - 1,n)
	}
	
	return binary_search(mid+1,e,n)
}