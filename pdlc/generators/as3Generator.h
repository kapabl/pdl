#pragma once

#define AS3GEN_TEMPLATE_IMPORT_BLOCK "[import-block]"
#define AS3GEN_TEMPLATE_PACKAGE "[package-name]"
#define AS3GEN_TEMPLATE_CLASS_ATTRS "[class-attrs]"
#define AS3GEN_TEMPLATE_CLASS_NAME "[class-name]"
#define AS3GEN_TEMPLATE_CLASS_INHERITANCE "[inheritance]"
#define AS3GEN_TEMPLATE_METHODS "[method-list]"
#define AS3GEN_TEMPLATE_PARENT_CONSTRUCTOR "[parent-constructor]"
#define AS3GEN_TEMPLATE_PROPERTIES "[property-list]"
#define AS3GEN_TEMPLATE_CONSTS "[const-list]"

namespace pam { namespace pdl { namespace codegen
{

class As3Generator : public Generator
{
public:
    As3Generator( pam::pdl::config::PdlConfig const& config );
	bool isInMainNamespace( std::string const& namespaceString ) const;


protected:
    virtual std::string doUsingList( pam::pdl::ast::UsingList const& usingList ) override;
    virtual std::string doParentClass( pam::pdl::ast::ClassNode const& classAstNode ) override;
    virtual std::string doClass( pam::pdl::ast::NamespaceNode const& astNamespace, pam::pdl::ast::ClassNode const&
                                 classAstNode ) override;
    virtual std::string doArgument( pam::pdl::ast::ArgumentNode const& astNode ) override;

    virtual std::string doConst( ast::ConstNode const& constNode ) override;
    virtual std::string doMethod( pam::pdl::ast::MethodNode const& astNode ) override;
    virtual std::string doProperty( pam::pdl::ast::PropertyNode const& astNode ) override;
    virtual std::string doShortProperty( pam::pdl::ast::ShortPropertyNode const& shortPropertyNode ) override;

    virtual std::string visitLiteralString( std::string const& value ) override;

    template<typename TPropertyNode>
    std::string doSinglePropertyWithField( TPropertyNode const& propertyNode,
        std::string const& visibility, PropertyAccess propertyAccess, std::string const& attributes )
    {
        const auto type = getPropertyType( propertyNode.propertyType );

        const auto name = toCamelString( propertyNode.name.name );
        const auto fieldName = fieldNameFromPropertyName( name );

        std::string getMethod;
        std::string setMethod;

        /*
        _formatter.singleGetProperty = "\n%4% function get %2%():%1% { return this.%3%; }";
        _formatter.singleSetProperty = "\n%4% function set %2%( value: %1% ):void { this.%3% = value };";
        */

        if ( propertyAccess == PropertyAccess::paRead || propertyAccess == PropertyAccess::paReadWrite )
        {
            getMethod = ( boost::format( _formatter.singleGetProperty )
                % type
                % name
                % fieldName
                % visibility ).str();
        }

        if ( propertyAccess == PropertyAccess::paWrite || propertyAccess == PropertyAccess::paReadWrite )
        {
            setMethod = ( boost::format( _formatter.singleSetProperty )
                % type
                % name
                % fieldName
                % visibility ).str();
        }

        auto result =
            singleFieldType( type, fieldName ) + "\n\n" +
            attributes +
            getMethod +
            setMethod;

        return result;
    }


    virtual bool outputClass( std::string const& classSource, pam::pdl::ast::ClassNode const& classAstNode ) override;

    virtual AttributeInfo processAttributeName( pam::pdl::ast::FullIdentifierNode const& attrName ) override;

    virtual boost::filesystem::path getFileOutputFolder() override;

    template<typename TPropertyNode>
    std::string doSinglePropertyWithCall( TPropertyNode const& propertyNode, std::string const& visibility, 
                                          PropertyAccess propertyAccess, std::string const& attributes )
    {
        const auto type = getPropertyType( propertyNode.propertyType );
        const auto name = toCamelString( propertyNode.name.name );

        std::string getMethod;
        std::string setMethod;

        if ( propertyAccess == PropertyAccess::paRead || propertyAccess == PropertyAccess::paReadWrite )
        {
            getMethod = ( boost::format( _controllableGetTemplate )
                % type
                % name
                % visibility ).str();
        }

        if ( propertyAccess == PropertyAccess::paWrite || propertyAccess == PropertyAccess::paReadWrite )
        {
            setMethod = ( boost::format( _controllableSetTemplate )
                % type
                % name
                % visibility ).str();
        }


        auto result =
            attributes +
            getMethod +
            setMethod;

        return result;
    }
    

    std::string doSingleProperty( pam::pdl::ast::PropertyNode const& astNode );
    std::string doSingleProperty( pam::pdl::ast::ShortPropertyNode const& astNode );

	virtual std::string getPropertyType( ast::PropertyType const& propertyType ) override;
	std::string doIndexerProperty( pam::pdl::ast::PropertyNode const& astNode );

    std::string internalGenerateMethod(
        std::string const& accessModifier,
        std::string const& type,
        std::string const& name, 
        std::string const& virtuality,
        pam::pdl::ast::ArgumentList const& astNode
    );

    

private:

    std::set<std::string> _importLines;
    std::string _controllableGetTemplate;
    std::string _controllableSetTemplate;

};

}}};