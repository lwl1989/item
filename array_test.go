package item

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	fmt.Println(InArray[int]([]int{1, 2, 3, 45}, 3))
	fmt.Println(InArray[int]([]int{1, 2, 3, 45}, 44))
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
	fmt.Println(InArrayCompare[cp]([]cp{{
		"1",
	}, {"2"}}, c))
	fmt.Println(InArrayCompare[cp]([]cp{{
		"2",
	}, {"3"}}, c))
}

func TestArrayMapCompare(t *testing.T) {
	fmt.Println(ArrayMapCompare[int]([]map[int]interface{}{
		{3: "4", 4: 5},
		{4: "3"},
	}, 3))
	fmt.Println(ArrayMapCompare[int]([]map[int]interface{}{
		{4: "3"},
	}, 3))
}

func TestArrayMapCompareValue(t *testing.T) {
	fmt.Println(ArrayMapCompareValue[string, ICompare]([]map[string]ICompare{
		{"a": cp{value: "dasdas"}},
		{"b": cp{value: "dasdas"}},
	}, "a"))
}

func TestArrayUnique(t *testing.T) {
	fmt.Println(ArrayUnique[string]([]string{"a", "aaa", "a", "bbb", "ccc", "a"}))
}

func TestArrayDiff(t *testing.T) {
	fmt.Println(ArrayDiff[string]([]string{"a", "aaa", "a", "bbb", "ccc", "a", "ggg"}, []string{"a", "aaa", "a", "bbb", "ccc", "a", "fff"}))
}
