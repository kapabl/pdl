#pragma once

#define CPPGEN_TEMPLATE_USING_BLOCK "[using-block]"
#define CPPGEN_TEMPLATE_NAMESPACE "[namespace-name]"
#define CPPGEN_TEMPLATE_CLASS_ATTRS "[class-attrs]"
#define CPPGEN_TEMPLATE_CLASS_NAME "[class-name]"
#define CPPGEN_TEMPLATE_CLASS_INHERITANCE "[inheritance]"
#define CPPGEN_TEMPLATE_METHODS "[method-list]"
#define CPPGEN_TEMPLATE_PROPERTIES "[property-list]"
#define CPPGEN_TEMPLATE_CONSTS "[const-list]"

namespace pam { namespace pdl { namespace codegen
{

class CppGenerator : public Generator
{
public:
    CppGenerator( pam::pdl::config::PdlConfig const& config );

protected:
    virtual std::string doUsingList( ast::UsingList const& astNode ) override;
    virtual std::string doParentClass( ast::ClassNode const& classAstNode ) override;
    virtual std::string doClass( ast::NamespaceNode const& astNamespace, ast::ClassNode const& astNode ) override;
    virtual std::string doArgument( ast::ArgumentNode const& astNode ) override;

    virtual std::string doMethod( ast::MethodNode const& astNode ) override;
    virtual std::string doConst( ast::ConstNode const& constNode ) override;
    virtual std::string doProperty( ast::PropertyNode const& astNode ) override;

    virtual std::string visitLiteralString( std::string const& value ) override;
    virtual bool outputClass( std::string const& classSource, ast::ClassNode const& classAstNode ) override;

    virtual AttributeInfo processAttributeName( ast::FullIdentifierNode const& attrName ) override;
    virtual boost::filesystem::path getFileOutputFolder() override;

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

    std::set<std::string> _usingLines;

};

}}};