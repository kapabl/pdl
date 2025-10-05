#pragma once

#define CSGEN_TEMPLATE_USING_BLOCK "[using-block]"
#define CSGEN_TEMPLATE_NAMESPACE "[namespace-name]"
#define CSGEN_TEMPLATE_CLASS_ATTRS "[class-attrs]"
#define CSGEN_TEMPLATE_CLASS_NAME "[class-name]"
#define CSGEN_TEMPLATE_CLASS_INHERITANCE "[inheritance]"
#define CSGEN_TEMPLATE_METHODS "[method-list]"
#define CSGEN_TEMPLATE_PROPERTIES "[property-list]"
#define CSGEN_TEMPLATE_CONSTS "[const-list]"

namespace pam { namespace pdl { namespace codegen
{

class CSharpGenerator : public Generator
{
public:
    CSharpGenerator( pam::pdl::config::PdlConfig const& config );

protected:
    virtual std::string doUsingList( pam::pdl::ast::UsingList const& astNode ) override;
    virtual std::string doClass( pam::pdl::ast::NamespaceNode const& astNamespace, pam::pdl::ast::ClassNode const& astNode ) override;
    virtual std::string doParentClass( pam::pdl::ast::ClassNode const& classAstNode ) override;
    virtual std::string doArgument( pam::pdl::ast::ArgumentNode const& astNode ) override;

    virtual std::string doConst( ast::ConstNode const& constNode ) override;
    virtual std::string doMethod( pam::pdl::ast::MethodNode const& astNode ) override;
    virtual std::string doProperty( pam::pdl::ast::PropertyNode const& astNode ) override;
	virtual std::string doShortProperty(pam::pdl::ast::ShortPropertyNode const& shortPropertyNode) override;

	virtual AttributeInfo processAttributeName( pam::pdl::ast::FullIdentifierNode const& attrName ) override;

    virtual std::string visitLiteralString( std::string const& value ) override;
    virtual bool outputClass( std::string const& classSource, pam::pdl::ast::ClassNode const& classAstNode ) override;
    virtual boost::filesystem::path getFileOutputFolder() override;

    std::string doSingleProperty( pam::pdl::ast::PropertyNode const& astNode );
	std::string doSingleProperty(pam::pdl::ast::ShortPropertyNode const& astNode);

    std::string doSingleLongProperty( pam::pdl::ast::PropertyNode const& astNode );
    std::string doSingleShortProperty( pam::pdl::ast::PropertyNode const& astNode );

    std::string doSingleLongProperty( pam::pdl::ast::ShortPropertyNode const& astNode );
    std::string doSingleShortProperty( pam::pdl::ast::ShortPropertyNode const& astNode );

    std::string doIndexerProperty( pam::pdl::ast::PropertyNode const& astNode );
    //std::string doIndexerLongProperty( pam::pdl::ast::PropertyNode const& astNode );
    //std::string doIndexerShortProperty( pam::pdl::ast::PropertyNode const& astNode );

    std::string internalGenerateMethod(
        std::string const& accessModifier,
        std::string const& type,
        std::string const& name, 
        std::string const& virtuality,
        pam::pdl::ast::ArgumentList const& astNode
    );

    std::set<std::string> _usingLines;

    web::json::value _jsonCsharpConfig;
    bool _emmitShortProperties;


};

}}};