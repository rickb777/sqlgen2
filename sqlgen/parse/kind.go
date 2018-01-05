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
	Map       = 102
	Slice     = 103
	Struct    = 104
)

func (k Kind) IsShort() bool {
	switch k {
	case Int8,
		Int16,
		Uint8,
		Uint16:
		return true
	}
	return false
}

func (k Kind) IsInteger() bool {
	switch k {
	case Int,
		Int32,
		Int64,
		Uint,
		Uint32,
		Uint64:
		return true
	}
	return false
}

func (k Kind) IsFloat() bool {
	switch k {
	case Float32,
		Float64:
		return true
	}
	return false
}

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

func (k Kind) Token() string {
	switch k {
	case Bool:
		return "bool"
	case Int:
		return "int"
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case Uint:
		return "uint"
	case Uint8:
		return "uint8"
	case Uint16:
		return "uint16"
	case Uint32:
		return "uint32"
	case Uint64:
		return "uint64"
	case String:
		return "string"
	case Interface:
		return "Interface"
	case Map:
		return "Map"
	case Slice:
		return "Slice"
	case Struct:
		return "Struct"
	}
	return "<unknown>"
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
