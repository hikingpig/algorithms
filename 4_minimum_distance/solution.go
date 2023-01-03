package solution

import "sort"

func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	// store the indices
	idx1, idx2 := 0, 0
	// store the pair of numbers
	first, second := array1[idx1], array2[idx2]
	// store the current distance
	dist := abs(first, second)
	// store the minimum distance so far
	minDist := dist

	for idx1 < len(array1) && idx2 < len(array2) {
		x1, x2 := array1[idx1], array2[idx2]
		if x1 > x2 {
			dist = x1 - x2
			idx2++ // move idx2 up so the distance can be reduced
		} else if x2 > x1 {
			dist = x2 - x1
			idx1++ // move idx1 up so the distance can be reduced
		} else { // x1 == x2, return them!
			return []int{x1, x2}
		}
		if dist < minDist { // find a new minimum distance
			minDist = dist
			first, second = x1, x2
		}
	}
	return []int{first, second}
}
