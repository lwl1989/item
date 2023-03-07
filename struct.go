package item

import "reflect"

// GetItemMap get struct item with fieldK
func GetItemMap[K comparable, V any, Item any](values []Item, fieldK, fieldV string) map[K]V {
	if len(values) == 0 {
		return nil
	}

	res := make(map[K]V)
	for _, v := range values {
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Pointer {
			if rv.IsNil() {
				continue
			}
			rv = rv.Elem()
		}

		rk := rv.FieldByName(fieldK)
		k, ok := rk.Interface().(K)
		if !ok {
			continue
		}

		rv1 := rv.FieldByName(fieldV)
		value, ok := rv1.Interface().(V)
		if !ok {
			continue
		}
		res[k] = value
	}

	return res
}
