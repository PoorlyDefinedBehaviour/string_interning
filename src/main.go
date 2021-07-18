package main

import (
	"fmt"
	"unsafe"

	"string_interning/src/vm"
)

// https://craftinginterpreters.com/hash-tables.html
func main() {
	vm := vm.New()

	fmt.Printf("\n\naaaaaaa  %+v\n\n", unsafe.Pointer(vm.String("hello")) == unsafe.Pointer(vm.String("hello")))
	fmt.Printf("\n\naaaaaaa  %+v\n\n", unsafe.Pointer(vm.String("hello")) == unsafe.Pointer(vm.String("hello 1")))
}
