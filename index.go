// Package index provides high level generic like functions for simple slices.
package index

import (
	"reflect"

	"github.com/pkg/errors"
)

// GetAll function allows you to check on which indexes your item
// is positioned in the slice. First argument is the slice in which
// you are searching for your item and the second one is the item.
func GetAll(givenSlice interface{}, givenItem interface{}) (indexes []int, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("index: incorrect usage of the package")
		}
	}()

	slice, err := getSliceFromInterface(givenSlice)
	if err != nil {
		return []int{}, err
	}

	indexes = []int{}
	for i, item := range slice {
		if err := checkItemComplexity(item); err != nil {
			return indexes, err
		}
		if item == givenItem {
			indexes = append(indexes, i)
		}
	}

	return indexes, nil
}

func getSliceFromInterface(sliceInterface interface{}) (slice []interface{}, err error) {
	value := reflect.ValueOf(sliceInterface)
	if value.Kind() != reflect.Slice {
		return nil, errors.New("index: given non-slice type")
	}

	slice = make([]interface{}, value.Len())
	for i := 0; i < value.Len(); i++ {
		slice[i] = value.Index(i).Interface()
	}

	return slice, nil
}

func checkItemComplexity(item interface{}) (err error) {
	value := reflect.ValueOf(item)
	if value.Kind() == reflect.Slice {
		return errors.New("index: given multidimensional slice")
	}
	if value.Kind() == reflect.Map {
		return errors.New("index: given too complex slice of maps")
	}

	return nil
}

// IsAny function allows you to check if there is at least one item
// in the slice that you gave. First argument is the slice in which
// you are searching for your item and the second one is the item.
func IsAny(slice interface{}, item interface{}) (isAny bool, err error) {
	indexes, err := GetAll(slice, item)
	if err != nil {
		return false, err
	}

	return len(indexes) > 0, nil
}

// HowMany function allows you to check how many items that you gave
// exist in the slice. First argument is the slice in which
// you are searching for your items and the second one is the item.
func HowMany(slice interface{}, item interface{}) (howMany int, err error) {
	indexes, err := GetAll(slice, item)
	if err != nil {
		return 0, err
	}

	return len(indexes), nil
}
