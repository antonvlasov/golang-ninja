package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AddOrRemove(aPtr *[]int, elem int) {
	count := 0
	for _, val := range *aPtr {
		if val == elem {
			count += 1
		}
	}

	if count == 0 {
		*aPtr = append(*aPtr, elem)
	} else {
		for idx, val := range *aPtr {
			if val == elem {
				*aPtr = append((*aPtr)[:idx], (*aPtr)[idx+1:]...)
			}
		}
	}
}

func TestAddOrRemove(t *testing.T) {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 20; i++ {
		AddOrRemove(&a, i)
	}

	expected := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	require.Equal(t, expected, a)
}
