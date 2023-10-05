package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AddOrRemove(aPtr *[]int, elem int) {
    f := false
    arr := aPtr
    var del_idx []int 
    for idx, value := range *arr{
        if value == elem{
            del_idx = append(del_idx, idx)
            f = true
        }
    }
    if !f{
        *arr = append(*arr, elem)
    }else{
        count := 0
        for _, value := range del_idx{
            *arr = append((*arr)[:value - count], (*arr)[value + 1 - count:]...)
            count += 1
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
