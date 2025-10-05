#pragma once

#define JAVAGEN_TEMPLATE_IMPORT_BLOCK "[import-block]"
#define JAVAGEN_TEMPLATE_PACKAGE "[package-name]"
#define JAVAGEN_TEMPLATE_CLASS_ATTRS "[class-attrs]"
#define JAVAGEN_TEMPLATE_CLASS_NAME "[class-name]"
#define JAVAGEN_TEMPLATE_CLASS_INHERITANCE "[inheritance]"
#define JAVAGEN_TEMPLATE_METHODS "[method-list]"
#define JAVAGEN_TEMPLATE_PROPERTIES "[property-list]"
#define JAVAGEN_TEMPLATE_CONSTS "[const-list]"

namespace pam { namespace pdl { namespace codegen
{

class JavaGenerator : public Generator
{
public:
	explicit JavaGenerator( config::PdlConfig const& config );


protected:
	std::string doUsingList( ast::UsingList const& astNode ) override;
	std::string doParentClass( ast::ClassNode const& classAstNode ) override;
	std::string doClass( ast::NamespaceNode const& astNamespace, ast::ClassNode const& astNode ) override;
	std::string doArgument( ast::ArgumentNode const& astNode ) override;

    virtual std::string doConst( ast::ConstNode const& constNode ) override;
	std::string doMethod( ast::MethodNode const& astNode ) override;
	std::string doProperty( ast::PropertyNode const& astNode ) override;

	std::string visitLiteralString( std::string const& value ) override;
	bool outputClass( std::string const& classSource, ast::ClassNode const& classAstNode ) override;

	AttributeInfo processAttributeName( ast::FullIdentifierNode const& attrName ) override;

	boost::filesystem::path getFileOutputFolder() override;

    std::string doSingleProperty( ast::PropertyNode const& astNode );
    std::string doIndexerProperty( ast::PropertyNode const& astNode );

    std::string internalGenerateMethod(
        std::string const& accessModifier,
        std::string const& type,
        std::string const& name, 
        std::string const& virtuality,
        ast::ArgumentList const& astNode
    );

    

private:

    std::set<std::string> _importLines;

};

}}};