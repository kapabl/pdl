#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "javaGenerator.h"

using namespace pam::pdl;
using namespace symbols;
using namespace codegen;
using namespace config;
using namespace ast;

using std::string;
namespace fs = boost::filesystem;

JavaGenerator::JavaGenerator( PdlConfig const& config ):
    Generator( config, "java" )
{
    _fileExt = ".java";
    _classTemplate += _fileExt;

    _formatter.attr = "@%1%(%2%)";
    _formatter.paramlessAttr = "@%1%";
    _formatter.singleField = "private %1% %2%;";
    _formatter.indexerField = "private HashMap<%3%,%1%> %2%;";

    _formatter.singleGetProperty = "\n%4% %1% get%2%() { return this.%3%; }";
    _formatter.singleSetProperty = "\n%4% void set%2%( %1% value ) { this.%3% = value; }";

    _formatter.indexerGetProperty = "\n%6% %1% get%2%( %4% %5% ) { return this.%3%.get( %5% ); }";
    _formatter.indexerSetProperty = "\n%6% void set%2%( %4% %5%, %1% value ) { this.%3%.put( %5%, value ); }";

    _typesMap["object"] = "Object";
    _typesMap["function" ] = "Object";
    _typesMap["string"] = "String";
    _typesMap["bool"] = "boolean";
    _typesMap["double"] = "double";
    _typesMap["uint"] = "long";
    _typesMap["array"] = "Object[]";

    _emptyValueOfType["Boolean"] = "false";
    _emptyValueOfType["String"] = "\"\"";
    _emptyValueOfType["int"] = "0";
    _emptyValueOfType["double"] = "0.0";
    _emptyValueOfType["Object"] = "null";
    _emptyValueOfType["Object[]"] = "null";


}

string JavaGenerator::doUsingList( UsingList const& astNode )
{
    for( UsingNode const& UsingNode : astNode )
    {
        _importLines.insert( string("import ") + SymbolTable::joinIdentifier( UsingNode.className ) + ";" );
    }
    string result = boost::join( _importLines, "\n" );
    return result;
}

string JavaGenerator::doParentClass( ClassNode const& classAstNode )
{
    string result = "";
    if ( classAstNode.parentClass )
    {
        auto pair = SymbolTable::parseFullClassName( classAstNode.parentClass.get() );
        result = string("extends ") + pair.second;
    }
    return result;
}

string JavaGenerator::doClass( NamespaceNode const& astNamespace, ClassNode const& classAstNode )
{
    _importLines.clear();

    std::vector<std::string> classAttributes;
    classAttributes.push_back( getAccessModifier( classAstNode.accessModifier ) );
    auto classAttrs = boost::join( classAttributes, " " );

    auto inheritance = doParentClass( classAstNode );
    
    doMembers( classAstNode.members );
    auto propertyBlock = boost::join( _properties, "\n\n" );
    boost::replace_all( propertyBlock, "\n", "\n" + _indent + _indent );

    auto methodBlock = boost::join( _methods, "\n\n" );
    boost::replace_all( methodBlock, "\n", "\n" + _indent + _indent );

    auto constBLock = boost::join( _consts, "\n\n" );
    boost::replace_all( constBLock, "\n", "\n" + _indent + _indent );

    auto importBlock = doUsingList( astNamespace.usings );

    auto templateCode = readTemplateCode();


    //boost::replace
    string result;
    result = boost::replace_all_copy( templateCode, GENERIC_TEMPLATE_HEADER, _outputHeader );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_IMPORT_BLOCK, importBlock );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_PACKAGE, _mainNamespace );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_CLASS_ATTRS, classAttrs );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_CLASS_NAME, classAstNode.name.name );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_CLASS_INHERITANCE, inheritance );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_CONSTS, constBLock );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_PROPERTIES, propertyBlock );
    result = boost::replace_all_copy( result, JAVAGEN_TEMPLATE_METHODS, methodBlock );
        
    return result;
}

string JavaGenerator::doProperty( PropertyNode const& astNode )
{
    string result = astNode.arguments.size() == 0 
        ? doSingleProperty( astNode )
        : doIndexerProperty( astNode );

    return result;
}



string JavaGenerator::visitLiteralString( string const& value )
{
    string javaScapedString = boost::replace_all_copy( value, "\"", "\\\"");
    return "\"" + javaScapedString + "\"";
}


string JavaGenerator::doSingleProperty( PropertyNode const& propertyNode)
{
	const auto type = getPropertyType(propertyNode.propertyType);
	const auto name = propertyNode.name.name;
	const auto fieldName = fieldNameFromPropertyName( name );

    string getMethod;
    string setMethod;

	const auto accessModifier = getAccessModifier(propertyNode.accessModifier );

	const auto paAccess = propertyNode.access.get();
    if ( paAccess == PropertyAccess::paRead || paAccess == PropertyAccess::paReadWrite )
    {
        getMethod = (boost::format( _formatter.singleGetProperty ) 
            % type 
            % name 
            % fieldName
            % accessModifier).str();
    }

    if ( paAccess == PropertyAccess::paWrite || paAccess == PropertyAccess::paReadWrite )
    {
        setMethod = (boost::format( _formatter.singleSetProperty ) 
            % type 
            % name 
            % fieldName
            % accessModifier).str();
    }

	const auto propertyAttributes = doAttributes(propertyNode.attributes );

	auto result = 
        singleFieldType( type, fieldName ) + "\n\n" +
        propertyAttributes + 
        getMethod + 
        setMethod;
    
    return result;
}

string JavaGenerator::doIndexerProperty( PropertyNode const& propertyNode)
{
	const auto type = getPropertyType(propertyNode.propertyType);
	const auto name = propertyNode.name.name;

	auto const& argument1 = propertyNode.arguments.front();
	const auto argument1Name = argument1.name.name;
	const auto argument1Type = translateType( SymbolTable::joinIdentifier( argument1.type ) );

	const auto fieldName = fieldNameFromPropertyName( name );

    string getMethod;
    string setMethod;

	const auto accessModifier = getAccessModifier(propertyNode.accessModifier );

	const auto paAccess = propertyNode.access.get();
    if ( paAccess == PropertyAccess::paRead || paAccess == PropertyAccess::paReadWrite )
    {
        getMethod = (boost::format( _formatter.indexerGetProperty ) 
            % type
            % name
            % fieldName 
            % argument1Type
            % argument1Name
            % accessModifier ).str();
    }

    if ( paAccess == PropertyAccess::paWrite || paAccess == PropertyAccess::paReadWrite )
    {
        setMethod = (boost::format( _formatter.indexerSetProperty ) 
            % type
            % name
            % fieldName 
            % argument1Type
            % argument1Name
            % accessModifier ).str();
    }

	const auto propertyAttributes = doAttributes(propertyNode.attributes );
	auto result = 
        indexerFieldType( type, fieldName, argument1Type ) + "\n\n" +
        propertyAttributes + 
        getMethod +
        setMethod;
    
    return result;

}

string JavaGenerator::doArgument( ArgumentNode const& astNode )
{
    return translateType( SymbolTable::joinIdentifier( astNode.type ) ) + " " + astNode.name.name;
}


string JavaGenerator::doMethod( MethodNode const& astNode )
{
    string result = internalGenerateMethod(
        getAccessModifier( astNode.accessModifier ),
        translateType( SymbolTable::joinIdentifier( astNode.type ) ),
        astNode.name.name,
        "abstract",
        astNode.arguments );

    return result;
}

string JavaGenerator::doConst( ConstNode const& constNode )
{
    //TODO
    string result = "";

    return result;
}


bool JavaGenerator::outputClass( string const& classSource, ClassNode const& classAstNode )
{
    auto className = classAstNode.name.name;
    auto fullClassName = _mainNamespace + "." + className;

    auto result = Generator::outputClass( classSource, fullClassName );

    return result;
}

fs::path JavaGenerator::getFileOutputFolder()
{
    return fs::path( _outputFolder ) / "java/src" / boost::replace_all_copy( _mainNamespace, ".", "/" );
}

string JavaGenerator::internalGenerateMethod( 
        string const& accessModifier,
        string const& javaType,
        string const& name, 
        string const& virtuality,
        pam::pdl::ast::ArgumentList const& ArgumentList
    )
{
    auto formatter = boost::format("%1% %2% %5% %3%( %4% );")
        % accessModifier
        % virtuality
        % name
        % doArgumentList( ArgumentList )
        % javaType;


    string result = formatter.str();

    return result;
}

Generator::AttributeInfo JavaGenerator::processAttributeName( FullIdentifierNode const& attrName )
{
    Generator::AttributeInfo result;
    result.classInfo = SymbolTable::parseFullClassName( attrName );
    result.name = result.classInfo.second;
    return result;
}
