package day20

func Mix(input []int, indices []int, multiplicator int) ([]int, []int) {
	copiedInput := make([]int, len(input))
	copy(copiedInput, input)

	copiedIndices := make([]int, len(indices))
	copy(copiedIndices, indices)

	for i := 0; i < len(copiedInput); i++ {
		from := 0
		for from < len(copiedIndices) {
			if copiedIndices[from] == i {
				break
			}
			from += 1
		}

		to := from + (copiedInput[from]*multiplicator)%(len(copiedInput)-1)
		for to <= 0 {
			to += len(copiedInput) - 1
		}
		for to >= len(copiedInput) {
			to -= len(copiedInput) - 1
		}

		n := copiedInput[from]
		nidx := copiedIndices[from]
		if to > from {
			copy(copiedInput[from:], copiedInput[from+1:to+1])
			copy(copiedIndices[from:], copiedIndices[from+1:to+1])
		} else {
			copy(copiedInput[to+1:], copiedInput[to:from])
			copy(copiedIndices[to+1:], copiedIndices[to:from])
		}
		copiedInput[to] = n
		copiedIndices[to] = nidx
	}

	return copiedInput, copiedIndices
}

func Sum(mixedInput []int, decryptionKey int) int {
	for j := 0; j < len(mixedInput); j++ {
		if mixedInput[j] == 0 {
			return mixedInput[(j+1000)%len(mixedInput)]*decryptionKey +
				mixedInput[(j+2000)%len(mixedInput)]*decryptionKey +
				mixedInput[(j+3000)%len(mixedInput)]*decryptionKey
		}
	}
	return 0
}
