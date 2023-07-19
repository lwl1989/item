package item

import "errors"

func Catch(r interface{}) (err error) {
	//check exactly what the panic was and create error.
	switch x := r.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = errors.New("unKnow panic")
	}
	return err
}
