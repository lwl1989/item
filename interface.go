package item


type ICompare interface {
	Compare(ICompare) bool
	GetCompareValue() interface{}
}
