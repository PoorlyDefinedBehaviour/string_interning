package vm

import (
	"unsafe"

	"string_interning/src/hashtable"
)

// Imagine we are implementing an interpreter
// in a language can't compare two strings
// using == and we are trying to implement the operator == for strings.
// We could traverse the strings and compare each characters
// but our comparison would always take O(n) time.
// String interning allows us to make it faster some of the time.
//
// The core problem is that it's possible to have
// different strings in memory with the same characters.
// Those need to behave like equivalent vaues even though
// they are distinct objects. They're essentially duplicates,
// and we have to compare all of their bytes to detect that.
//
// String interning is a process of deduplication.
// We create a collection of interned strings.
// Any string in that collection is guaranteed to be textually
// distinct fromm all others. When you internet a string,
// you look for a matching string in the collection.
// If found, you use that original one. Otherwise, the string you have
// is unique, so you add it to the collection.
//
// In our language, the runtime keeps a collection of strings,
// when a string is created, the runtime interns it and returns a pointer.
//
// When the runtime needs to compare two strings, it can just do pointer comparison.
// Given the strings a and b where both are interned, if both are the same pointer, they are equal.
//
// How other languages do it:
// Lua interns all strings.
// Lisp, Scheme, Smalltalk, Ruby and others have a separate string-like
// type called symbol that is implicitly interned. This is why they say
// symbols are faster in Ruby.
// Java interns constant strings by default, and provides an API to let you
// explicitly intern any stirng you give it.
type vm struct {
	strings hashtable.T
}

type T = vm

type String struct {
	len   int
	chars unsafe.Pointer
}

func newString(s string) *String {
	return &String{
		len:   len(s),
		chars: unsafe.Pointer(&s),
	}
}

func (vm *vm) intern(s string) *String {
	internedString := vm.strings.Get(s)
	if internedString != nil {
		return internedString.(*String)
	}

	vm.strings.Set(s, newString(s))

	return vm.strings.Get(s).(*String)
}

func (vm *vm) String(s string) *String {
	return vm.intern(s)
}

func New() T {
	return T{
		strings: hashtable.New(),
	}
}
