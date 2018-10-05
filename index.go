package index

import (
	"github.com/pkg/errors"
	"reflect"
)

func GetAll(_slice interface{}, _item interface{}) (indexes []int, err error) {
	indexes = []int{}

	slice, err := sliceFromInterface(_slice)
	if err != nil {
		return indexes, err
	}

	for index, item := range slice {
		if item == _item {
			indexes = append(indexes, index)
		}
	}

	return indexes, nil
}

func sliceFromInterface(sliceInterface interface{}) (r0 []interface{}, err error) {
	slice := reflect.ValueOf(sliceInterface)
	if slice.Kind() != reflect.Slice {
		return nil, errors.New("index: given non-slice type")
	}

	r0 = make([]interface{}, slice.Len())

	for i := 0; i < slice.Len(); i++ {
		r0[i] = slice.Index(i).Interface()
	}

	return r0, nil
}
