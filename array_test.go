package item

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	fmt.Println(InArray([]int{1, 2, 3, 45}, 3))
	fmt.Println(InArray([]int{1, 2, 3, 45}, 44))
}

type cp struct {
	value string
}

func (cp cp) Compare(compare ICompare) bool {
	return cp.GetCompareValue() == compare.GetCompareValue()
}
func (cp cp) GetCompareValue() interface{} {
	return cp.value
}

func TestInArrayCompare(t *testing.T) {
	c := cp{
		value: "1",
	}
	fmt.Println(InArrayCompare([]cp{{
		"1",
	}, {"2"}}, c))
	fmt.Println(InArrayCompare([]cp{{
		"2",
	}, {"3"}}, c))
}