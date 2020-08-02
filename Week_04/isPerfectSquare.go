func isPerfectSquare(num int) bool {

	if num < 2 {
		return true
	}
	
	left,right := 0,num/2 + 1
	for left <= right {
		mid := (left+right)/2
		if mid * mid == num {
			return true
		} else if mid * mid < num {
			left = mid + 1
		}else {
			right = mid - 1
		}
	}

	return false
}