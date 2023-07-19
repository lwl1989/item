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

func TestArrayValues(t *testing.T) {
	mp := map[string]cp{
		"a": cp{
			value: "a",
		},
		"b": cp{
			value: "b",
		},
	}
	res := ArrayValues[string, cp](mp)
	if len(res) != 2 {
		t.Fatal("res error with len")
	}
}

func TestArrayIntersect(t *testing.T) {
	fmt.Println(ArrayIntersect[string]([]string{"a", "aaa", "a", "bbb", "ccc", "a", "ggg"}, []string{"a", "aaa", "a", "bbb", "ccc", "a", "fff"}))
}

func TestArraySub(t *testing.T) {
	fmt.Println(ArraySub[string]([]string{"a", "aaa", "a", "bbb", "ccc", "a", "ggg"}, []string{"a", "aaa", "a", "bbb", "ccc", "a", "fff"}))
}

func TestArrayColumns(t *testing.T) {
	fmt.Println(ArrayColumns[string, string]([]map[string]string{{"a": "a"}, {"a": "b"}, {"a": "c"}, {"a": "d"}}, "a"))
}

func TestArrayMap(t *testing.T) {
	fmt.Println(ArrayMap([]map[string]string{{"a": "a"}, {"a": "b"}, {"a": "c"}, {"a": "d"}}, "a"))
}

func TestGetArrayIndex(t *testing.T) {
	fmt.Println(GetArrayIndex[int]([]int{1, 3, 5, 7, 98}, 5))
	fmt.Println(GetArrayIndex[string]([]string{"a", "aaa", "a", "bbb", "ggg", "a", "ggg"}, "ggg"))
}

func TestArrayMapColumn(t *testing.T) {
	t.Logf("%+v", ArrayMapColumn[map[string]string, string]([]map[string]string{{"a": "a"}, {"a": "b"}, {"a": "c"}, {"a": "d"}}, "a"))
	type A struct {
		A string
		B int
		C float64
	}
	t.Logf("%+v", ArrayMapColumn[A, string]([]A{
		{
			A: "3",
			B: 3,
			C: 3.3,
		},
		{
			A: "4",
			B: 4,
			C: 4.3,
		},
		{
			A: "5",
			B: 5,
			C: 5.3,
		},
	}, "A"))
	t.Logf("%+v", ArrayMapColumn[A, int]([]A{
		{
			A: "3",
			B: 3,
			C: 3.3,
		},
		{
			A: "4",
			B: 4,
			C: 4.3,
		},
		{
			A: "5",
			B: 5,
			C: 5.3,
		},
	}, "B"))
	t.Logf("%+v", ArrayMapColumn[A, float64]([]A{
		{
			A: "3",
			B: 3,
			C: 3.3,
		},
		{
			A: "4",
			B: 4,
			C: 4.3,
		},
		{
			A: "5",
			B: 5,
			C: 5.3,
		},
	}, "C"))
}
