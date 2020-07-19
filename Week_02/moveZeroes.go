func moveZeroes(nums []int) {
	
	l := len(nums) - 1 
	i := 0
	
	for {

		if i >= l {
			break
		} else {	
			if nums[i] == 0 {		
				nums = append(nums[0:i],nums[i+1:]...)
				nums = append(nums,0)
				l = l - 1			
			} else {
				i++
			}
		}

	}

	return 
	
}