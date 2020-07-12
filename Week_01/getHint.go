func getHint(secret string, guess string) string {

	sliceGue := strings.Split(guess, "")
	sliceSec := strings.Split(secret, "")
	aCount := 0
	bCount := 0
	//统计匹配的数并记录下标
	var idx []int
	for i := 0; i < len(sliceSec); i++ {
		if sliceGue[i] == sliceSec[i] {
			idx = append(idx, i)
			aCount++
		}
	}
	//根据匹配的下标删除指定元素(从后往前删)
	for i := len(idx) - 1; i >= 0; i-- {
		point := idx[i]
		sliceGue = append(sliceGue[0:point], sliceGue[point+1:]...)
		sliceSec = append(sliceSec[0:point], sliceSec[point+1:]...)
	}
	for _, vs := range sliceSec {
		for idxg, vg := range sliceGue {
			if vs == vg {
				//删除
				sliceGue = append(sliceGue[0:idxg], sliceGue[idxg+1:]...)
				bCount++
				break
			}
		}
	}
	return fmt.Sprintf("%vA%vB", aCount, bCount)

}