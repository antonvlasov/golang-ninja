package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func AddOrRemove(aPtr *[]int, elem int) {
	ans := make([]int, 0)
	flag := false
	for _, v := range *aPtr {
		if v == elem {
			//	fmt.Println(aPtr[index], index, v, elem) объясните пожалуйста почему в го нельзя обращаться к указателю по индексу я пробовал и так aPtr[index] и так *aPtr[index]
			// upd понял ошибку, оказывается надо было к указателю по индексу обращаться (*aPtr)[index], а я не брал в скобки
			flag = true
			//*aPtr = append(*aPtr[:index], *aPtr[index+1:]...) так тоже не получилось из-за обращения по индексу
		} else {
			ans = append(ans, v)

		}
	}
	if !flag {
		ans = append(ans, elem)
	}
	*aPtr = ans
}

func TestAddOrRemove(t *testing.T) {
	a := []int{1, 2, 3, 4}

	for i := 0; i < 20; i++ {
		AddOrRemove(&a, i)
	}

	expected := []int{0, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	require.Equal(t, expected, a)
}
