// Package index provides high level generic like functions for simple slices.
package index

import (
	"reflect"

	"github.com/pkg/errors"
)

// GetAll function allows you to check on which indexes your item
// is positioned in the slice. First argument is the slice in which
// you are searching for your item and the second one is the item.
func GetAll(_slice interface{}, _item interface{}) (indexes []int, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("index: incorrect usage of the package")
		}
	}()

	slice, err := sliceFromInterface(_slice)
	if err != nil {
		return []int{}, err
	}

	indexes = []int{}
	for index, item := range slice {
		if err := checkItemComplexity(item); err != nil {
			return []int{}, err
		}
		if item == _item {
			indexes = append(indexes, index)
		}
	}

	return indexes, nil
}

func sliceFromInterface(sliceInterface interface{}) (ret []interface{}, err error) {
	slice := reflect.ValueOf(sliceInterface)
	if slice.Kind() != reflect.Slice {
		return nil, errors.New("index: given non-slice type")
	}

	ret = make([]interface{}, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		ret[i] = slice.Index(i).Interface()
	}

	return ret, nil
}

func checkItemComplexity(_item interface{}) (err error) {
	item := reflect.ValueOf(_item)
	if item.Kind() == reflect.Slice {
		return errors.New("index: given multidimensional slice")
	}
	if item.Kind() == reflect.Map {
		return errors.New("index: given too complex slice of maps")
	}

	return nil
}

// IsAny function allows you to check if there is at least one item
// in the slice that you gave. First argument is the slice in which
// you are searching for your item and the second one is the item.
func IsAny(_slice interface{}, _item interface{}) (isAny bool, err error) {
	indexes, err := GetAll(_slice, _item)
	if err != nil {
		return false, err
	}

	return len(indexes) > 0, nil
}

// HowMany function allows you to check how many items that you gave
// exist in the slice. First argument is the slice in which
// you are searching for your items and the second one is the item.
func HowMany(_slice interface{}, _item interface{}) (howMany int, err error) {
	indexes, err := GetAll(_slice, _item)
	if err != nil {
		return 0, err
	}

	return len(indexes), nil
}
