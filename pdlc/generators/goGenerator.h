#pragma once

namespace pam { namespace pdl { namespace codegen
{

class GoGenerator : public Generator
{
public:
    explicit GoGenerator( config::PdlConfig const& config );

protected:
    std::string doUsingList( ast::UsingList const& astNode ) override;
    std::string doClass( ast::NamespaceNode const& astNamespace, ast::ClassNode const& astNode ) override;
    std::string doParentClass( ast::ClassNode const& classAstNode ) override;
    std::string doArgument( ast::ArgumentNode const& astNode ) override;
    std::string doConst( ast::ConstNode const& constNode ) override;
    std::string doMethod( ast::MethodNode const& astNode ) override;
    std::string doProperty( ast::PropertyNode const& astNode ) override;
    std::string doShortProperty( ast::ShortPropertyNode const& astNode ) override;
    AttributeInfo processAttributeName( ast::FullIdentifierNode const& attrName ) override;
    std::string visitLiteralString( std::string const& value ) override;
    bool outputClass( std::string const& classSource, ast::ClassNode const& astClass ) override;
    boost::filesystem::path getFileOutputFolder() override;

private:
    std::string buildImportsBlock( std::set<std::string> const& imports ) const;
    std::string buildConstBlock() const;
    std::string buildStructBlock( std::string const& structName, std::string const& embedded, std::vector<std::string> const& fields ) const;
    std::string buildMethodsBlock() const;
    std::string goExportName( std::string const& identifier ) const;
    std::string goFieldTag( std::string const& name ) const;
    std::string goPackageName() const;
    std::string translateFullIdentifier( ast::FullIdentifierNode const& identifier );
    std::string buildPropertyField( std::string const& name, ast::PropertyType const& type );
    std::string defaultReturnValue( std::string const& type ) const;

    std::set<std::string> _imports;
    std::string _currentStructName;
};

}}};
