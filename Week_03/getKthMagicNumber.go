func getKthMagicNumber(k int) int {

	ans,a,b,c := []int{1},0,0,0

	for i:=1; i < k; i++ {
		ans = append(ans,min(ans[a]*3,ans[b]*5,ans[c]*7))
		if ans[i] == ans[a]*3 {
			a++
		}
		if ans[i] == ans[b]*5 {
			b++
		}
		if ans[i] == ans[c]*7 {
			c++
		}
	}

	return ans[k-1]
}

func min(x,y,z int) int {
	r := x
	if y < r {
		r = y
	} 
	
	if z < r {
		r = z
	}

	return r
}