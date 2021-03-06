package index_test

import (
	"testing"

	"github.com/franpog859/index"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndex_GetAll(t *testing.T) {
	t.Run("returns proper indexes", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := 1

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, []int{0, 3}, indexes)
	})

	t.Run("returns empty slice if no items match the template", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := 5

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns empty slice if given slice is empty", func(t *testing.T) {
		//given
		slice := []int{}
		item := 4

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns an error if first argument is not a slice", func(t *testing.T) {
		//given
		slice := "not a slice"
		item := 4

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, "index: given non-slice type", err.Error())
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns empty slice if the template is different type", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3}
		item := "not an int"

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns an error if first argument is nil", func(t *testing.T) {
		//given
		var slice interface{}
		item := 1

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, "index: given non-slice type", err.Error())
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns empty slice if no structs do not match template", func(t *testing.T) {
		//given
		type str struct {
			value1 int
			value2 string
		}
		slice := []str{str{5, "first"}, str{1, "second"}}
		item := str{5, "third"}

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns proper indexes for slice of structs", func(t *testing.T) {
		//given
		type str struct {
			value1 int
			value2 string
		}
		slice := []str{str{5, "first"}, str{1, "second"}, str{5, "first"}}
		item := str{5, "first"}

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, []int{0, 2}, indexes)
	})

	t.Run("returns an error if the given slice is multidimensional", func(t *testing.T) {
		//given
		slice := []([]int){{1, 2}, {3}}
		item := []int{3}

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, "index: given multidimensional slice", err.Error())
		assert.Equal(t, []int{}, indexes)
	})

	t.Run("returns an error if the given slice is too complex and contains maps", func(t *testing.T) {
		//given
		slice := make([]map[string]int, 1, 1)
		item := make(map[string]int)
		item["key"] = 1
		slice[0] = item

		//when
		indexes, err := index.GetAll(slice, item)

		//then
		require.Equal(t, "index: given too complex slice of maps", err.Error())
		assert.Equal(t, []int{}, indexes)
	})
}

func TestIndex_IsAny(t *testing.T) {
	t.Run("returns true if item exists in the slice", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := 1

		//when
		isAny, err := index.IsAny(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, true, isAny)
	})

	t.Run("returns false if item does not exist in the slice", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := "not even an int"

		//when
		isAny, err := index.IsAny(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, false, isAny)
	})

	t.Run("returns an error if first argument is not a slice", func(t *testing.T) {
		//given
		slice := "not a slice"
		item := 1

		//when
		isAny, err := index.IsAny(slice, item)

		//then
		require.Equal(t, "index: given non-slice type", err.Error())
		assert.Equal(t, false, isAny)
	})
}

func TestIndex_HowMany(t *testing.T) {
	t.Run("returns correct number of items in the slice", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := 1

		//when
		howMany, err := index.HowMany(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, 2, howMany)
	})

	t.Run("returns zero if there is no proper item in the slice", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := "not even an int"

		//when
		howMany, err := index.HowMany(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, 0, howMany)
	})

	t.Run("returns zero if the slice is empty", func(t *testing.T) {
		//given
		slice := []int{}
		item := 1

		//when
		howMany, err := index.HowMany(slice, item)

		//then
		require.Equal(t, nil, err)
		assert.Equal(t, 0, howMany)
	})

	t.Run("returns an error if first argument is not a slice", func(t *testing.T) {
		//given
		slice := "not a slice"
		item := 1

		//when
		howMany, err := index.HowMany(slice, item)

		//then
		require.Equal(t, "index: given non-slice type", err.Error())
		assert.Equal(t, 0, howMany)
	})
}
