
func fourSum(nums []int,target int) [][]int {

	if(len(nums) < 4) {
		return nil
	}

	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums) - 3; i++ {
		
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1; j < len(nums) - 2; j++ {
			
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			
			k ,l := j+1,len(nums) - 1
			
			for k < l {
				
				distinct := target - (nums[i]+nums[j]+nums[k]+nums[l])
				if( distinct == 0 ){
					res = append(res,[]int{nums[i],nums[j],nums[k],nums[l]})
				}
				if distinct > 0 {
					for cur := l; k < l && nums[cur] == nums[l]; l-- {

					}
				} else {
					for cur := k; k < l && nums[cur] == nums[k]; k++ {

					}
				}
			}



		}

	}
	return res
}