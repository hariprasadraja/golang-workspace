package others

// Educational purpose only, extremley poor performance

type data []int32

func (dataList data) PancakeSort() {
	for uns := len(dataList) - 1; uns > 0; uns-- {
		// find largest in unsorted range
		lx, lg := 0, dataList[0]
		for i := 1; i <= uns; i++ {
			if dataList[i] > lg {
				lx, lg = i, dataList[i]
			}
		}
		// move to final position in two flips
		dataList.flip(lx)
		dataList.flip(uns)
	}
}

func (dataList data) flip(r int) {
	for l := 0; l < r; l, r = l+1, r-1 {
		dataList[l], dataList[r] = dataList[r], dataList[l]
	}
}
