#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "CppGenerator.h"

using namespace pam::pdl;
using namespace symbols;
using namespace codegen;
using namespace ast;

using std::string;
namespace fs = boost::filesystem;

CppGenerator::CppGenerator( config::PdlConfig const& config ) :
    Generator( config, "cpp" )
{
    _fileExt = ".h";
    _classTemplate += _fileExt;

    //TODO

//    _formatter.paramlessAttr = "[%1%]";
//    _formatter.singleField = "private var %2%: %1%;";
//    _formatter.indexerField = "private var %2%:Object = {};/* %3% -> %1% */";
/*
    _formatter.singleGetProperty = "\n%4% function get %2%():%1% { return this.%3%; }";
    _formatter.singleSetProperty = "\n%4% function set %2%( value: %1% ):void { this.%3% = value };";

    //there are no indexer or property with arguments in as3... so we use functions
    _formatter.indexerGetProperty = "\n%6% function get%2%( %5%: %4% ):%1% { return this.%3%[ %5% ]; }";
    _formatter.indexerSetProperty = "\n%6% function set%2%( %5%: %4%, value: %1% ):void { this.%3%[ %5% ] = value; }";

    _typesMap["object"] = "void*";
    _typesMap["uint"] = "unsignit int";
    _typesMap["function" ] = "void*";
    _typesMap["string"] = "std::string";
    _typesMap["bool"] = "bool";
    _typesMap["double"] = "double";
    _typesMap["array"] = "void**";

    _emptyValueOfType["bool"] = "false";
    _emptyValueOfType["std::string"] = "\"\"";
    _emptyValueOfType["int"] = "0";
    _emptyValueOfType["double"] = "0.0";
    _emptyValueOfType["array"] = "null";
*/


}

string CppGenerator::doUsingList( UsingList const& astNode )
{
    //TODO
    /*
    for( UsingNode const& UsingNode : astNode )
    {
        _importLines.insert( _indent + string("import ") + SymbolTable::joinIdentifier( UsingNode.className ) + ";" );
    }
    string result = boost::join( _importLines, "\n" );
    return result;
    */
    return "";
}

string CppGenerator::doParentClass( ClassNode const& classAstNode )
{
    //TODO
    /*
    string result = "";
    if ( classAstNode.parentClass )
    {
        auto pair = SymbolTable::parseFullClassName( classAstNode.parentClass.get() );
        result = string("extends ") + pair.second;
    }
    return result;
    */
    return "";
}

string CppGenerator::doClass( NamespaceNode const& astNamespace, ClassNode const& classAstNode )
{

    _usingLines.clear();


    std::vector<std::string> classAttributes;
    classAttributes.push_back( getAccessModifier( classAstNode.accessModifier ) );
    auto classAttrs = boost::join( classAttributes, " " );

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


    //boost::replace
    string result;
    result = boost::replace_all_copy( templateCode, GENERIC_TEMPLATE_HEADER, _outputHeader );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_USING_BLOCK, usingBlock );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_NAMESPACE, _mainNamespace );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_CLASS_ATTRS, classAttrs );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_CLASS_NAME, classAstNode.name.name );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_CLASS_INHERITANCE, inheritance );

    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_CONSTS, constBlock );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_PROPERTIES, propertyBlock );
    result = boost::replace_all_copy( result, CPPGEN_TEMPLATE_METHODS, methodBlock );

    return result;
}

string CppGenerator::doProperty( PropertyNode const& astNode )
{
    //TODO
    /*
    string result = astNode.arguments.size() == 0
        ? doSingleProperty( astNode )
        : doIndexerProperty( astNode );

    return result;
    */
    return "";
}



string CppGenerator::visitLiteralString( string const& value )
{
    //TODO
    /*
    string as3ScapedString = boost::replace_all_copy( value, "\"", "\\\"");
    return "\"" + as3ScapedString + "\"";
    */
    return "";
}


string CppGenerator::doSingleProperty( PropertyNode const& astNode )
{
    //TODO
    /*
    string type = translateType( SymbolTable::joinIdentifier( astNode.type ) );
    string name = astNode.name.name;
    string fieldName = fieldNameFromPropertyName( name );

    string getMethod;
    string setMethod;

    string accessModifier = getAccessModifier( astNode.accessModifier );

    PropertyAccess paAccess = astNode.access.get();
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

    string propertyAttributes = doAttributes( astNode.attributes );

    string result =
        singleFieldType( type, fieldName ) + "\n\n" +
        propertyAttributes +
        getMethod +
        setMethod;

    return result;
    */
    return "";

}

string CppGenerator::doIndexerProperty( PropertyNode const& astNode )
{
    //TODO
    /*
    string type = translateType( SymbolTable::joinIdentifier( astNode.type ) );
    string name = astNode.name.name;

    ArgumentNode const& argument1 = astNode.arguments.front();
    string argument1Name = argument1.name.name;
    string argument1Type = translateType( SymbolTable::joinIdentifier( argument1.type ) );

    string fieldName = fieldNameFromPropertyName( name );

    string getMethod;
    string setMethod;

    //there are no indexer or property with arguments in as3... so we use functions

    string accessModifier = getAccessModifier( astNode.accessModifier );

    PropertyAccess paAccess = astNode.access.get();
    if ( paAccess == PropertyAccess::paRead || paAccess == PropertyAccess::paReadWrite )
    {
        //_formatter.indexerGetProperty = "\n%6% function get %2%( %5%: %4% ):%1% { return this.%3%[ %5% ]; }";
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
        //_formatter.indexerSetProperty = "\n%6% function set %2%( %5%: %4%, %1: value ):void { this.%3%[ %5% ] = value; }";
        setMethod = (boost::format( _formatter.indexerSetProperty )
            % type
            % name
            % fieldName
            % argument1Type
            % argument1Name
            % accessModifier ).str();
    }

    string propertyAttributes = doAttributes( astNode.attributes );
    string result =
        indexerFieldType( type, fieldName, argument1Type ) + "\n\n" +
        propertyAttributes +
        getMethod +
        setMethod;

    return result;
    */
    return "";

}

string CppGenerator::doArgument( ArgumentNode const& astNode )
{
    //TODO
    /*
    return astNode.name.name + ": " + translateType( SymbolTable::joinIdentifier( astNode.type ) );
    */
    return "";
}

string CppGenerator::doConst( ast::ConstNode const& constNode )
{
    //TODO
    return "";
}


string CppGenerator::doMethod( MethodNode const& astNode )
{
    //TODO

    return "";
}


bool CppGenerator::outputClass( string const& classSource, ClassNode const& classAstNode )
{
    //TODO
    /*
    string className = classAstNode.name.name;
    string fullClassName = _mainNamespace + "." + className;

    auto outputPath = fs::path( _outputFolder ) / "src" / boost::replace_all_copy( _mainNamespace, ".", "/" );

    fs::path outputFolder( outputPath );
    if ( !fs::is_directory( outputPath ) )
    {
        fs::create_directories(  outputPath );
    }

    std::string fullFileName = (outputPath / ( className  + _fileExt )).string();

    bool result = Generator::outputClass( classSource, fullClassName );

    return result;
    */
    return false;
}

string CppGenerator::internalGenerateMethod(
    string const& accessModifier,
    string const& as3type,
    string const& name,
    string const& virtuality,
    pam::pdl::ast::ArgumentList const& ArgumentList
    )
{
    //TODO

    //auto formatter = boost::format("%1% %2% function %3%( %4% ): %5% { /*TODO Override*/ return %6%; }")
    /*
        % accessModifier
        % virtuality
        % name
        % doArgumentList( ArgumentList )
        % as3type
        % getEmptyValueOfType( as3type );

    string result = formatter.str();

    return result;
    */
    return "";
}

Generator::AttributeInfo CppGenerator::processAttributeName( FullIdentifierNode const& attrName )
{
    //TODO
    /*
    //custom attribute class is ignored in as3
    Generator::AttributeInfo result;
    result.classInfo = SymbolTable::parseFullClassName( attrName );
    result.name = result.classInfo.second;
    return result;
    */
    return AttributeInfo();
}

fs::path CppGenerator::getFileOutputFolder()
{
    return fs::path( _outputFolder ) / "cpp";
}
