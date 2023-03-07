package item

func GetArrayIndex[T comparable](params []T, value T) int {
	for index, param := range params {
		if param == value {
			return index
		}
	}
	return -1
}

func InArray[T comparable](params []T, value T) bool {
	for _, param := range params {
		if param == value {
			return true
		}
	}
	return false
}

func InArrayCompare[T ICompare](params []T, value T) bool {
	for _, param := range params {
		if param.Compare(value) {
			return true
		}
	}
	return false
}

func ArrayMap(params []map[string]interface{}, key string) map[string]map[string]interface{} {
	if len(params) < 1 {
		return nil
	}
	res := make(map[string]map[string]interface{})
	for _, v := range params {
		_, ok := v[key]
		if !ok {
			continue
		}
		res[key] = v
	}
	return res
}

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

func ArrayColumns[K comparable, V any](params []map[K]V, key K) []V {
	if len(params) < 1 {
		return nil
	}
	res := make([]V, len(params))
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

func ArrayColumnValues[K comparable, V any](params []map[K]interface{}, key, VKey K) map[K]V {
	if len(params) < 1 {
		return nil
	}
	res := make(map[K]V)
	for _, m := range params {
		_, ok := m[key]
		if !ok {
			// res[i]= *new(V) do not create it
			continue
		}
		v, ok := m[VKey]
		if !ok {
			continue
		}
		value, ok := v.(V)
		if ok {
			res[key] = value
		}
	}
	return res
}

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
