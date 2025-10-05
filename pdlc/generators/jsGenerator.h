#pragma once

#define JSGEN_TEMPLATE_USING_BLOCK "[using-block]"
#define JSGEN_TEMPLATE_NAMESPACE "[namespace]"
#define JSGEN_TEMPLATE_CLASS_ATTRS "[class-attrs]"
#define JSGEN_TEMPLATE_CLASS_NAME "[class-name]"
#define JSGEN_TEMPLATE_FULL_PARENT_CLASS "[full-parent-class]"
#define JSGEN_TEMPLATE_METHODS "[method-list]"
#define JSGEN_TEMPLATE_PROPERTIES "[property-list]"
#define JSGEN_TEMPLATE_CONSTS "[const-list]"
#define JSGEN_TEMPLATE_PROPERTIES_ATTRIBUTES "[property-attributes]"
#define JSGEN_TEMPLATE_CLASS_DECLARATION "[class-declaration]"
#define JSGEN_TEMPLATE_JSDOC_CLASS "[jsdoc-class]"
#define JSGEN_TEMPLATE_JSDOC_PROPERTIES "[jsdoc-properties]"

#define JSGEN_ATTRI_JSTYPE "jstype"

namespace pam { namespace pdl { namespace codegen
{

class JsGenerator : public Generator
{
public:
    JsGenerator( config::PdlConfig const& config );

    virtual std::string doNamespace( ast::NamespaceNode const& astNode ) override;
protected:

    virtual std::string doUsingList( ast::UsingList const& astNode ) override;
    virtual std::string doParentClass( ast::ClassNode const& classAstNode ) override;
	std::set<std::string> readNamespaceFile( std::string namespaceFile );
	void addNamespaceToFile( std::string const& namespaceFile );
	void processNamespaces();
    std::string generateJsDocProperties() const;
    virtual std::string doClass( ast::NamespaceNode const& astNamespace, ast::ClassNode const& astNode ) override;
    virtual std::string doArgument( ast::ArgumentNode const& astNode ) override;

    virtual std::string doConst( ast::ConstNode const& constNode ) override;
    virtual std::string doMethod( ast::MethodNode const& astNode ) override;
    virtual std::string doProperty( ast::PropertyNode const& astNode ) override;
	virtual std::string doShortProperty(ast::ShortPropertyNode const& astNode) override;

    virtual std::string visitLiteralString( std::string const& value ) override;
    void generateJsDocProperty(const std::string& type, const std::string& name);
    virtual bool outputClass( std::string const& classSource, ast::ClassNode const& classAstNode ) override;

	std::string getTypeFromAttribute(const ast::AttributeListNode& attributes);
    virtual boost::filesystem::path getFileOutputFolder() override;

    std::string doSingleProperty( ast::PropertyNode const& astNode );
    std::string doSingleProperty( ast::ShortPropertyNode const& astNode );

    std::string doIndexerProperty( ast::PropertyNode const& astNode );

    virtual std::string visitAttrOptionalParams( ast::AttrOptionalParams const& astNode ) override;

	//std::string getPropertyType(ast::PropertyType const& propertyType, pam::pdl::symbols::SymbolPtr const& typeSymbol );

	virtual std::string visitAttrRequiredParams( ast::AttrRequiredParams const& astNode ) override;
    virtual std::string visitAttrRequiredAndOptionals( ast::AttrRequiredAndOptionals const& astNode ) override;

	AttributeInfo processAttributeName(ast::FullIdentifierNode const& attrName) override;

	std::string getJsPropertyType( ast::PropertyNode const& propertyNode );
	std::string getJsPropertyType( ast::ShortPropertyNode const& propertyNode );


private:
	struct JsCodeFormatters
	{
		std::string jsDocType;
		std::string jsDocVar;
		std::string jsDocParam;
		std::string jsDocClass;
		std::string jsDocReturns;

	};

	JsCodeFormatters _jsFormatter;

    std::set<std::string> _usingLines;
    std::string _namespaceSimulation;
    std::vector<std::string> _namespaceLines;
	std::vector<std::string> _propertyAttributes;
	std::vector<std::string> _jsDocProperties;

	std::string _fullParentClass;
	std::string _classInNamespace;
    bool _generateAsObject;
    std::string _variableDeclarationKeyword;

    std::string generateJsDocClass( ast::ClassNode const& classNode ) const;

	std::string internalGetJsPropertyType( pam::pdl::symbols::PropertySymbolPtr const& propertySymbol, ast::PropertyType const& propertyType );
};

}}};
