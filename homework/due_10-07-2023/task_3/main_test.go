package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AddOrRemove(aPtr *[]int, elem int) {
	res := make([]int, 0)
	flag := false
	for _, val := range *aPtr {
		if val == elem {
			flag = true
		} else {
			res = append(res, val)
		}
	}
	if !flag {
		res = append(res, elem)
	}
	*aPtr = res
}

func TestAddOrRemove(t *testing.T) {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 20; i++ {
		AddOrRemove(&a, i)
	}

	expected := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	require.Equal(t, expected, a)
}
