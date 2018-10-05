package index_test

import (
	"github.com/franpog859/index"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
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
}
