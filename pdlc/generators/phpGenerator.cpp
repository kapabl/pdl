#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "phpGenerator.h"

using namespace pam::pdl;
using namespace pam::pdl::symbols;
using namespace pam::pdl::codegen;
using namespace pam::pdl::ast;

using std::string;
namespace fs = boost::filesystem;

using namespace pam::pdl::config;

PhpGenerator::PhpGenerator(PdlConfig const& config) :
    Generator(config, "php")
{
    _fileExt = ".php";
    _classTemplate += _fileExt;

    readPhpConfig();

    _formatter.paramlessAttr = "'%1%' => []";
    _formatter.attr = string("'%1%' => [\n%2%\n]");
    _formatter.constMember = "const %1% = %2%;";

    _typesMap["object"] = "object";
    _typesMap["function"] = "callable";
    _typesMap["string"] = "string";
    _typesMap["bool"] = "bool";
    _typesMap["double"] = "float";
    _typesMap["int"] = "int";
    _typesMap["uint"] = "int";
    _typesMap["array"] = "array";

    _emptyValueOfType["bool"] = "false";
    _emptyValueOfType["string"] = "''";
    _emptyValueOfType["int"] = "0";
    _emptyValueOfType["double"] = "0.0";
    _emptyValueOfType["array"] = "null";
    _emptyValueOfType["object"] = "null";
    _emptyValueOfType["function"] = "null";

    _attributeSeparator = ",\n";
}

std::string PhpGenerator::fullIdentifierToPsr4(FullIdentifierNode const& fullIdentifier)
{
    std::vector< std::string > vString;
    for (auto const& identifier : fullIdentifier)
    {
        auto singleNamespace = identifier.name;
        singleNamespace[0] = toupper(singleNamespace[0]);
        vString.push_back(singleNamespace);
    }

    auto result = boost::algorithm::join(vString, ".");
    return result;
}


std::string PhpGenerator::getTargetLanguageNamespace(ast::NamespaceNode const& namespaceNode)
{
    auto result = fullIdentifierToString(namespaceNode.name);
    return result;
}

bool PhpGenerator::isIgnoreNamespace(std::string const& namespaceString)
{
    const auto result = _ignoreNamespaces.find(namespaceString) !=
        _ignoreNamespaces.end();

    return result;
}

bool PhpGenerator::isInMainNamespace(std::string const& namespaceString) const
{
    const auto result = _pdlMainNamespace == namespaceString;
    return result;
}

string PhpGenerator::fullIdentifierToString(FullIdentifierNode const& fullIdentifier) const
{
    string result;

    if (_phpPsr4)
    {
        result = fullIdentifierToPsr4(fullIdentifier);
    }
    else
    {
        result = SymbolTable::joinIdentifier(fullIdentifier);
    }

    return result;
}

string PhpGenerator::doUsingList(UsingList const& usingList)
{
    for (auto const& usingNode : usingList)
    {
        if (usingNode.symbol->getUseCount() > 0)
        {
            const auto namespaceString = usingNode.symbol->getNamespace()->name();
            if (!isIgnoreNamespace(namespaceString) && !isInMainNamespace(namespaceString))
            {
                const auto fullClassName = boost::replace_all_copy(fullIdentifierToString(usingNode.className), ".",
                    "\\");
                _useLines.insert(string("use ") + fullClassName + ";");
            }
        }
    }
    auto result = boost::join(_useLines, "\n");
    return result;
}

string PhpGenerator::doShortProperty(ShortPropertyNode const& shortPropertyNode)
{
    auto result = doSingleProperty(shortPropertyNode);
    return result;
}

string PhpGenerator::doProperty(PropertyNode const& astNode)
{
    auto result = astNode.arguments.size() == 0
        ? doSingleProperty(astNode)
        : doIndexerProperty(astNode);

    return result;
}


string PhpGenerator::doParentClass(ClassNode const& classNode)
{
    string result = "";
    if (classNode.parentClass)
    {
        std::string parentClassname;
        auto const& parentClass = classNode.parentClass.get();
        const auto parsedClassname = SymbolTable::parseFullClassName(parentClass);
        if (parsedClassname.second != classNode.symbol->name())
        {
            parentClassname = parsedClassname.second;
        }
        else
        {
            auto const& parentClassSymbol = classNode.symbol->getParentClass();

            parentClassname = "\\" + fullClassNameToPsr4(parentClassSymbol->getNamespace()->name()
                + '.' + parentClassSymbol->name());
            parentClassSymbol->decUseCount();
        }

        result = string("extends ") + parentClassname;
    }
    return result;
}


string PhpGenerator::doClass(NamespaceNode const& astNamespace, ClassNode const& classNode)
{
    _useLines.clear();
    _propertyAttributes.clear();

    std::vector< std::string > classAttributes;
    classAttributes.push_back(getAccessModifier(classNode.accessModifier));
    const auto classAttrs = boost::join(classAttributes, " ");

    const auto phpNamespace = boost::replace_all_copy(_mainNamespace, ".", "\\");
    const auto className = classNode.name.name;
    const auto inheritance = doParentClass(classNode);

    std::string phpDocInheritance = " *";
    if (classNode.parentClass)
    {
        phpDocInheritance = (boost::format(" * @%1%") % inheritance).str();
    }

    const auto phpDocClass = (boost::format(" * @class %1%") % className).str();
    const auto phpDocPackage = (boost::format(" * @package %1%") % phpNamespace).str();

    doMembers(classNode.members);

    const auto propertyBlock = boost::join(_properties, "\n");
    const auto propertyAttributesBlock = generatePropertyAttributeBlock();
    string propertyAttributesVar;

    if (!_propertyAttributes.empty())
    {
        propertyAttributesVar = "private $_propertyAttributes;";
    }


    auto methodBlock = boost::join(_methods, "\n");
    boost::replace_all(methodBlock, "\n", "\n" + _indent);

    auto constBlock = boost::join(_consts, "\n");
    boost::replace_all(constBlock, "\n", "\n" + _indent);

    const auto useBlock = doUsingList(astNamespace.usings);

    const auto templateCode = readTemplateCode();

    const auto parentConstructor = classNode.parentClass
        ? "parent::__construct();\n"
        : "";

    const auto constructor = (boost::format(
        "public function __construct()\n\
    {\n\
        %1%\
        %2%\n\
    }\n"
        )
        % propertyAttributesBlock
        % parentConstructor).str();

    auto result = boost::replace_all_copy(templateCode, GENERIC_TEMPLATE_HEADER, _outputHeader);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_USE_BLOCK, useBlock);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_NAMESPACE, phpNamespace);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_PHPDOC_PACKAGE, phpDocPackage);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_CLASS_ATTRS, classAttrs);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_CLASS_NAME, className);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_CLASS_INHERITANCE, inheritance);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_PHPDOC_INHERITANCE, phpDocInheritance);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_PHPDOC_CLASS, phpDocClass);

    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_CONSTS, constBlock);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_PROPERTIES, propertyBlock);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_METHODS, methodBlock);

    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_PROPERTY_ATTRS_BLOCK, propertyAttributesBlock);
    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_PROPERTY_ATTRS_VAR, propertyAttributesVar);

    result = boost::replace_all_copy(result, PHPGEN_TEMPLATE_CONSTRUCTOR, constructor);


    return result;
}

string PhpGenerator::generatePropertyAttributeBlock() const
{
    string result;

    if (!_propertyAttributes.empty())
    {
        std::vector< string > attributeBlock;
        for (auto attribute : _propertyAttributes)
        {
            if (!attribute.second.empty())
            {
                //auto attrValues = _indent + boost::replace_all_copy(attribute.second, "\n", "\n" + _indent );
                attributeBlock.push_back((boost::format("'%1%' => [ %2% ]")
                    % attribute.first
                    % attribute.second).str());
            }
        }

        if (!attributeBlock.empty())
        {
            const auto attributes = boost::replace_all_copy(boost::join(attributeBlock, ",\n"), "\n",
                "\n" + _indent + _indent + _indent);
            result = "$this->_propertyAttributes = [\n"
                + _indent + _indent + _indent + attributes + "\n"
                + _indent + _indent + "];\n";
        }
    }

    return result;
}


string PhpGenerator::visitLiteralString(string const& value)
{
    const auto phpScapedString = boost::replace_all_copy(value, "\"", "\\\"");
    return "\"" + phpScapedString + "\"";
}

string PhpGenerator::doSinglePropertyBody(Identifier const& name, PropertyType const& propertyType)
{
    const auto type = translateType(SymbolTable::joinIdentifier(propertyType.type));
    const auto propertyName = toCamelString(name.name);

    const auto brackets = arrayNotationList2Brackets( propertyType.arrayNotationList );

    const string phpDocProperty = " * @property %1%%3% $%2%";

    auto result = (boost::format(phpDocProperty)
        % type
        % propertyName
        % brackets).str();

    return result;
}

string PhpGenerator::doSingleProperty(PropertyNode const& propertyNode)
{
    const auto name = toCamelString(propertyNode.name.name);
    _propertyAttributes[name] = doPropertyAttributes(propertyNode);

    auto result = doSinglePropertyBody(propertyNode.name, propertyNode.propertyType);

    return result;
}

string PhpGenerator::doSingleProperty(ShortPropertyNode const& shortPropertyNode)
{
    auto result = doSinglePropertyBody(shortPropertyNode.name, shortPropertyNode.propertyType);

    return result;
}

std::string PhpGenerator::fullClassNameToPsr4(std::string const& fullClassName)
{
    std::vector< std::string > parts;
    boost::split(parts, fullClassName, boost::is_any_of("."));

    for (auto& part : parts)
    {
        part[0] = toupper(part[0]);
    }

    auto result = boost::join(parts, "\\");

    return result;
}

string PhpGenerator::doIndexerProperty(PropertyNode const& propertyNode)
{
    const auto type = getPropertyType(propertyNode.propertyType);
    const auto name = toCamelString(propertyNode.name.name);

    auto const& argument1 = propertyNode.arguments.front();
    auto argument1Name = argument1.name.name;
    auto argument1Type = translateType(SymbolTable::joinIdentifier(argument1.type));

    _propertyAttributes[name] = doAttributes(propertyNode.attributes);

    const auto brackets = arrayNotationList2Brackets(propertyNode.propertyType.arrayNotationList);

    const string phpDocProperty = " * @property %1%%3% $%2%";

    auto result = (boost::format(phpDocProperty)
        % type
        % name
        % brackets
        ).str();

    return result;
}

string PhpGenerator::doArgument(ArgumentNode const& astNode)
{
    const auto type = translateType(SymbolTable::joinIdentifier(astNode.type));

    auto result = (boost::format("/*%1%*/ %2%") % type % astNode.name.name).str();

    return result;
}


string PhpGenerator::doMethod(MethodNode const& astNode)
{
    //public function checkPlayerLic(/*int*/ i, /*double*/ d ) /*int*/ { return 0; } 
    const auto retType = translateType(SymbolTable::joinIdentifier(astNode.type));

    const auto access = getAccessModifier(astNode.accessModifier);
    const auto arguments = doArgumentList(astNode.arguments);

    auto result = (boost::format("%3% function %2%( %4% ) /*%1%*/ { return %5%; }")
        % retType
        % astNode.name.name
        % access
        % arguments
        % getEmptyValueOfType(retType)
        ).str();

    return result;
}

string PhpGenerator::doConst( ConstNode const& constNode )
{

    const auto value = boost::apply_visitor( _literalVisitor, constNode.value );
    //const auto type = translateType( SymbolTable::joinIdentifier( constNode.type ) );

    auto formatter = boost::format( _formatter.constMember )
        % constNode.name
      //  % type
        % value;

    auto result = formatter.str();

    return result;
}


bool PhpGenerator::outputClass(string const& classSource, ClassNode const& classNode)
{
    const auto fullClassName = _mainNamespace + "." + classNode.name.name;
    const auto result = Generator::outputClass(classSource, fullClassName);
    return result;
}

fs::path PhpGenerator::getFileOutputFolder()
{
    const auto classNamespace = _mainNamespace;
    if (_phpPsr4)
    {
        //classNamespace = '';
    }

    return fs::path(_outputFolder) / "php" / boost::replace_all_copy(classNamespace, ".", "/");
}


Generator::AttributeInfo PhpGenerator::processAttributeName(FullIdentifierNode const& attrName)
{
    //custom attribute class is ignored in as3 and js
    Generator::AttributeInfo result;
    result.classInfo = SymbolTable::parseFullClassName(attrName);
    result.name = result.classInfo.second;
    return result;
}

string PhpGenerator::visitAttrRequiredParams(AttrRequiredParams const& astNode)
{
    string result;

    if (!astNode.empty())
    {
        std::vector< string > paramVector;
        auto index = 1;
        for (auto const& param : astNode)
        {
            paramVector.push_back("'default" + std::to_string(index) + "' => "
                + boost::apply_visitor(_literalVisitor, param));
            index++;
        }
        result = _indent + boost::join(paramVector, ",\n" + _indent);
    }
    return result;
}

string PhpGenerator::visitAttrRequiredAndOptionals(AttrRequiredAndOptionals const& astNode)
{
    auto result = visitAttrRequiredParams(astNode.required) + ",\n" + visitAttrOptionalParams(astNode.optionals);
    return result;
}

string PhpGenerator::visitAttrOptionalParams(AttrOptionalParams const& astNode)
{
    std::vector< string > paramVector;
    string result;
    if (!astNode.empty())
    {
        for (auto const& param : astNode)
        {
            paramVector.push_back("'" + param.name.name + "' => "
                + boost::apply_visitor(_literalVisitor, param.value));
        }
        result = _indent + boost::join(paramVector, ",\n" + _indent);
    }
    return result;
}

void PhpGenerator::readPhpConfig()
{
    _phpPsr4 = PdlConfig::as_bool(_languageConfig, "psr4");
    auto ignoreNamespaces = PdlConfig::as_array(_languageConfig, "ignoreNamespaces");
    auto iter = ignoreNamespaces.begin();
    while (iter != ignoreNamespaces.end())
    {
        auto jsonValue = *iter;
        _ignoreNamespaces.insert(PdlConfig::to_string(jsonValue.as_string()));
        ++iter;
    }
}
