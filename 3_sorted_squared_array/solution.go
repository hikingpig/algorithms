package solution

func square(x int) int {
	return x * x
}

// the negative number can be the smallest one but can create the largest square
// so we start from 2 ends of an array and choose the largest square from there
// and put that into the result array
// after the result is filled, all elements will be visited
func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))
	minIdx := 0
	maxIdx := len(array) - 1
	for i := len(array) - 1; i >= 0; i-- {
		maxSquare1 := square(array[minIdx])
		maxSquare2 := square(array[maxIdx])
		if maxSquare1 > maxSquare2 {
			sortedSquares[i] = maxSquare1
			minIdx++
		} else {
			sortedSquares[i] = maxSquare2
			maxIdx--
		}
	}
	return sortedSquares
}
