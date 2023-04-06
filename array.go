package item

// GetArrayIndex get value index with slice, if not found return -1
// 获取数据下标，不存在返回-1
func GetArrayIndex[T comparable](params []T, value T) int {
	for index, param := range params {
		if param == value {
			return index
		}
	}
	return -1
}

// InArray
// 判断数据是否存在 compare
func InArray[T comparable](params []T, value T) bool {
	for _, param := range params {
		if param == value {
			return true
		}
	}
	return false
}

// InArrayCompare
// 判断数据是否存在 ICompare
func InArrayCompare[T ICompare](params []T, value T) bool {
	for _, param := range params {
		if param.Compare(value) {
			return true
		}
	}
	return false
}

// ArrayMap
// 返回输入数组中某个单一列的值的map string interface
func ArrayMap[V any](params []map[string]V, key string) map[string][]map[string]V {
	if len(params) < 1 {
		return nil
	}
	res := make(map[string][]map[string]V)
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = append(res[key], v)
	}
	return res
}

// ArrayMapCompare
// 返回输入数组中某个单一列的值的map interface
func ArrayMapCompare[T comparable](params []map[T]interface{}, key T) map[T]map[T]interface{} {
	if len(params) < 1 {
		return nil
	}
	res := make(map[T]map[T]interface{})
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = v
	}
	return res
}

// ArrayMapCompareValue
// 返回输入数组中某个单一列的值的map
func ArrayMapCompareValue[K comparable, V any](params []map[K]V, key K) map[K]map[K]V {
	if len(params) < 1 {
		return nil
	}
	res := make(map[K]map[K]V)
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = v
	}
	return res
}

// ArrayColumns
// 返回输入数组中某个单一列的值
func ArrayColumns[K comparable, V any](params []map[K]V, key K) []V {
	if len(params) < 1 {
		return nil
	}
	res := make([]V, len(params), len(params))
	for i, m := range params {
		v, ok := m[key]
		if !ok {
			// res[i]= *new(V) do not create it
			continue
		}
		res[i] = v
	}
	return res
}

// ArrayUnique unique with slice
// 对切片进行去重
func ArrayUnique[V comparable](params []V) []V {
	mp := map[V]struct{}{}
	res := make([]V, len(params))
	var index int
	for _, v := range params {
		if _, ok := mp[v]; ok {
			continue
		}
		mp[v] = struct{}{}
		res[index] = v
		index++
	}
	return res[:index]
}

// ArrayDiff diff with slice values
// 取多个切片的差集
func ArrayDiff[V comparable](params ...[]V) []V {
	var all []V
	mp := map[V]int8{}
	for _, v := range params {
		all = append(all, v...)
	}
	for _, v := range all {
		if num, ok := mp[v]; ok {
			mp[v] = num + 1
			continue
		}
		mp[v] = 1
	}
	var res []V
	for v, num := range mp {
		if num == 1 {
			res = append(res, v)
		}
	}
	return res
}

// ArraySub diff with slice values
// example:  [1,2,3]  [2,3,4]  array_sub(a,b) => [1]
func ArraySub[V comparable](arr1, arr2 []V) []V {
	var res []V
	for _, v1 := range arr1 {
		found := false
		for _, v2 := range arr2 {
			if v1 == v2 {
				found = true
			}
		}
		if !found {
			res = append(res, v1)
		}
	}
	return res
}

// ArrayIntersect Intersect with slice values
// 取多个切片的交集  params max 127
func ArrayIntersect[V comparable](params ...[]V) []V {
	var all []V
	mp := map[V]int8{}
	l := int8(len(params))
	for _, v := range params {
		v = ArrayUnique[V](v)
		all = append(all, v...)
	}
	for _, v := range all {
		if num, ok := mp[v]; ok {
			mp[v] = num + 1
			continue
		}
		mp[v] = 1
	}
	var res []V
	for v, num := range mp {
		if num == l {
			res = append(res, v)
		}
	}
	return res
}

// ArrayValues  map transfer to slice
func ArrayValues[K comparable, V any](mp map[K]V) []V {
	res := make([]V, len(mp), len(mp))
	var i int
	for _, v := range mp {
		res[i] = v
		i++
	}
	return res
}
