package parse

import "go/types"

type Kind types.BasicKind

const (
	Invalid    = Kind(types.Invalid)
	Bool       = Kind(types.Bool)
	Int        = Kind(types.Int)
	Int8       = Kind(types.Int8)
	Int16      = Kind(types.Int16)
	Int32      = Kind(types.Int32)
	Int64      = Kind(types.Int64)
	Uint       = Kind(types.Uint)
	Uint8      = Kind(types.Uint8)
	Uint16     = Kind(types.Uint16)
	Uint32     = Kind(types.Uint32)
	Uint64     = Kind(types.Uint64)
	Float32    = Kind(types.Float32)
	Float64    = Kind(types.Float64)
	Complex64  = Kind(types.Complex64)
	Complex128 = Kind(types.Complex128)
	String     = Kind(types.String)

	Interface = 101
	Bytes     = 102
	Map       = 103
	Ptr       = 104
	Slice     = 105
	Struct    = 106
)

//var SimpleTypes = map[string]Kind{
//	"bool":        Bool,
//	"int":         Int,
//	"int8":        Int8,
//	"int16":       Int16,
//	"int32":       Int32,
//	"int64":       Int64,
//	"uint":        Uint,
//	"uint8":       Uint8,
//	"uint16":      Uint16,
//	"uint32":      Uint32,
//	"uint64":      Uint64,
//	"string":      String,
//}

func (k Kind) IsSimpleType() bool {
	switch k {
	case Bool,
		Int,
		Int8,
		Int16,
		Int32,
		Int64,
		Uint,
		Uint8,
		Uint16,
		Uint32,
		Uint64,
		String:
		return true
	}
	return false
}

// EncodableTypes lists the types that must be encoded for storage (native floats are not supported)
//var EncodableTypes = map[string]Kind{
//	"float32":     Float32,
//	"float64":     Float64,
//	"complex64":   Complex64,
//	"complex128":  Complex128,
//	"interface{}": Interface,
//	"[]byte":      Bytes,
//}
