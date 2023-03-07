package item

func MapKeys[T any](mp map[string]T) (keys []string) {
	if len(mp) == 0 {
		return
	}
	for k, _ := range mp {
		keys = append(keys, k)
	}
	return keys
}

func MapGetKey[T comparable](mp map[string]T, item T) string {
	if len(mp) == 0 {
		return ""
	}
	for k, v := range mp {
		if v == item {
			return k
		}
	}
	return ""
}

func MapGetKeyWithCompare[T ICompare](mp map[string]T, item T) string {
	if len(mp) == 0 {
		return ""
	}
	for k, v := range mp {
		if item.Compare(v) {
			return k
		}
	}
	return ""
}
