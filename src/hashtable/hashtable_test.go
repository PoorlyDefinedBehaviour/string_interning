package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashtable(t *testing.T) {
	t.Parallel()

	t.Run("Size", func(t *testing.T) {
		t.Parallel()

		table := New()

		assert.Equal(t, 0, table.Size())

		table.Set("abc", struct{ x int }{x: 1})

		assert.Equal(t, 1, table.Size())

		table.Set("foo", "hello")

		assert.Equal(t, 2, table.Size())

		table.Remove("foo")

		assert.Equal(t, 1, table.Size())

		table.Remove("foo")

		assert.Equal(t, 1, table.Size())

		table.Remove("abc")

		assert.Equal(t, 0, table.Size())

		table.Remove("keythatdoesnotexist")

		assert.Equal(t, 0, table.Size())
	})

	t.Run("IsEmpty", func(t *testing.T) {
		t.Parallel()

		table := New()

		assert.True(t, table.IsEmpty())

		table.Set("a", 1)

		assert.False(t, table.IsEmpty())

		table.Remove("b")

		assert.False(t, table.IsEmpty())

		table.Remove("a")

		assert.True(t, table.IsEmpty())
	})

	t.Run("Has", func(t *testing.T) {
		t.Parallel()

		table := New()

		assert.False(t, table.Has("hello"))

		table.Set("hello", "world")

		assert.True(t, table.Has("hello"))

		assert.False(t, table.Has("hell"))

		table.Remove("hello")

		assert.False(t, table.Has("hello"))
	})

	t.Run("Get", func(t *testing.T) {
		t.Parallel()

		table := New()

		assert.Nil(t, table.Get("a"))

		table.Set("a", 1)

		assert.Equal(t, 1, table.Get("a"))

		assert.Nil(t, table.Get("b"))

		table.Remove("a")

		assert.Nil(t, table.Get("a"))
	})
}
