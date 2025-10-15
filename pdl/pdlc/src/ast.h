#pragma once

#include <boost/config/warning_disable.hpp>
#include <boost/variant/recursive_variant.hpp>
#include <boost/fusion/include/adapt_struct.hpp>
#include <boost/fusion/include/io.hpp>
#include <boost/optional.hpp>
#include <list>



namespace io { namespace pdl { namespace ast
{
    struct Tagged
    {
        int id; // Used to annotate the AST with the iterator position.
                // This id is used as a key to a map<int, Iterator>
                // (not really part of the AST.)
    };

    struct Nil {};

    struct NamespaceNode;
    struct ClassNode;
    struct PropertyNode;
    struct MethodNode;

    struct BaseNode
    {
        bool isPending;
        BaseNode() : isPending( false ) {}
    };

    struct Identifier : Tagged, BaseNode
    {
        Identifier(std::string const& name = "")  : 
            BaseNode(),
            name( name )
        {}

        std::string name;
    };

    typedef std::list<Identifier> FullIdentifierList;

    struct ClassList : std::list<ClassNode>, BaseNode
    {
    };

    struct FullIdentifierNode: FullIdentifierList, BaseNode
    {
    };

    struct UsingNode: BaseNode
    {
        FullIdentifierNode className;

        io::pdl::symbols::ClassSymbolPtr symbol;
    };

    struct UsingList : std::list<UsingNode>, BaseNode
    {
    };



    struct NamespaceNode: BaseNode
    {
        FullIdentifierNode name;
        UsingList usings;
        ClassList classes;

        io::pdl::symbols::NamespaceSymbolPtr symbol;
    };

    typedef boost::variant<
        std::string,
        double,
        int,
        bool
    >
    LiteralValueNode;



    struct AttrOptionalParam: BaseNode
    {
        Identifier name;
        LiteralValueNode value;
    };

    typedef std::list<AttrOptionalParam> AttrOptionalParams;
    typedef std::list<LiteralValueNode> AttrRequiredParams;

    struct AttrRequiredAndOptionals: BaseNode
    {
        AttrRequiredParams required;
        AttrOptionalParams optionals;
    };

    typedef boost::variant<
        AttrOptionalParams,
        AttrRequiredAndOptionals,
        AttrRequiredParams
    >
    AttrParams;

    struct AttributeNode: BaseNode
    {
        FullIdentifierNode name;
        AttrParams params;
    };

    struct AttributeListNode: std::list<AttributeNode>, BaseNode
    {
    };


    struct ArrayNotationList : std::list<std::string>, BaseNode
    {
    };


    struct PropertyType : BaseNode
    {
        FullIdentifierNode type;
        ArrayNotationList arrayNotationList;
    };

    struct ShortPropertyNode : BaseNode
    {
        PropertyType propertyType;
        Identifier name;

        symbols::PropertySymbolPtr symbol;
    };

    struct ConstNode : BaseNode
    {
        //FullIdentifierNode type;
        FullIdentifierNode type;
        Identifier name;
        LiteralValueNode value;

        symbols::ConstSymbolPtr symbol;
    };




    typedef boost::variant<
          ConstNode, 
          MethodNode, 
          PropertyNode,
          ShortPropertyNode
        >
    MemberNode;
    

    struct MemberList: std::list<MemberNode>, BaseNode 
    {
    };


    struct ClassNode: BaseNode
    {
        AccessModifiers accessModifier;
        Identifier name;
        boost::optional<FullIdentifierNode> parentClass;

        MemberList members;

        symbols::ClassSymbolPtr symbol;
    };


    struct ArgumentNode: BaseNode
    {
        FullIdentifierNode type;
        Identifier name;
    };

    struct ArgumentList: std::list<ArgumentNode>, BaseNode
    {
    };

    struct MethodNode: BaseNode
    {
        AccessModifiers accessModifier;
        FullIdentifierNode type;
        Identifier name;
        ArgumentList arguments;

        symbols::MethodSymbolPtr symbol;
        
    };



    struct PropertyNode: BaseNode
    {
        AttributeListNode attributes;
        AccessModifiers accessModifier;
        boost::optional<PropertyAccess> access;
        //FullIdentifierNode type;
        PropertyType propertyType;
        Identifier name;
        ArgumentList arguments;


        symbols::PropertySymbolPtr symbol;
    };


    // print functions for debugging
    inline std::ostream& operator<<(std::ostream& out, Nil)
    {
        out << "nil"; return out;
    }

    inline std::ostream& operator<<(std::ostream& out, Identifier const& id)
    {
        out << id.name; return out;
    }

}}}

namespace io_pdl = ::io::pdl;


BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::ShortPropertyNode,
    (io_pdl::ast::PropertyType, propertyType)
    (io_pdl::ast::Identifier, name)
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::PropertyType,
    (io_pdl::ast::FullIdentifierNode, type)
    (io_pdl::ast::ArrayNotationList, arrayNotationList)
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::PropertyNode,
    (io_pdl::ast::AttributeListNode, attributes)
    (io_pdl::AccessModifiers, accessModifier)
    (boost::optional<io_pdl::PropertyAccess>, access)
    (io_pdl::ast::PropertyType, propertyType)
    (io_pdl::ast::Identifier, name)
    (io_pdl::ast::ArgumentList, arguments)
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::ConstNode,
    ( io_pdl::ast::FullIdentifierNode, type )
    ( io_pdl::ast::Identifier, name )
    ( io_pdl::ast::LiteralValueNode, value )
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::MethodNode,
    (io_pdl::AccessModifiers, accessModifier)
    (io_pdl::ast::FullIdentifierNode, type)
    (io_pdl::ast::Identifier, name)
    (io_pdl::ast::ArgumentList, arguments)
)


BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::NamespaceNode,
    (io_pdl::ast::FullIdentifierNode, name)
    (io_pdl::ast::UsingList, usings)
    (io_pdl::ast::ClassList, classes)
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::ClassNode,
    (io_pdl::AccessModifiers, accessModifier)
    (io_pdl::ast::Identifier, name)
    (boost::optional<io_pdl::ast::FullIdentifierNode>, parentClass)
    (io_pdl::ast::MemberList, members )
)


BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::ArgumentNode,
    (io_pdl::ast::FullIdentifierNode, type)
    (io_pdl::ast::Identifier, name)
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::UsingNode,
    (io_pdl::ast::FullIdentifierNode, className)
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::AttributeNode,
    (io_pdl::ast::FullIdentifierNode, name )
    (io_pdl::ast::AttrParams, params )
)

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::AttrRequiredAndOptionals,
    (io_pdl::ast::AttrRequiredParams, required)
    (io_pdl::ast::AttrOptionalParams, optionals)
)
/*
BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::attrRequiredParam,
    (io_pdl::ast::LiteralValueNode, value )
)
*/

BOOST_FUSION_ADAPT_STRUCT(
    io_pdl::ast::AttrOptionalParam,
    (io_pdl::ast::Identifier, name )
    (io_pdl::ast::LiteralValueNode, value )
)



std::string operator+(const char* right, io_pdl::ast::Identifier const& left );
std::string operator+(io_pdl::ast::Identifier const& left, const char* right );

std::string operator+(const char* left, io_pdl::ast::FullIdentifierList const& right);
std::string operator+( io_pdl::ast::FullIdentifierList const& left, const char* right);

