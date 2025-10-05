#pragma once

#define PHPGEN_TEMPLATE_USE_BLOCK "[use-block]"
#define PHPGEN_TEMPLATE_NAMESPACE "[namespace-name]"
#define PHPGEN_TEMPLATE_CLASS_ATTRS "[class-attrs]"
#define PHPGEN_TEMPLATE_CLASS_NAME "[class-name]"
#define PHPGEN_TEMPLATE_CLASS_INHERITANCE "[inheritance]"

#define PHPGEN_TEMPLATE_PHPDOC_INHERITANCE "[phpdoc-inheritance]"
#define PHPGEN_TEMPLATE_PHPDOC_CLASS "[phpdoc-class]"
#define PHPGEN_TEMPLATE_PHPDOC_PACKAGE "[phpdoc-package]"

#define PHPGEN_TEMPLATE_METHODS "[method-list]"
#define PHPGEN_TEMPLATE_CONSTS "[const-list]"
#define PHPGEN_TEMPLATE_PROPERTIES "[property-list]"
#define PHPGEN_TEMPLATE_PROPERTY_ATTRS_BLOCK "[property-attributes]"
#define PHPGEN_TEMPLATE_PROPERTY_ATTRS_VAR "[property-attributes-var]"
#define PHPGEN_TEMPLATE_CONSTRUCTOR "[constructor]"

namespace pam { namespace pdl { namespace codegen
{

using namespace pam::pdl::symbols;

class PhpGenerator : public Generator
{
public:
	virtual ~PhpGenerator()
	{
	}

	void readPhpConfig();
	PhpGenerator( config::PdlConfig const& pdlConfig );
	static std::string fullIdentifierToPsr4(ast::FullIdentifierNode const& fullIdentifier);

protected:

	virtual std::string getTargetLanguageNamespace(ast::NamespaceNode const& astNode) override;


	bool isIgnoreNamespace( std::string const& namespaceString );
	bool isInMainNamespace(std::string const& namespaceString) const;

	virtual std::string doUsingList( ast::UsingList const& astNode ) override;
    virtual std::string doParentClass( ast::ClassNode const& classAstNode ) override;
    virtual std::string doClass( ast::NamespaceNode const& astNamespace, pam::pdl::ast::ClassNode const& astNode ) override;
    virtual std::string doArgument( ast::ArgumentNode const& astNode ) override;

    virtual std::string doConst( ast::ConstNode const& constNode ) override;
    virtual std::string doMethod( ast::MethodNode const& astNode ) override;
    virtual std::string doProperty( ast::PropertyNode const& astNode ) override;
	virtual std::string doShortProperty(ast::ShortPropertyNode const& astNode) override;
	

    virtual std::string visitLiteralString( std::string const& value ) override;
    virtual bool outputClass( std::string const& classSource, ast::ClassNode const& classAstNode ) override;

    virtual AttributeInfo processAttributeName( ast::FullIdentifierNode const& attrName ) override;
    virtual boost::filesystem::path getFileOutputFolder() override;

    std::string generatePropertyAttributeBlock() const;

	std::string doSinglePropertyBody(ast::Identifier const& name, ast::PropertyType const& propertyType);
	std::string doSingleProperty( ast::PropertyNode const& astNode );
	std::string doSingleProperty( ast::ShortPropertyNode const& astNode );
	static std::string fullClassNameToPsr4(std::string const& fullClassName );

	template<typename TPropertyNode>
	std::string doPropertyAttributes(TPropertyNode const& propertyNode)
	{
		std::vector<std::string> attributeVector;

		for (auto const& attribute : propertyNode.attributes)
		{
			std::string params = boost::apply_visitor(_attrParamsVisitor, attribute.params);
			const auto attrInfo = processAttributeName(attribute.name);

			if (params == "")
			{
				const auto fullName = SymbolTable::joinIdentifier(attribute.name);
				if (fullName == "browseableclass")
				{
					auto& classSymbol = *static_cast<ClassSymbol*>(propertyNode.symbol->getTypeSymbol().get());

					const auto fullClassName = fullClassNameToPsr4(classSymbol.getNamespace()->name() + "." + classSymbol.name());

					params = _indent + "'default1' => '" + fullClassName + "'";
				}
			}

			if (params == "")
			{
				attributeVector.push_back(
					(boost::format(_formatter.paramlessAttr)
						% attrInfo.name
						).str());
			}
			else
			{
				attributeVector.push_back(
					(boost::format(_formatter.attr)
						% attrInfo.name
						% params).str());
			}
		}

		auto result = boost::join(attributeVector, _attributeSeparator);

		return result;
	}

	std::string doIndexerProperty( ast::PropertyNode const& astNode );
	std::string fullIdentifierToString(ast::FullIdentifierNode const& fullIdentifier) const;

    virtual std::string visitAttrOptionalParams( ast::AttrOptionalParams const& astNode ) override;
    virtual std::string visitAttrRequiredParams( ast::AttrRequiredParams const& astNode ) override;
    virtual std::string visitAttrRequiredAndOptionals( ast::AttrRequiredAndOptionals const& astNode ) override;

//    std::string internalGenerateMethod(
//        std::string const& accessModifier,
//        std::string const& type,
//        std::string const& name, 
//        std::string const& virtuality,
//        pam::pdl::ast::ArgumentList const& astNode
//    );

    

private:

    std::set<std::string> _useLines;
    std::map<std::string,std::string> _propertyAttributes;
	bool _phpPsr4 = false;
	std::set<std::string> _ignoreNamespaces;
};

}}};