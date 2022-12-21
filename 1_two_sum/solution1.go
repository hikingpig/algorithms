package solution

// TwoSum1 do 2 things in each iteration
// - build the map - search for match
// because the index is always advanced, we don't have to
// worry if target and potential match have the same value
// and at the same position
func TwoSum1(array []int, target int) []int {
	m := map[int]bool{}
	for _, x := range array {
		y := target - x
		if _, ok := m[y]; ok {
			return []int{x, y}
		}
		m[x] = true
	}
	return []int{}
}

// TwoSum2 build the map first and search for the match
// There's 1 more loop compare to TwoSum1
// And in case the match is the same value, then we have
// to check if the position is the same
// No as elegant as TwoSum1
func TwoSum2(array []int, target int) []int {
	m := map[int]int{}
	for _, x := range array {
		m[x]++
	}
	for _, x := range array {
		y := target - x
		if y != x {
			if _, ok := m[y]; ok {
				return []int{x, y}
			}
		} else {
			if m[x] > 2 {
				return []int{x, y}
			}
		}
	}
	return []int{}

}
