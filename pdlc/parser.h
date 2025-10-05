#pragma once

using namespace boost::spirit;

namespace pam
{
    namespace pdl
    {
#define DECLARE_RULE( name, astNode ) qi::rule<Iterator, astNode, skipper<Iterator> > name;

        template <typename Iterator>
        struct PdlGrammar : qi::grammar< Iterator, ast::NamespaceNode(), skipper< Iterator > >
        {
            explicit PdlGrammar( ErrorHandler< Iterator >& errorHandler ) :
                PdlGrammar::base_type( start )
            {
                qi::_1_type _1;
                //qi::_2_type _2;
                qi::_3_type _3;
                qi::_4_type _4;

                using qi::lit;
                using qi::lexeme;
                using qi::alpha;
                using qi::alnum;
                using qi::string;

                using qi::on_error;
                using qi::on_success;
                using qi::fail;
                using qi::char_;

                using qi::int_;
                using qi::double_;
                using qi::bool_;

                attrRequiredParams.name( "Required Attribute Params" );
                attrOptionalParam.name( "Optional Attribute Param" );
                attrOptionalParams.name( "Optional Attribute Params" );
                attrParams.name( "Required Attribute Params" );
                attribute.name( "Attribute" );
                attributeList.name( "Attributes" );
                start.name( "Namespace Declaration" );
                classList.name( "Classes" );
                classDeclaration.name( "Class Declaration" );
                fullIdentifier.name( "Full Identifier" );
                constDeclaration.name( "Const Declaration" );
                property.name( "Property Declaration" );
                method.name( "Method Declaration" );
                member.name( "Member Declaration" );
                memberList.name( "Members" );
                argument.name( "Argument" );
                argumentList.name( "Arguments" );
                usingLine.name( "Using Directive" );
                usingList.name( "Using Directives" );
                booleanLiteral.name( "Boolean" );
                identifier.name( "Identifier" );
                name.name( "name" );
                literalValue.name( "Literal" );


                using boost::phoenix::function;

                typedef function< ErrorHandler< Iterator > > ErrorHandlerFunction;
                typedef function< Annotation< Iterator > > AnnotationFunction;


                memberAccessModifiers.add
                    ( ACCESS_MOD_PUBLIC_STR, AccessModifiers::amPublic )
                    ( ACCESS_MOD_PROTECTED_STR, AccessModifiers::amProtected )
                    ( ACCESS_MOD_PRIVATE_STR, AccessModifiers::amPrivate )
                    ( ACCESS_MOD_INTERNAL_STR, AccessModifiers::amInternal );

                classAccessModifiers.add
                    ( ACCESS_MOD_PUBLIC_STR, AccessModifiers::amPublic )
                    ( ACCESS_MOD_PROTECTED_STR, AccessModifiers::amProtected )
                    ( ACCESS_MOD_PRIVATE_STR, AccessModifiers::amPrivate )
                    ( ACCESS_MOD_INTERNAL_STR, AccessModifiers::amInternal );

                propertyAccess.add
                    ( "read", PropertyAccess::paRead )
                    ( "write", PropertyAccess::paWrite )
                    ( "readwrite", PropertyAccess::paReadWrite );

                keywords.add
                    ( NAMESPACE_KEYWORD )
                    ( CLASS_KEYWORD )
                    ( METHOD_KEYWORD )
                    ( BOOL_LITERAL_TRUE )
                    ( BOOL_LITERAL_FALSE )
                    ( PROPERTY_KEYWORD );

                nonIdentifierKeywords.add
                    ( CLASS_KEYWORD )
                    ( BOOL_LITERAL_TRUE )
                    ( BOOL_LITERAL_FALSE );

                intrinsicTypes.add
                    ( ITYPE_VOID )
                    ( ITYPE_BOOL )
                    ( ITYPE_INT )
                    ( ITYPE_UINT )
                    ( ITYPE_DOUBLE )
                    ( ITYPE_FUNCTION )
                    ( ITYPE_ARRAY )
                    ( ROOT_OBJECT )
                    ( ITYPE_STRING );


                name =
                    /*   !nonIdentifierKeywords
                   >>  */raw[ lexeme[ ( alpha | '_' ) >> *( alnum | '_' ) ] ];

                identifier = name/* - nonIdentifierKeywords*/;

                fullIdentifier = identifier % '.';

                auto arrayNotation = string( "[]" );

                propertyType = fullIdentifier >> *arrayNotation;

                argument =
                    fullIdentifier
                    > identifier;

                argumentList = argument % ',';

                quotedString = '"' >> quotedStringContent >> '"';
                quotedStringContent = raw[ *( escapeChar | ~char_( '"' ) ) ];
                escapeChar = '\\' >> char_( "\"" );

                /*            quotedString =
                                   omit [ char_("\"") [ _a =_1 ] ]
                                >> no_skip [ *(char_ - char_( _a ) ) ]
                                >> lit( _a );
                                */

                booleanLiteral = string( "true" ) | string( "false" );

                literalValue = quotedString | int_ | double_ | bool_;

                defineAttributes();
                defineProperty();
                defineConst();
                defineMethod();

                member = constDeclaration | property | method | shortProperty;

                memberList = *member;

                //classAccess = classAccessModifiers;

                defineClass();

                usingLine =
                    lit( "using" )
                    > fullIdentifier
                    > ';';

                usingList = *usingLine;

                start =
                    lit( NAMESPACE_KEYWORD )
                    > fullIdentifier
                    > '{'
                    > usingList
                    > classList
                    > '}';


                BOOST_SPIRIT_DEBUG_NODE(attributeList);
                BOOST_SPIRIT_DEBUG_NODE(attribute);
                BOOST_SPIRIT_DEBUG_NODE(fullIdentifier);
                BOOST_SPIRIT_DEBUG_NODE(identifier);

                on_error< fail >( start,
                                  ErrorHandlerFunction( errorHandler )(
                                      "Error! Expecting ", _4, _3 ) );

                on_success( identifier, AnnotationFunction( errorHandler.iters )( _val, _1 ) );
            }

            void defineClass()
            {
                classDeclaration =
                    classAccessModifiers
                    >> lit( CLASS_KEYWORD )
                    > identifier
                    > -( ':' > fullIdentifier )
                    > '{'
                    >> memberList
                    > '}';

                classList = *classDeclaration;
            }

            void defineConst()
            {
                constDeclaration =
                    lit( CONST_KEYWORD )
                    >> fullIdentifier
                    >> identifier
                    >> '='
                    >> literalValue
                    > ';';
            }


            void defineProperty()
            {
                property =
                    attributeList
                    >> memberAccessModifiers
                    >> lit( PROPERTY_KEYWORD )
                    >> -propertyAccess
                    >> propertyType
                    >> identifier
                    >> -( '[' > argumentList > ']' )
                    > ';';

                shortProperty =
                    propertyType
                    >> identifier
                    >> ';';
            }

            void defineMethod()
            {
                method =
                    memberAccessModifiers
                    >> lit( METHOD_KEYWORD )
                    >> fullIdentifier
                    >> identifier
                    >> '(' >> argumentList >> ')'
                    > ';';
            }

            DECLARE_RULE(attrRequiredParams, ast::AttrRequiredParams())
            DECLARE_RULE(attrOptionalParam, ast::AttrOptionalParam())
            DECLARE_RULE(attrOptionalParams, ast::AttrOptionalParams())
            DECLARE_RULE(attrParams, ast::AttrParams())
            DECLARE_RULE(attribute, ast::AttributeNode())
            DECLARE_RULE(attributeList, ast::AttributeListNode())


            void defineAttributes()
            {
                attrOptionalParam =
                    identifier
                    >> '='
                    > literalValue;

                attrRequiredParams = literalValue % ',';
                attrOptionalParams = attrOptionalParam % ',';

                attrParams =
                    attrOptionalParams
                    | ( attrRequiredParams >> ',' > attrOptionalParams )
                    | attrRequiredParams;


                attribute =
                    '['
                    > fullIdentifier
                    > -( '(' > attrParams > ')' )
                    > ']';

                attributeList = *attribute;
            }

            qi::symbols< char > intrinsicTypes;
            qi::symbols< char > keywords;
            qi::symbols< char > nonIdentifierKeywords;
            qi::symbols< char, AccessModifiers > memberAccessModifiers;
            qi::symbols< char, AccessModifiers > classAccessModifiers;
            qi::symbols< char, PropertyAccess > propertyAccess;


            qi::rule< Iterator, std::string(), skipper< Iterator >, qi::locals< char > > quotedString;
            qi::rule< Iterator, std::string(), skipper< Iterator >, qi::locals< char > > quotedStringContent;
            qi::rule< Iterator, std::string(), skipper< Iterator >, qi::locals< char > > escapeChar;

            DECLARE_RULE(start, ast::NamespaceNode())

            DECLARE_RULE(classList, ast::ClassList())
            DECLARE_RULE(classDeclaration, ast::ClassNode())
            DECLARE_RULE(fullIdentifier, ast::FullIdentifierNode())

            DECLARE_RULE( constDeclaration, ast::ConstNode() )

            DECLARE_RULE(propertyType, ast::PropertyType())
            DECLARE_RULE(property, ast::PropertyNode())
            DECLARE_RULE(shortProperty, ast::ShortPropertyNode())
            DECLARE_RULE(method, ast::MethodNode())

            DECLARE_RULE(member, ast::MemberNode())

            DECLARE_RULE(memberList, ast::MemberList())

            DECLARE_RULE(argument, ast::ArgumentNode())
            DECLARE_RULE(argumentList, ast::ArgumentList())

            DECLARE_RULE(usingLine, ast::UsingNode())
            DECLARE_RULE(usingList, ast::UsingList())

            DECLARE_RULE(booleanLiteral, std::string())
            DECLARE_RULE(identifier, ast::Identifier())
            DECLARE_RULE(name, std::string())
            DECLARE_RULE(literalValue, ast::LiteralValueNode())


#undef DECLARE_RULE
        };
    }
}
