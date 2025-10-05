#pragma once

//keywods
#define NAMESPACE_KEYWORD  "namespace"
#define CLASS_KEYWORD  "class"
#define METHOD_KEYWORD  "method"
#define PROPERTY_KEYWORD  "property"
#define CONST_KEYWORD  "const"

//boolean literals
#define BOOL_LITERAL_TRUE  "true"
#define BOOL_LITERAL_FALSE  "false"

//intrinsic types
#define ITYPE_BOOL "bool"
#define ITYPE_VOID "void"
#define ITYPE_INT "int"
#define ITYPE_UINT "uint"
#define ITYPE_DOUBLE "double"
#define ITYPE_STRING "string"
#define ITYPE_ARRAY "array"
#define ITYPE_FUNCTION "function"
#define ROOT_OBJECT "object"

//access modifiers
#define ACCESS_MOD_PUBLIC_STR "public"
#define ACCESS_MOD_PROTECTED_STR "protected"
#define ACCESS_MOD_PRIVATE_STR "private"
#define ACCESS_MOD_INTERNAL_STR "internal"

//class attributes
#define CLASS_ATTR_FINAL "final"
#define CLASS_ATTR_SEALED "sealed"

#define PROPERTY_CONTROL "propertyControl"

namespace pam { namespace pdl {

    enum class AccessModifiers
    {
        amNone = 0,
        amPublic = 1,
        amInternal,
        amProtected,
        amPrivate
    };

    enum class PropertyAccess
    {
        paRead = 1,
        paWrite,
        paReadWrite
    };

}};

/*
std::ostream& operator << (std::ostream& stream, const pam::pdl::AccessModifiers& am);
std::ostream& operator << (std::ostream& stream, const pam::pdl::PropertyAccess& pa);
*/


