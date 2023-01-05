package permute

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"b", "a", "c"}
	s3 := []string{"b", "d", "c"}
	s4 := []string{"b", "a", "c", "d"}

	_, err := ComputePermutation(s1, s3)
	assert.NotNil(t, err)

	_, err = ComputePermutation(s1, s4)
	assert.NotNil(t, err)

	perm, err := ComputePermutation(s1, s2)
	assert.Nil(t, err)
	assert.Equal(t, []int{1, 0, 2}, perm)

	permuted, err := ApplyPermutation(s1, perm)
	assert.Nil(t, err)
	assert.Equal(t, permuted, s2)

	_, err = ApplyPermutation(s1, []int{0, 1, 2, 3})
	assert.NotNil(t, err)

	// invalid permutation
	_, err = ApplyPermutation(s1, []int{0, 1, 3})
	assert.NotNil(t, err)

	// identity permutation
	permuted, err = ApplyPermutation(s1, []int{0, 1, 2})
	assert.Nil(t, err)
	assert.Equal(t, permuted, s1)
}
