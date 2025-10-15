#pragma once

#include "commonHeaders.hpp"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "pdlConfig.h"
#include "symbolTable.h"

namespace io {
namespace pdl {
namespace astjson {

class AstJsonWriter {
public:
    explicit AstJsonWriter(config::PdlConfig& compilerConfig);
    void write(ast::NamespaceNode const& namespaceNode);

private:
    struct MemberJsonVisitor;

    boost::filesystem::path buildOutputPath() const;
    web::json::value buildDocument(ast::NamespaceNode const& namespaceNode) const;
    web::json::value buildSource(ast::NamespaceNode const& namespaceNode) const;
    web::json::value buildNamespace(ast::NamespaceNode const& namespaceNode) const;
    web::json::value buildUsings(ast::UsingList const& usingList) const;
    web::json::value buildClasses(ast::ClassList const& classList) const;
    web::json::value buildClass(ast::ClassNode const& classNode) const;
    web::json::value buildMembers(ast::MemberList const& memberList) const;
    web::json::value buildProperty(ast::PropertyNode const& propertyNode) const;
    web::json::value buildShortProperty(ast::ShortPropertyNode const& propertyNode) const;
    web::json::value buildMethod(ast::MethodNode const& methodNode) const;
    web::json::value buildConst(ast::ConstNode const& constNode) const;
    web::json::value buildAttributes(ast::AttributeListNode const& attributes) const;
    web::json::value buildAttribute(ast::AttributeNode const& attributeNode) const;
    web::json::value buildAttributeParams(ast::AttrParams const& params) const;
    web::json::value buildOptionalAttributeParams(ast::AttrOptionalParams const& params) const;
    web::json::value buildRequiredAttributeParams(ast::AttrRequiredParams const& params) const;
    web::json::value buildAttributeOptionalParam(ast::AttrOptionalParam const& paramNode) const;
    web::json::value buildArguments(ast::ArgumentList const& argumentList) const;
    web::json::value buildArgument(ast::ArgumentNode const& argumentNode) const;
    web::json::value buildPropertyType(ast::PropertyType const& propertyType) const;
    web::json::value buildIdentifier(ast::FullIdentifierNode const& identifierNode) const;
    web::json::value buildIdentifierSegments(ast::FullIdentifierNode const& identifierNode) const;
    web::json::value buildArrayNotation(ast::ArrayNotationList const& arrayNotationList) const;
    web::json::value buildLiteral(ast::LiteralValueNode const& literalNode) const;
    web::json::value buildMemberHeader(std::string const& kindName, std::string const& memberName, AccessModifiers accessModifier) const;
    utility::string_t toUtilityString(std::string const& value) const;
    std::string accessModifierName(AccessModifiers accessModifier) const;
    std::string propertyAccessName(boost::optional<PropertyAccess> const& accessKind) const;

    boost::filesystem::path _outputDirectory;
    boost::filesystem::path _projectRoot;
    std::string _inputFilePath;
};

struct AstJsonWriter::MemberJsonVisitor {
    explicit MemberJsonVisitor(AstJsonWriter const& writerReference);
    web::json::value operator()(ast::PropertyNode const& propertyNode) const;
    web::json::value operator()(ast::ShortPropertyNode const& propertyNode) const;
    web::json::value operator()(ast::MethodNode const& methodNode) const;
    web::json::value operator()(ast::ConstNode const& constNode) const;

private:
    AstJsonWriter const& writer;
};

} // namespace astjson
} // namespace pdl
} // namespace io
