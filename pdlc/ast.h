#pragma once

#include <boost/config/warning_disable.hpp>
#include <boost/variant/recursive_variant.hpp>
#include <boost/fusion/include/adapt_struct.hpp>
#include <boost/fusion/include/io.hpp>
#include <boost/optional.hpp>
#include <list>



namespace pam { namespace pdl { namespace ast
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

        pam::pdl::symbols::ClassSymbolPtr symbol;
    };

    struct UsingList : std::list<UsingNode>, BaseNode
    {
    };



    struct NamespaceNode: BaseNode
    {
        FullIdentifierNode name;
        UsingList usings;
        ClassList classes;

        pam::pdl::symbols::NamespaceSymbolPtr symbol;
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



BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::ShortPropertyNode,
    (pam::pdl::ast::PropertyType, propertyType)
    (pam::pdl::ast::Identifier, name)
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::PropertyType,
    (pam::pdl::ast::FullIdentifierNode, type)
    (pam::pdl::ast::ArrayNotationList, arrayNotationList)
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::PropertyNode,
    (pam::pdl::ast::AttributeListNode, attributes)
    (pam::pdl::AccessModifiers, accessModifier)
    (boost::optional<pam::pdl::PropertyAccess>, access)
    (pam::pdl::ast::PropertyType, propertyType)
    (pam::pdl::ast::Identifier, name)
    (pam::pdl::ast::ArgumentList, arguments)
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::ConstNode,
    ( pam::pdl::ast::FullIdentifierNode, type )
    ( pam::pdl::ast::Identifier, name )
    ( pam::pdl::ast::LiteralValueNode, value )
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::MethodNode,
    (pam::pdl::AccessModifiers, accessModifier)
    (pam::pdl::ast::FullIdentifierNode, type)
    (pam::pdl::ast::Identifier, name)
    (pam::pdl::ast::ArgumentList, arguments)
)


BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::NamespaceNode,
    (pam::pdl::ast::FullIdentifierNode, name)
    (pam::pdl::ast::UsingList, usings)
    (pam::pdl::ast::ClassList, classes)
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::ClassNode,
    (pam::pdl::AccessModifiers, accessModifier)
    (pam::pdl::ast::Identifier, name)
    (boost::optional<pam::pdl::ast::FullIdentifierNode>, parentClass)
    (pam::pdl::ast::MemberList, members )
)


BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::ArgumentNode,
    (pam::pdl::ast::FullIdentifierNode, type)
    (pam::pdl::ast::Identifier, name)
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::UsingNode,
    (pam::pdl::ast::FullIdentifierNode, className)
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::AttributeNode,
    (pam::pdl::ast::FullIdentifierNode, name )
    (pam::pdl::ast::AttrParams, params )
)

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::AttrRequiredAndOptionals,
    (pam::pdl::ast::AttrRequiredParams, required)
    (pam::pdl::ast::AttrOptionalParams, optionals)
)
/*
BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::attrRequiredParam,
    (pam::pdl::ast::LiteralValueNode, value )
)
*/

BOOST_FUSION_ADAPT_STRUCT(
    pam::pdl::ast::AttrOptionalParam,
    (pam::pdl::ast::Identifier, name )
    (pam::pdl::ast::LiteralValueNode, value )
)



std::string operator+(const char* right, pam::pdl::ast::Identifier const& left );
std::string operator+(pam::pdl::ast::Identifier const& left, const char* right );

std::string operator+(const char* left, pam::pdl::ast::FullIdentifierList const& right);
std::string operator+( pam::pdl::ast::FullIdentifierList const& left, const char* right);

