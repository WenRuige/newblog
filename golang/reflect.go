package main

import (
	"unsafe"
	"fmt"
	"reflect"
)

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  *rtype
	word unsafe.Pointer
}

type rtype struct {
	size    uintptr
	ptrdata uintptr // number of bytes in the type that can contain pointers
	hash    uint32  // hash of type; avoids computation in hash tables
	//tflag      tflag    // extra type information flags
	align      uint8 // alignment of variable with this type
	fieldAlign uint8 // alignment of struct field with this type
	kind       uint8 // enumeration for C
	//alg        *typeAlg // algorithm table
	gcdata *byte // garbage collection data
	//	str        nameOff  // string form
	//	ptrToThis  typeOff  // type for pointer to this type, may be zero
}

func main() {
	var s int = 42
	res := *(*emptyInterface)(unsafe.Pointer(&s))
	///res.typ
	fmt.Println(res.word)
	fmt.Println(reflect.TypeOf(s))
}
