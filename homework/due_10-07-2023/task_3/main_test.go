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
	switch count {
	case 0:
		*aPtr = append(*aPtr, elem)
	case 1:
		index := 0
		for idx, val := range *aPtr {
			if val == elem {
				index = idx
			}
		}
		*aPtr = append((*aPtr)[:index], (*aPtr)[index+1:]...)

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
