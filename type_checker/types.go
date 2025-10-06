package type_checker

import (
	"fmt"
	"log"

	"github.com/ElrohirGT/5318008Lang/lib"
)

// =====================
// WHAT IS A TYPE
// =====================

// ID given for each defined type.
type TypeIdentifier string

func NewArrayTypeIdentifier(baseType TypeIdentifier) TypeIdentifier {
	return TypeIdentifier(fmt.Sprintf("%s[]", baseType))
}

// Metadata about a type. It works as an enum, all types share common attributes
// Extra attributes may be stored on its own field as an Optional.
type TypeInfo struct {
	BaseType  bool // If this a simple type (does not carry class/array info)
	Size      uint // In bytes
	ArrayType lib.Optional[ArrayTypeInfo]
	ClassType lib.Optional[ClassTypeInfo]
}

// Data Structure to store the types defined within a compiscript code.
type TypeTable = map[TypeIdentifier]TypeInfo

// Data Structure to store the types of the literal expressions found in the program.
type LiteralTable = map[string]TypeIdentifier

func NewTypeInfo_Base(size uint) TypeInfo {
	return TypeInfo{
		BaseType:  true,
		Size:      size,
		ArrayType: lib.NewOpEmpty[ArrayTypeInfo](),
		ClassType: lib.NewOpEmpty[ClassTypeInfo](),
	}
}

// =====================
// BASE/PRIMITIVE TYPES
// =====================

var BASE_TYPES = struct {
	INTEGER TypeIdentifier
	BOOLEAN TypeIdentifier
	STRING  TypeIdentifier
	NULL    TypeIdentifier
	// Used when the type of something is not known, can be used as a default type before computing stuff.
	UNKNOWN TypeIdentifier
	// Special type used when we know for sure the type of something is invalid but still need to assign a type!
	INVALID TypeIdentifier
}{
	INTEGER: "integer",
	BOOLEAN: "boolean",
	STRING:  "string",
	NULL:    "null",
	UNKNOWN: "**unknown**",
	INVALID: "**invalid**",
}

var BASE_TYPE_ARRAY = []TypeIdentifier{
	BASE_TYPES.INTEGER,
	BASE_TYPES.BOOLEAN,
	BASE_TYPES.STRING,
	BASE_TYPES.NULL,
	BASE_TYPES.INVALID,
	BASE_TYPES.UNKNOWN,
}

// =====================
// ARRAYS
// =====================

type ArrayTypeInfo struct {
	Type TypeIdentifier
}

func NewTypeInfo_Array(arrInfo ArrayTypeInfo) TypeInfo {
	return TypeInfo{
		BaseType:  false,
		ArrayType: lib.NewOpValue(arrInfo),
		ClassType: lib.NewOpEmpty[ClassTypeInfo](),
	}
}

// =====================
// METHOD/FUNCTIONS
// =====================

type ParameterInfo struct {
	Name string
	Type TypeIdentifier
}

type MethodInfo struct {
	ParameterList []ParameterInfo
	ReturnType    TypeIdentifier
}

// =====================
// CLASSES
// =====================

const CONSTRUCTOR_NAME = "constructor"

type ClassTypeInfo struct {
	Name           TypeIdentifier
	InheritsFrom   TypeIdentifier
	FieldOrder     []string
	Fields         map[string]TypeIdentifier
	ConstantFields lib.Set[string]
	Methods        map[string]MethodInfo
	constructor    *MethodInfo
}

func (classInfo *ClassTypeInfo) GetFieldOffset(l *Listener, fieldName string) uint {
	var computedOffset uint = 0
	// FIXME: FUCK INHERITANCE
	// ALL MY HOMIES HATE INHERITANCE
	for _, fName := range classInfo.FieldOrder {
		fieldType, found := classInfo.GetFieldType(fieldName, l)
		if !found {
			log.Panicf(
				"Failed to find type for field `%s` of class `%s`",
				fieldName,
				classInfo.Name,
			)
		}

		fieldTypeInfo, found := l.GetTypeInfo(fieldType)
		if !found {
			log.Panicf(
				"Failed to find the type information for field `%s` of class `%s`",
				fieldName,
				classInfo.Name,
			)
		}

		if fName == fieldName {
			break
		}

		computedOffset += fieldTypeInfo.Size
	}

	return computedOffset
}

func (c *ClassTypeInfo) GetConstructor(l *Listener) MethodInfo {
	if c.constructor != nil {
		return *c.constructor
	}

	if fatherInfo, found := l.GetTypeInfo(c.InheritsFrom); c.InheritsFrom != "" && found {
		fatherInf := fatherInfo.ClassType.GetValue()
		return fatherInf.GetConstructor(l)
	}

	return MethodInfo{
		ReturnType: c.Name,
	}
}

func (c *ClassTypeInfo) GetFieldType(fieldName string, listener *Listener) (TypeIdentifier, bool) {
	fieldType, found := c.Fields[fieldName]
	if !found {
		if fatherInfo, found := listener.GetTypeInfo(c.InheritsFrom); c.InheritsFrom != "" && found {
			fatherInf := fatherInfo.ClassType.GetValue()
			return fatherInf.GetFieldType(fieldName, listener)
		}
	}

	return fieldType, found
}

func (c *ClassTypeInfo) UpsertField(name string, _type TypeIdentifier) {
	c.FieldOrder = append(c.FieldOrder, name)
	c.Fields[name] = _type
}

func (c *ClassTypeInfo) UpsertMethod(name string, info MethodInfo) {
	c.Methods[name] = info
}

func NewClassTypeInfo(className string) ClassTypeInfo {
	return ClassTypeInfo{
		Name:       TypeIdentifier(className),
		Fields:     make(map[string]TypeIdentifier),
		Methods:    make(map[string]MethodInfo),
		FieldOrder: make([]string, 0),
	}
}

func NewTypeInfo_Class(classInfo ClassTypeInfo, size uint) TypeInfo {
	return TypeInfo{
		BaseType:  false,
		Size:      size,
		ArrayType: lib.NewOpEmpty[ArrayTypeInfo](),
		ClassType: lib.NewOpValue(classInfo),
	}
}
