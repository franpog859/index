package index_test

import (
	"github.com/franpog859/index"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIndex_GetAll(t *testing.T) {
	t.Run("asda", func(t *testing.T) {
		//given
		slice := []int{1, 2, 3, 1}
		item := 1

		//when
		indexes, ok := index.getAll(slice, item)

		//then
		require.Equal(t, true, ok)
		assert.Equal(t, []int{0, 3}, indexes)
	})
}
