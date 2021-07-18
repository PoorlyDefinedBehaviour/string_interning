package vm

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func Test_vm_String(t *testing.T) {
	t.Parallel()

	t.Run("the same pointer should always be returned for strings that have the same contents", func(t *testing.T) {
		t.Parallel()

		vm := New()

		assert.Equal(t, unsafe.Pointer(vm.String("hello")), unsafe.Pointer(vm.String("hello")))
		assert.NotEqual(t, unsafe.Pointer(vm.String("hello")), unsafe.Pointer(vm.String("hello 1")))
		assert.Equal(t, unsafe.Pointer(vm.String("hello 1")), unsafe.Pointer(vm.String("hello 1")))
	})
}
