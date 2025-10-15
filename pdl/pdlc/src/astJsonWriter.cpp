#include "astJsonWriter.h"

using namespace io::pdl;
using namespace io::pdl::ast;
using namespace io::pdl::astjson;
using namespace io::pdl::config;
using namespace io::pdl::symbols;

AstJsonWriter::AstJsonWriter(PdlConfig& compilerConfig) {
    auto& jsonConfig = compilerConfig.getJsonConfig();
    auto outputValue = PdlConfig::to_string(jsonConfig[U("out")].as_string());
    boost::filesystem::path outputPath(outputValue);
    auto configFileKey = PdlConfig::to_wstring("configFile");
    if (jsonConfig.has_field(configFileKey)) {
        auto configPath = boost::filesystem::path(PdlConfig::to_string(jsonConfig[configFileKey].as_string()));
        if (!configPath.is_absolute()) {
            configPath = boost::filesystem::absolute(configPath);
        }
        auto configDir = configPath.parent_path();
        if (!configDir.empty()) {
            auto projectRoot = configDir.parent_path();
            if (!projectRoot.empty()) {
                _projectRoot = projectRoot;
            }
        }
    }
    if (_projectRoot.empty()) {
        auto parentOutput = outputPath.parent_path();
        if (!parentOutput.empty()) {
            _projectRoot = parentOutput;
        } else {
            _projectRoot = outputPath;
        }
    }
    auto astKey = PdlConfig::to_wstring("ast");
    boost::filesystem::path astOutput;
    if (jsonConfig.has_field(astKey)) {
        auto astConfig = jsonConfig[astKey];
        auto dirKey = PdlConfig::to_wstring("outputDir");
        if (astConfig.has_field(dirKey)) {
            astOutput = boost::filesystem::path(PdlConfig::to_string(astConfig[dirKey].as_string()));
        }
    }
    if (astOutput.empty()) {
        auto baseForAst = outputPath.parent_path();
        if (baseForAst.empty()) {
            baseForAst = outputPath;
        }
        astOutput = baseForAst / "ast";
    }
    _outputDirectory = astOutput;
    _inputFilePath = compilerConfig.getInputFileName();
}

utility::string_t AstJsonWriter::toUtilityString(std::string const& value) const {
    auto result = PdlConfig::to_wstring(value);
    return result;
}

std::string AstJsonWriter::accessModifierName(AccessModifiers accessModifier) const {
    std::string result;
    switch (accessModifier) {
        case AccessModifiers::amPublic:
            result = ACCESS_MOD_PUBLIC_STR;
            break;
        case AccessModifiers::amInternal:
            result = ACCESS_MOD_INTERNAL_STR;
            break;
        case AccessModifiers::amProtected:
            result = ACCESS_MOD_PROTECTED_STR;
            break;
        case AccessModifiers::amPrivate:
            result = ACCESS_MOD_PRIVATE_STR;
            break;
        default:
            result = "";
            break;
    }
    return result;
}

std::string AstJsonWriter::propertyAccessName(boost::optional<PropertyAccess> const& accessKind) const {
    std::string result;
    if (!accessKind) {
        result = "";
        return result;
    }
    if (accessKind.get() == PropertyAccess::paRead) {
        result = "read";
        return result;
    }
    if (accessKind.get() == PropertyAccess::paWrite) {
        result = "write";
        return result;
    }
    result = "readWrite";
    return result;
}

boost::filesystem::path AstJsonWriter::buildOutputPath() const {
    boost::filesystem::path result = _outputDirectory;
    boost::filesystem::path inputPath(_inputFilePath);
    auto parentPath = inputPath.parent_path();
    if (!parentPath.empty()) {
        auto relativeParent = parentPath;
        if (!_projectRoot.empty()) {
            auto candidate = parentPath.lexically_relative(_projectRoot);
            auto candidateText = candidate.string();
            if (!candidate.empty() && candidateText.find("..") == std::string::npos) {
                relativeParent = candidate;
            } else if (parentPath.is_absolute()) {
                relativeParent = parentPath.filename();
            }
        } else if (parentPath.is_absolute()) {
            relativeParent = parentPath.filename();
        }
        if (!relativeParent.empty()) {
            result /= relativeParent;
        }
    }
    auto stemName = inputPath.stem().string();
    result /= stemName + ".ast.json";
    return result;
}

web::json::value AstJsonWriter::buildDocument(NamespaceNode const& namespaceNode) const {
    web::json::value result = web::json::value::object();
    result[U("version")] = web::json::value::number(1);
    result[U("source")] = buildSource(namespaceNode);
    result[U("namespace")] = buildNamespace(namespaceNode);
    return result;
}

web::json::value AstJsonWriter::buildSource(NamespaceNode const& namespaceNode) const {
    web::json::value result = web::json::value::object();
    result[U("file")] = web::json::value::string(toUtilityString(_inputFilePath));
    result[U("namespace")] = buildIdentifier(namespaceNode.name);
    return result;
}

web::json::value AstJsonWriter::buildNamespace(NamespaceNode const& namespaceNode) const {
    web::json::value result = web::json::value::object();
    result[U("name")] = buildIdentifier(namespaceNode.name);
    result[U("usings")] = buildUsings(namespaceNode.usings);
    result[U("classes")] = buildClasses(namespaceNode.classes);
    return result;
}

web::json::value AstJsonWriter::buildUsings(UsingList const& usingList) const {
    web::json::value result = web::json::value::array();
    std::size_t usingIndex = 0;
    for (auto const& usingNode : usingList) {
        result[usingIndex++] = buildIdentifier(usingNode.className);
    }
    return result;
}

web::json::value AstJsonWriter::buildClasses(ClassList const& classList) const {
    web::json::value result = web::json::value::array();
    std::size_t classIndex = 0;
    for (auto const& classNode : classList) {
        result[classIndex++] = buildClass(classNode);
    }
    return result;
}

web::json::value AstJsonWriter::buildClass(ClassNode const& classNode) const {
    web::json::value result = web::json::value::object();
    result[U("name")] = web::json::value::string(toUtilityString(classNode.name.name));
    result[U("accessModifier")] = web::json::value::string(toUtilityString(accessModifierName(classNode.accessModifier)));
    if (classNode.parentClass) {
        result[U("parent")] = buildIdentifier(classNode.parentClass.get());
    } else {
        result[U("parent")] = web::json::value::null();
    }
    result[U("members")] = buildMembers(classNode.members);
    return result;
}

web::json::value AstJsonWriter::buildMembers(MemberList const& memberList) const {
    web::json::value result = web::json::value::array();
    MemberJsonVisitor visitor(*this);
    std::size_t memberIndex = 0;
    for (auto const& memberNode : memberList) {
        result[memberIndex++] = boost::apply_visitor(visitor, memberNode);
    }
    return result;
}

web::json::value AstJsonWriter::buildMemberHeader(std::string const& kindName, std::string const& memberName, AccessModifiers accessModifier) const {
    web::json::value result = web::json::value::object();
    result[U("kind")] = web::json::value::string(toUtilityString(kindName));
    result[U("name")] = web::json::value::string(toUtilityString(memberName));
    result[U("accessModifier")] = web::json::value::string(toUtilityString(accessModifierName(accessModifier)));
    return result;
}

web::json::value AstJsonWriter::buildProperty(PropertyNode const& propertyNode) const {
    web::json::value result = buildMemberHeader("property", propertyNode.name.name, propertyNode.accessModifier);
    result[U("type")] = buildPropertyType(propertyNode.propertyType);
    result[U("attributes")] = buildAttributes(propertyNode.attributes);
    result[U("arguments")] = buildArguments(propertyNode.arguments);
    result[U("access")] = web::json::value::string(toUtilityString(propertyAccessName(propertyNode.access)));
    return result;
}

web::json::value AstJsonWriter::buildShortProperty(ShortPropertyNode const& propertyNode) const {
    web::json::value result = buildMemberHeader("shortProperty", propertyNode.name.name, AccessModifiers::amNone);
    result[U("type")] = buildPropertyType(propertyNode.propertyType);
    result[U("attributes")] = web::json::value::array();
    result[U("arguments")] = web::json::value::array();
    result[U("access")] = web::json::value::string(toUtilityString(""));
    return result;
}

web::json::value AstJsonWriter::buildMethod(MethodNode const& methodNode) const {
    web::json::value result = buildMemberHeader("method", methodNode.name.name, methodNode.accessModifier);
    result[U("returnType")] = buildIdentifier(methodNode.type);
    result[U("arguments")] = buildArguments(methodNode.arguments);
    return result;
}

web::json::value AstJsonWriter::buildConst(ConstNode const& constNode) const {
    web::json::value result = buildMemberHeader("const", constNode.name.name, AccessModifiers::amNone);
    result[U("type")] = buildIdentifier(constNode.type);
    result[U("value")] = buildLiteral(constNode.value);
    return result;
}

web::json::value AstJsonWriter::buildArguments(ArgumentList const& argumentList) const {
    web::json::value result = web::json::value::array();
    std::size_t argumentIndex = 0;
    for (auto const& argumentNode : argumentList) {
        result[argumentIndex++] = buildArgument(argumentNode);
    }
    return result;
}

web::json::value AstJsonWriter::buildArgument(ArgumentNode const& argumentNode) const {
    web::json::value result = web::json::value::object();
    result[U("name")] = web::json::value::string(toUtilityString(argumentNode.name.name));
    result[U("type")] = buildIdentifier(argumentNode.type);
    return result;
}

web::json::value AstJsonWriter::buildPropertyType(PropertyType const& propertyType) const {
    web::json::value result = web::json::value::object();
    result[U("type")] = buildIdentifier(propertyType.type);
    result[U("arrayNotation")] = buildArrayNotation(propertyType.arrayNotationList);
    return result;
}

web::json::value AstJsonWriter::buildIdentifier(FullIdentifierNode const& identifierNode) const {
    web::json::value result = web::json::value::object();
    auto qualified = SymbolTable::joinIdentifier(identifierNode);
    result[U("segments")] = buildIdentifierSegments(identifierNode);
    result[U("qualifiedName")] = web::json::value::string(toUtilityString(qualified));
    return result;
}

web::json::value AstJsonWriter::buildIdentifierSegments(FullIdentifierNode const& identifierNode) const {
    web::json::value result = web::json::value::array();
    std::size_t segmentIndex = 0;
    for (auto const& identifier : identifierNode) {
        result[segmentIndex++] = web::json::value::string(toUtilityString(identifier.name));
    }
    return result;
}

web::json::value AstJsonWriter::buildArrayNotation(ArrayNotationList const& arrayNotationList) const {
    web::json::value result = web::json::value::array();
    std::size_t notationIndex = 0;
    for (auto const& notation : arrayNotationList) {
        result[notationIndex++] = web::json::value::string(toUtilityString(notation));
    }
    return result;
}

web::json::value AstJsonWriter::buildAttributes(AttributeListNode const& attributes) const {
    web::json::value result = web::json::value::array();
    std::size_t attributeIndex = 0;
    for (auto const& attributeNode : attributes) {
        result[attributeIndex++] = buildAttribute(attributeNode);
    }
    return result;
}

web::json::value AstJsonWriter::buildAttribute(AttributeNode const& attributeNode) const {
    web::json::value result = web::json::value::object();
    result[U("name")] = buildIdentifier(attributeNode.name);
    result[U("params")] = buildAttributeParams(attributeNode.params);
    return result;
}

web::json::value AstJsonWriter::buildAttributeParams(AttrParams const& params) const {
    web::json::value result = web::json::value::object();
    result[U("required")] = web::json::value::array();
    result[U("optional")] = web::json::value::array();
    if (auto optionalParams = boost::get<AttrOptionalParams>(&params)) {
        auto optionalValues = buildOptionalAttributeParams(*optionalParams);
        result[U("optional")] = optionalValues;
        return result;
    }
    if (auto requiredParams = boost::get<AttrRequiredParams>(&params)) {
        auto requiredValues = buildRequiredAttributeParams(*requiredParams);
        result[U("required")] = requiredValues;
        return result;
    }
    auto bothParams = boost::get<AttrRequiredAndOptionals>(&params);
    if (bothParams == nullptr) {
        return result;
    }
    auto requiredValues = buildRequiredAttributeParams(bothParams->required);
    auto optionalValues = buildOptionalAttributeParams(bothParams->optionals);
    result[U("required")] = requiredValues;
    result[U("optional")] = optionalValues;
    return result;
}

web::json::value AstJsonWriter::buildOptionalAttributeParams(AttrOptionalParams const& params) const {
    web::json::value result = web::json::value::array();
    std::size_t paramIndex = 0;
    for (auto const& paramNode : params) {
        result[paramIndex++] = buildAttributeOptionalParam(paramNode);
    }
    return result;
}

web::json::value AstJsonWriter::buildRequiredAttributeParams(AttrRequiredParams const& params) const {
    web::json::value result = web::json::value::array();
    std::size_t paramIndex = 0;
    for (auto const& paramValue : params) {
        result[paramIndex++] = buildLiteral(paramValue);
    }
    return result;
}

web::json::value AstJsonWriter::buildAttributeOptionalParam(AttrOptionalParam const& paramNode) const {
    web::json::value result = web::json::value::object();
    result[U("name")] = web::json::value::string(toUtilityString(paramNode.name.name));
    result[U("value")] = buildLiteral(paramNode.value);
    return result;
}

web::json::value AstJsonWriter::buildLiteral(LiteralValueNode const& literalNode) const {
    web::json::value result;
    if (auto stringValue = boost::get<std::string>(&literalNode)) {
        result = web::json::value::string(toUtilityString(*stringValue));
        return result;
    }
    if (auto doubleValue = boost::get<double>(&literalNode)) {
        result = web::json::value::number(*doubleValue);
        return result;
    }
    if (auto intValue = boost::get<int>(&literalNode)) {
        result = web::json::value::number(*intValue);
        return result;
    }
    auto boolValue = boost::get<bool>(literalNode);
    result = web::json::value::boolean(boolValue);
    return result;
}

AstJsonWriter::MemberJsonVisitor::MemberJsonVisitor(AstJsonWriter const& writerReference) : writer(writerReference) {
}

web::json::value AstJsonWriter::MemberJsonVisitor::operator()(PropertyNode const& propertyNode) const {
    auto result = writer.buildProperty(propertyNode);
    return result;
}

web::json::value AstJsonWriter::MemberJsonVisitor::operator()(ShortPropertyNode const& propertyNode) const {
    auto result = writer.buildShortProperty(propertyNode);
    return result;
}

web::json::value AstJsonWriter::MemberJsonVisitor::operator()(MethodNode const& methodNode) const {
    auto result = writer.buildMethod(methodNode);
    return result;
}

web::json::value AstJsonWriter::MemberJsonVisitor::operator()(ConstNode const& constNode) const {
    auto result = writer.buildConst(constNode);
    return result;
}

namespace {

std::string formatJson(std::string const& rawJson) {
    std::string result;
    result.reserve(rawJson.size() * 2);
    int indentLevel = 0;
    bool insideString = false;
    bool escaping = false;
    auto appendIndent = [&](int level) {
        result.append(static_cast<std::size_t>(level) * 2, ' ');
    };
    for (char ch : rawJson) {
        if (insideString) {
            result.push_back(ch);
            if (escaping) {
                escaping = false;
                continue;
            }
            if (ch == '\\') {
                escaping = true;
                continue;
            }
            if (ch == '"') {
                insideString = false;
            }
            continue;
        }
        switch (ch) {
            case '{':
            case '[':
                result.push_back(ch);
                result.push_back('\n');
                indentLevel++;
                appendIndent(indentLevel);
                break;
            case '}':
            case ']':
                result.push_back('\n');
                indentLevel = std::max(indentLevel - 1, 0);
                appendIndent(indentLevel);
                result.push_back(ch);
                break;
            case ',':
                result.push_back(ch);
                result.push_back('\n');
                appendIndent(indentLevel);
                break;
            case ':':
                result.push_back(ch);
                result.push_back(' ');
                break;
            case '"':
                result.push_back(ch);
                insideString = true;
                break;
            default:
                if (std::isspace(static_cast<unsigned char>(ch))) {
                    break;
                }
                result.push_back(ch);
                break;
        }
    }
    result.push_back('\n');
    return result;
}

} // namespace

void AstJsonWriter::write(NamespaceNode const& namespaceNode) {
    auto document = buildDocument(namespaceNode);
    auto outputPath = buildOutputPath();
    auto parentDirectory = outputPath.parent_path();
    if (!parentDirectory.empty()) {
        boost::filesystem::create_directories(parentDirectory);
    } else {
        boost::filesystem::create_directories(_outputDirectory);
    }
    auto rawSerialized = utility::conversions::to_utf8string(document.serialize());
    auto serialized = formatJson(rawSerialized);
    std::ofstream outputStream(outputPath.string(), std::ios::out | std::ios::trunc);
    outputStream << serialized;
    outputStream.close();
}
