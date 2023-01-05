package permute

import (
	"fmt"
	"sort"
)

// ComputePermutation returns the permutation that transforms
// a to b; error if such a permutation does not exist.
func ComputePermutation[T comparable](a, b []T) ([]int, error) {
	empty := []int{}
	if len(a) != len(b) {
		return empty, fmt.Errorf("lengths do not match")
	}
	elem2IndexA := make(map[T]int)
	for i, elem := range a {
		elem2IndexA[elem] = i
	}
	var perm []int
	for _, elem := range b {
		if _, ok := elem2IndexA[elem]; !ok {
			return empty, fmt.Errorf("slices contain different elements")
		}
		perm = append(perm, elem2IndexA[elem])
	}
	return perm, nil
}

// ApplyPermutation returns a slice that is obtained by permuting
// a according to the given permutation.
func ApplyPermutation[T comparable](a []T, permutation []int) ([]T, error) {
	res := []T{}
	if len(a) != len(permutation) {
		return res, fmt.Errorf("input and permutation have different lengths")
	}
	if !IsPermutation(permutation) {
		return res, fmt.Errorf("invalid permutation")
	}

	for _, idx := range permutation {
		res = append(res, a[idx])
	}
	return res, nil
}

// IsPermutation returns true if permutation is a valid permutation,
// i.e., non repeating values in [0, len-1]
func IsPermutation(permutation []int) bool {
	var copyPerm []int
	copyPerm = append(copyPerm, permutation...)
	sort.Ints(copyPerm)
	if copyPerm[0] != 0 || copyPerm[len(copyPerm)-1] != len(copyPerm)-1 {
		return false
	}
	for i := 1; i < len(copyPerm); i++ {
		if copyPerm[i-1] == copyPerm[i] {
			return false
		}
	}
	return true
}
