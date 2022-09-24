package sorts

// BubbleSort sorts numbers
func BubbleSort(numbers *[]int) {
	var swapped bool = true
	var nofIteration int
	for swapped {
		swapped = false
		for j := 0; j < len(*numbers)-nofIteration-1; j++ {
			if (*numbers)[j] > (*numbers)[j+1] {
				(*numbers)[j], (*numbers)[j+1] = (*numbers)[j+1], (*numbers)[j]
				swapped = true
			}
		}
		nofIteration += 1
	}
}
