package generator

const (
	TypeBool     = "bool"
	TypeVoid     = "void"
	TypeInt      = "int"
	TypeUInt     = "uint"
	TypeDouble   = "double"
	TypeString   = "string"
	TypeArray    = "array"
	TypeFunction = "function"
	TypeObject   = "object"
)

var intrinsicTypes = map[string]struct{}{
	TypeBool:     {},
	TypeVoid:     {},
	TypeInt:      {},
	TypeUInt:     {},
	TypeDouble:   {},
	TypeString:   {},
	TypeArray:    {},
	TypeFunction: {},
	TypeObject:   {},
}


func IsIntrinsicType(raw string) bool {
	_, exists := intrinsicTypes[raw]
	return exists
}


