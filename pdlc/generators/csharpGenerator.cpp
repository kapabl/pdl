#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "csharpGenerator.h"

using namespace pam::pdl;
using namespace pam::pdl::config;
using namespace pam::pdl::symbols;
using namespace pam::pdl::codegen;
using namespace pam::pdl::ast;

using std::string;
namespace fs = boost::filesystem;


CSharpGenerator::CSharpGenerator( PdlConfig const& config ):
    Generator( config, "csharp" )
{
    _fileExt = ".cs";

    _emmitShortProperties = PdlConfig::as_bool( _languageConfig, "shortProperties" );

    _classTemplate += _fileExt;

    _formatter.singleField = "private %1% %2%;";
    _formatter.indexerField = "private Dictionary<%3%,%1%> %2%;";
    _formatter.constMember = "public const %2% %1% = %3%;";

    _formatter.singleGetProperty = "\n" + _indent + "get { return %1%; }";
    _formatter.singleSetProperty = "\n" + _indent + "set { %1% = value; }";

    _formatter.indexerGetProperty = "\n" + _indent + "get { return %1%[ %2% ]; }";
    _formatter.indexerSetProperty = "\n" + _indent + "set { %1%[ %2% ] = value; }";

    
    
    _typesMap["object"] = "object";
    _typesMap["function"] = "Action";

    //nullable CSharp types
    _typesMap["int"] = "int?";
    _typesMap["uint"] = "uint?";
    _typesMap["double"] = "double?";
    _typesMap["string"] = "string";
    _typesMap["bool"] = "bool?";
    
    _typesMap["array"] = "object[]";

    _emptyValueOfType["bool"] = "false";
    _emptyValueOfType["string"] = "\"\"";
    _emptyValueOfType["int"] = "0";
    _emptyValueOfType["double"] = "0.0";
    _emptyValueOfType["object"] = "null";
    _emptyValueOfType["function"] = "null";
    _emptyValueOfType["object[]"] = "null";
   

}

bool CSharpGenerator::outputClass( string const& classSource, ClassNode const& classAstNode )
{
	auto fullClassName = _mainNamespace + "." + classAstNode.name.name;
	auto result = Generator::outputClass( classSource, fullClassName );
    return result;
}

fs::path CSharpGenerator::getFileOutputFolder()
{
    //return fs::path( _outputFolder ) / "csharp";
    return fs::path( _outputFolder ) / "csharp/src" / boost::replace_all_copy( _mainNamespace, ".", "/" );
}

string CSharpGenerator::doUsingList( UsingList const& astNode )
{
    //std::set<string> usingLines;

    for( auto const& UsingNode : astNode )
    {
        auto pair = SymbolTable::parseFullClassName( UsingNode.className );
        if ( !pair.first.empty() )
        {
            _usingLines.insert( "using " + pair.first + ";" );
        }
    }
    string result = boost::join( _usingLines, "\n" );
    return result;
}

string CSharpGenerator::doParentClass( ClassNode const& classAstNode )
{
    string result = "";
    if ( classAstNode.parentClass )
    {
        auto pair = SymbolTable::parseFullClassName( classAstNode.parentClass.get() );
        result = ": " + pair.second;
    }
    return result;
}

string CSharpGenerator::doClass( NamespaceNode const& astNamespace, ClassNode const& classAstNode )
{
    _usingLines.clear();

    std::vector<std::string> attributes;
    attributes.push_back( getAccessModifier( classAstNode.accessModifier ) );
	auto classAttrs = boost::join( attributes, " " );

	auto inheritance = doParentClass( classAstNode );

    doMembers( classAstNode.members );

	auto constBlock = boost::join( _consts, "\n\n" );
    boost::replace_all( constBlock, "\n", "\n" + _indent + _indent );

	auto propertyBlock = boost::join( _properties, "\n\n" );
    boost::replace_all( propertyBlock, "\n", "\n" + _indent + _indent );

	auto methodBlock = boost::join( _methods, "\n\n" );
    boost::replace_all( methodBlock, "\n", "\n" + _indent + _indent );

	auto usingBlock = doUsingList( astNamespace.usings );

	auto templateCode = readTemplateCode();

    string result;
    result = boost::replace_all_copy( templateCode, GENERIC_TEMPLATE_HEADER, _outputHeader );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_USING_BLOCK, usingBlock );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_NAMESPACE, _mainNamespace );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_CLASS_ATTRS, classAttrs );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_CLASS_NAME, classAstNode.name.name );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_CLASS_INHERITANCE, inheritance );

    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_CONSTS, constBlock );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_PROPERTIES, propertyBlock );
    result = boost::replace_all_copy( result, CSGEN_TEMPLATE_METHODS, methodBlock );
        
    return result;
}

string CSharpGenerator::doProperty( PropertyNode const& astNode )
{
	auto result = astNode.arguments.size() == 0 
        ? doSingleProperty( astNode )
        : doIndexerProperty( astNode );

    return result;
}

string CSharpGenerator::doShortProperty(ShortPropertyNode const& shortPropertyNode)
{
	auto result = doSingleProperty(shortPropertyNode);
	return result;
}

std::string CSharpGenerator::doSingleLongProperty( ShortPropertyNode const& shortPropertyNode)
{
	const auto type = getPropertyType(shortPropertyNode.propertyType );

	const auto name = toPascalString( shortPropertyNode.name.name );
	const auto fieldName = fieldNameFromPropertyName( name );

	const auto fieldDeclaration = singleFieldType( type, fieldName );

	auto getMethod = (boost::format(_formatter.singleGetProperty) % fieldName).str();
	auto setMethod = (boost::format(_formatter.singleSetProperty) % fieldName).str();


    auto accessFormatter = boost::format("%1% %2% %3%\n{%4%%5%\n}")
        % getAccessModifier(AccessModifiers::amPublic )
        % type
        % name
        % getMethod
        % setMethod;

	const auto accessMethods = accessFormatter.str();

	auto result = fieldDeclaration;

    result += "\n" + accessMethods;
    
    return result;
}

std::string CSharpGenerator::doSingleLongProperty( PropertyNode const& propertyNode)
{
	const auto type = getPropertyType( propertyNode.propertyType );

	const auto name = toPascalString( propertyNode.name.name );
	const auto fieldName = fieldNameFromPropertyName( name );

	const auto fieldDeclaration = singleFieldType( type, fieldName );

    string getMethod;
    string setMethod;

	const auto paAccess = propertyNode.access.get();
    if ( paAccess == PropertyAccess::paRead || paAccess == PropertyAccess::paReadWrite )
    {
        getMethod = (boost::format( _formatter.singleGetProperty ) % fieldName).str();
    }

    if ( paAccess == PropertyAccess::paWrite || paAccess == PropertyAccess::paReadWrite )
    {
        setMethod = (boost::format( _formatter.singleSetProperty ) % fieldName).str();
    }

    auto accessFormatter = boost::format("%1% %2% %3%\n{%4%%5%\n}")
        % getAccessModifier(propertyNode.accessModifier )
        % type
        % name
        % getMethod
        % setMethod;

	const auto accessMethods = accessFormatter.str();
	auto attributes = doAttributes(propertyNode.attributes );

	auto result = fieldDeclaration;

    if ( !attributes.empty() )
    {
        result += "\n" + attributes;
    }


    result += "\n" + accessMethods;
    
    return result;
}

std::string CSharpGenerator::doSingleShortProperty( pam::pdl::ast::PropertyNode const& propertyNode )
{
	//TODO check array type
	const auto type = getPropertyType(propertyNode.propertyType);
	const auto name = toPascalString(propertyNode.name.name );

    string getMethod;
    string setMethod;

	const auto paAccess = propertyNode.access.get();
    if ( paAccess == PropertyAccess::paRead || paAccess == PropertyAccess::paReadWrite )
    {
        getMethod = "get;";
    }

    if ( paAccess == PropertyAccess::paWrite || paAccess == PropertyAccess::paReadWrite )
    {
        setMethod = "set;";
    }

    auto accessFormatter = boost::format("%1% %2% %3% { %4% %5% }")
        % getAccessModifier(propertyNode.accessModifier )
        % type
        % name
        % getMethod
        % setMethod;

	const auto accessMethods = accessFormatter.str();
	auto attributes = doAttributes(propertyNode.attributes );

	auto result = !attributes.empty() ? ( doAttributes(propertyNode.attributes ) + "\n" ) : "";
    result += accessMethods;
    
    return result;
}

std::string CSharpGenerator::doSingleShortProperty(pam::pdl::ast::ShortPropertyNode const& shortPropertyNode)
{
	//TODO check array type
	const auto type = getPropertyType(shortPropertyNode.propertyType);
	const auto name = toPascalString( shortPropertyNode.name.name );

	

	auto accessFormatter = boost::format("%1% %2% %3% { %4% %5% }")
		% getAccessModifier( AccessModifiers::amPublic )
		% type
		% name
		% "get;"
		% "set;";

	auto result = accessFormatter.str();

	return result;
}

//std::string CSharpGenerator::doIndexerLongProperty( pam::pdl::ast::PropertyNode const& astNode )
//{
//    string result;
//    //TODO
//    return result;
//}

//std::string CSharpGenerator::doIndexerShortProperty( pam::pdl::ast::PropertyNode const& astNode )
//{
//    string result;
//    //TODO
//    return result;
//}

std::string CSharpGenerator::doSingleProperty( pam::pdl::ast::PropertyNode const& astNode )
{
	auto result = ( _emmitShortProperties ) 
		? doSingleShortProperty( astNode ) 
		: doSingleLongProperty( astNode );

    return result;
}

std::string CSharpGenerator::doSingleProperty( pam::pdl::ast::ShortPropertyNode const& shortPropertyNode)
{
	auto result = ( _emmitShortProperties ) 
		? doSingleShortProperty(shortPropertyNode)
		: doSingleLongProperty(shortPropertyNode);

    return result;
}

string CSharpGenerator::doIndexerProperty( pam::pdl::ast::PropertyNode const& propertyNode)
{
	const auto type = getPropertyType( propertyNode.propertyType );

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
        getMethod = (boost::format( _formatter.indexerGetProperty ) % fieldName % argument1Name ).str();
    }

    if ( paAccess == PropertyAccess::paWrite || paAccess == PropertyAccess::paReadWrite )
    {
        setMethod = (boost::format( _formatter.indexerSetProperty ) % fieldName % argument1Name ).str();
    }

    auto accessFormatter = boost::format("%1% %2% this[%3%]\n{%4%%5%\n}")
        % accessModifier
        % type 
        % doArgument( argument1 )
        % getMethod
        % setMethod;

	auto access = accessFormatter.str();

    //as of Microsoft: http://msdn.microsoft.com/en-us/library/2549tw02%28v=vs.71%29.aspx
    if ( name != "Item" )
    {
        auto indexerAttr = boost::format("[System.Runtime.CompilerServices.IndexerName(\"%1%\")]") % name;
        access = indexerAttr.str() + "\n" + access;
    }

	auto result = 
        indexerFieldType( type, fieldName, argument1Type ) + "\n\n" +
        doAttributes(propertyNode.attributes ) + "\n" +
        access;

    return result;
}

string CSharpGenerator::doMethod( MethodNode const& astNode )
{
    //In the case of C# most method should be virtual so 
    //you can override them in your project

	auto result = internalGenerateMethod(
        getAccessModifier( astNode.accessModifier ),
        //SymbolTable::joinIdentifier( astNode.type ),
        translateType( SymbolTable::joinIdentifier( astNode.type ) ),
        astNode.name.name,
        "abstract",
        astNode.arguments ) + ";" ;

    return result;
}

string CSharpGenerator::doConst( ConstNode const& constNode )
{

    const auto value = boost::apply_visitor( _literalVisitor, constNode.value );
    const auto type = translateType( SymbolTable::joinIdentifier( constNode.type ) );

    auto formatter = boost::format( _formatter.constMember )
        % constNode.name
        % type
        % value;

    auto result = formatter.str();

    return result;
}

string CSharpGenerator::doArgument( ArgumentNode const& astNode )
{
    //return SymbolTable::joinIdentifier( astNode.type ) + " " + astNode.name.name;
     return translateType( SymbolTable::joinIdentifier( astNode.type ) ) + " " + astNode.name.name;
}


string CSharpGenerator::internalGenerateMethod( string const& accessModifier,
        string const& type,
        string const& name, 
        string const& virtuality,
        pam::pdl::ast::ArgumentList const& ArgumentList
    )
{
    auto formatter = boost::format("%1% %2% %3% %4%( %5% )")
        % accessModifier
        % virtuality
        % type
        % name
        % doArgumentList( ArgumentList );

	auto result = formatter.str();

    return result;
}


string CSharpGenerator::visitLiteralString( string const& value )
{
	const auto validCSharpString = boost::replace_all_copy( value, "\"", "\"\"" );
    return "@\"" + validCSharpString + "\"";
}

Generator::AttributeInfo CSharpGenerator::processAttributeName( FullIdentifierNode const& attrName )
{
    AttributeInfo result;
    result.classInfo = SymbolTable::parseFullClassName( attrName );
    result.name = result.classInfo.second;

    //add namespace to using list
    if ( !result.classInfo.first.empty() )
    {
		auto line = string("using ") + result.classInfo.first + ";";
        _usingLines.insert( line );
    }
    return result;
}


