package solution

// Loop over the sequence
// For each element of the sequence
// Loop over the array to find that element
// Keep the index of the last matching element in the array
// and start search from that index + 1
func IsSubsequence1(array []int, sequence []int) bool {
	arrayStartIdx := 0
	for _, s := range sequence {
		// i must be checked outside the inner loop
		i := arrayStartIdx
		for ; i < len(array); i++ {
			// found matching element
			// update the start index to search array for next
			// sequence element
			if s == array[i] {
				arrayStartIdx = i + 1
				break
			}
		}
		// matching element not found.
		if i == len(array) {
			return false
		}
	}
	return true
}

// IsSubsequence2 applied the same logic as IsSubsequence1
// but in a more elegant way
func IsSubsequence2(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0
	// looping over both array and sequence at the same time
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			// found matching element, search for next element
			seqIdx++
		}
		// search for sequence element in the array
		arrIdx++
	}
	// if all the sequence elements are searched, seqIdx will be at the end
	// else only arrIdx is at the end
	return seqIdx == len(sequence)
}

// IsSubsequence3 search by looping over array
func IsSubsequence3(array []int, sequence []int) bool {
	seqIdx := 0
	for _, v := range array {
		// search all sequence elements, success
		if seqIdx == len(sequence) {
			break
		}
		if v == sequence[seqIdx] {
			seqIdx++
		}
	}
	// finish looping over array. check if all element in sequence searched and found
	return seqIdx == len(sequence)
}
