#pragma once

#define GENERIC_TEMPLATE_HEADER "[header]"

namespace pam
{
    namespace pdl
    {
        namespace codegen
        {
            class Generator
            {
            public:
                typedef std::map< AccessModifiers, std::string > AccessModifierMap;

                Generator( config::PdlConfig const& pdlConfig, std::string const& language );

                bool isSameSource() const { return _isSameSource; }
                void setIsSameSource( bool value ) { _isSameSource = value; }

                virtual std::string doNamespace( ast::NamespaceNode const& astNamespace );

                static bool isEnabled( config::PdlConfig& pdlConfig, std::string const& language );

            protected:

                struct AttributeInfo;

                virtual ~Generator()
                {
                }

                void addScalarType( std::string const& type ) { _scalarTypes.insert( type ); }

                bool isScalarType( std::string const& type )
                {
                    auto result = _scalarTypes.find( type ) != _scalarTypes.end();
                    return result;
                }


                void readConfig();
                AccessModifierMap _accessModifierMap;
                bool _isSameSource;
                virtual std::string getTargetLanguageNamespace( ast::NamespaceNode const& astNamespace );

                virtual std::string doUsingList( ast::UsingList const& astUsingList ) = 0;
                virtual std::string doParentClass( ast::ClassNode const& astClass ) = 0;
                virtual std::string doClass( ast::NamespaceNode const& astNamespace,
                                             ast::ClassNode const& astClass ) = 0;
                virtual std::string doProperty( ast::PropertyNode const& astProperty ) = 0;
                virtual std::string doShortProperty( ast::ShortPropertyNode const& astProperty ) = 0;
                virtual std::string doMethod( ast::MethodNode const& astMethod ) = 0;
                virtual std::string doConst( ast::ConstNode const& constNode ) = 0;

                virtual std::string doArgument( ast::ArgumentNode const& astArgument ) = 0;
                virtual AttributeInfo processAttributeName( ast::FullIdentifierNode const& astFullIndenfier ) = 0;

                virtual std::string doArgumentList( ast::ArgumentList const& astArgumentList );

                virtual std::string doAttributes( ast::AttributeListNode const& astAttributeList );

                virtual void doMembers( ast::MemberList const& astMemberList );

                virtual std::string visitProperty( ast::PropertyNode const& astProperty );
                virtual std::string visitShortProperty( ast::ShortPropertyNode const& shortPropertyNode );

                virtual std::string visitMethod( ast::MethodNode const& astMethod );
                virtual std::string visitConst( ast::ConstNode const& constNode );
                virtual std::string visitLiteralString( std::string const& value ) = 0;
                virtual std::string visitAttrOptionalParams( ast::AttrOptionalParams const& astAttrOptionalParams );
                virtual std::string visitAttrRequiredParams( ast::AttrRequiredParams const& astAttrRequiredParams );
                virtual std::string visitAttrRequiredAndOptionals(
                    ast::AttrRequiredAndOptionals const& astAttrRequiredAndOptionnals );
                virtual std::string visitLiteralInt( int const& value ) { return std::to_string( value ); }
                virtual std::string visitLiteralDouble( double const& value ) { return std::to_string( value ); }
                virtual std::string visitLiteralBool( bool const& value ) { return value ? "true" : "false"; }
                virtual std::string readTemplateCode();

                virtual std::string createFileHeader( std::string const& fullClassName );

                virtual bool outputClass( std::string const& classSource, ast::ClassNode const& astClass ) = 0;
                virtual std::string getPropertyType( ast::PropertyType const& propertyType );
                std::string arrayNotationList2Brackets( ast::ArrayNotationList const& arrayNotationList );
                virtual boost::filesystem::path getFileOutputFolder() = 0;
                void writeOutputFile( std::string const& source ) const;
                virtual bool outputClass( std::string const& classSource, std::string const& fullClassName );

                virtual std::string fieldNameFromPropertyName( std::string name );

                virtual std::string singleFieldType( std::string const& fieldType, std::string const& fieldName );
                virtual std::string indexerFieldType( std::string const& fieldType, std::string const& fieldName,
                                                      std::string const& indexerType );

                virtual std::string translateType( std::string const& pdlType );
                virtual std::string getEmptyValueOfType( std::string const& type );
                virtual void prepareOutputFolder( ast::NamespaceNode const& astNamespace,
                                                  ast::ClassNode const& astClass );


                std::string getAccessModifier( AccessModifiers const accessModifier );

                static bool isFullyQualified( std::string const& name )
                {
                    return name.find( '.' ) != std::string::npos;
                }

                static bool isFullyQualified( ast::FullIdentifierNode const& type ) { return type.size() > 1; }


                struct MemberVisitor
                {
                    typedef std::string ResultType;

                    Generator& gen;

                    explicit MemberVisitor( Generator& gen ) : gen( gen )
                    {
                    };

                    ResultType operator()( ast::PropertyNode const& astProperty ) const
                    {
                        return gen.visitProperty( astProperty );
                    }

                    ResultType operator()( ast::ShortPropertyNode const& astProperty ) const
                    {
                        return gen.visitShortProperty( astProperty );
                    }

                    ResultType operator()( ast::MethodNode const& astMethod ) const
                    {
                        return gen.visitMethod( astMethod );
                    }

                    ResultType operator()( ast::ConstNode const& constNode ) const
                    {
                        return gen.visitConst( constNode );
                    }
                };

                struct LiteralVisitor
                {
                    typedef std::string ResultType;

                    Generator& gen;

                    explicit LiteralVisitor( Generator& gen ) : gen( gen )
                    {
                    };

                    ResultType operator()( int const& value ) const { return gen.visitLiteralInt( value ); }
                    ResultType operator()( double const& value ) const { return gen.visitLiteralDouble( value ); }
                    ResultType operator()( std::string const& value ) const { return gen.visitLiteralString( value ); }
                    ResultType operator()( bool const& value ) const { return gen.visitLiteralBool( value ); }
                };

                struct AttrParamsVisitor
                {
                    typedef std::string ResultType;

                    Generator& gen;

                    explicit AttrParamsVisitor( Generator& generator ) : gen( generator )
                    {
                    };

                    ResultType operator()( ast::AttrOptionalParams const& astAttrOptionalParams ) const
                    {
                        return gen.visitAttrOptionalParams( astAttrOptionalParams );
                    }

                    ResultType operator()( ast::AttrRequiredParams const& astAttrRequiredParams ) const
                    {
                        return gen.visitAttrRequiredParams( astAttrRequiredParams );
                    }

                    ResultType operator()( ast::AttrRequiredAndOptionals const& astAttrRequiredAndOptionals ) const
                    {
                        return gen.visitAttrRequiredAndOptionals( astAttrRequiredAndOptionals );
                    }
                };

                struct AttributeInfo
                {
                    symbols::SymbolTable::ParsedClassname classInfo;
                    std::string name;
                };

                struct CodeFormatters
                {
                    std::string paramlessAttr;
                    std::string attr;

                    std::string singleField;
                    std::string indexerField;

                    std::string singleGetProperty;
                    std::string indexerGetProperty;

                    std::string singleSetProperty;
                    std::string indexerSetProperty;

                    std::string constMember;
                };


                static std::string toCamelString( std::string const& value )
                {
                    auto result = value;
                    result[ 0 ] = tolower( value[ 0 ] );
                    return result;
                }

                static std::string toPascalString( std::string const& value )
                {
                    auto result = value;
                    result[ 0 ] = toupper( value[ 0 ] );
                    return result;
                }


                std::string _fileExt;
                //std::map<std::string, std::string> _config;

                config::PdlConfig _config;
                web::json::value _languageConfig;
                std::string _language;

                bool _hasPropertyControl;

                std::string _indent;
                std::string _outputFolder;
                std::string _inputFolder;
                std::string _classTemplate;

                LiteralVisitor _literalVisitor;
                MemberVisitor _memberVisitor;
                AttrParamsVisitor _attrParamsVisitor;

                std::vector< std::string > _methods;
                std::vector< std::string > _consts;
                std::vector< std::string > _properties;

                /*std::string _usingBlock;*/
                std::string _mainNamespace;

                std::string _pdlMainNamespace;

                CodeFormatters _formatter;

                //used to convernt from PDL data types to
                //language specific data types
                std::unordered_map< std::string, std::string > _typesMap;
                std::unordered_map< std::string, std::string > _emptyValueOfType;
                std::string _nullObjectValue;

                std::string _attributeSeparator;
                std::string _outputFileName;
                std::string _outputHeader;

                std::unordered_set< std::string > _scalarTypes;
            };
        }
    }
};
