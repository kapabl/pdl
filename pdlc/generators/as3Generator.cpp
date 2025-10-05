#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "as3Generator.h"

using namespace pam::pdl;
using namespace pam::pdl::symbols;
using namespace pam::pdl::codegen;
using namespace pam::pdl::ast;

using std::string;
namespace fs = boost::filesystem;

As3Generator::As3Generator( pam::pdl::config::PdlConfig const& config ) :
    Generator( config, "as3" )
{
    _fileExt = ".as";

    _classTemplate += _fileExt;

    _formatter.constMember = "public static const %1%:%2% = %3%;";

    _formatter.paramlessAttr = "[%1%]";
    _formatter.singleField = "private var %2%: %1%;";
    _formatter.indexerField = "private var %2%:Object = {};/* %3% -> %1% */";

    _formatter.singleGetProperty = "\n%4% function get %2%():%1% { return this.%3%; }";
    _formatter.singleSetProperty = "\n%4% function set %2%( value: %1% ):void { this.%3% = value; }";

    //there are no indexer or property with arguments in as3... so we use functions
    _formatter.indexerGetProperty = "\n%6% function get%2%( %5%: %4% ):%1% { return this.%3%[ %5% ]; }";
    _formatter.indexerSetProperty = "\n%6% function set%2%( %5%: %4%, value: %1% ):void { this.%3%[ %5% ] = value; }";

    _controllableGetTemplate = "\n%3% function get %2%():%1% { return getPropertyValue( '%2%', '%1%'); }";
    _controllableSetTemplate = "\n%3% function set %2%(value:%1%):void { setPropertyValue( '%2%', value ); }";


    _typesMap[ "object" ] = "Object";
    _typesMap[ "function" ] = "Function";
    _typesMap[ "string" ] = "String";
    _typesMap[ "bool" ] = "Boolean";
    _typesMap[ "double" ] = "Number";
    _typesMap[ "array" ] = "Array";

    addScalarType( _typesMap[ "double" ] );
    addScalarType( _typesMap[ "bool" ] );
    addScalarType( _typesMap[ "string" ] );


    _emptyValueOfType[ "Boolean" ] = "false";
    _emptyValueOfType[ "String" ] = "\"\"";
    _emptyValueOfType[ "int" ] = "0";
    _emptyValueOfType[ "Number" ] = "0.0";
    _emptyValueOfType[ "Array" ] = "[]";
}

bool As3Generator::isInMainNamespace( std::string const& namespaceString ) const
{
    const auto result = _pdlMainNamespace == namespaceString;
    return result;
}

string As3Generator::doUsingList( UsingList const& usingList )
{
    for ( auto const& usingNode : usingList )
    {
        if ( usingNode.symbol->getUseCount() > 0 )
        {
            const auto namespaceString = usingNode.symbol->getNamespace()->name();
            if ( !usingNode.symbol->isAttributeClass() && !isInMainNamespace( namespaceString ) )
            {
                _importLines.insert(
                    _indent + string( "import " ) + SymbolTable::joinIdentifier( usingNode.className ) + ";" );
            }
            else
            {
                _importLines.insert(
                    _indent + string( "//import " ) + SymbolTable::joinIdentifier( usingNode.className ) + ";" );
            }
        }
    }
    auto result = boost::join( _importLines, "\n" );

    return result;
}

string As3Generator::doParentClass( ClassNode const& classAstNode )
{
    string result;
    if ( classAstNode.parentClass )
    {
        const auto pair = SymbolTable::parseFullClassName( classAstNode.parentClass.get() );
        result = string( "extends " ) + pair.second;
    }
    return result;
}

string As3Generator::doClass( NamespaceNode const& astNamespace, ClassNode const& classAstNode )
{
    _importLines.clear();

    _hasPropertyControl = classAstNode.symbol->hasPropertyControl();

    std::vector< std::string > classAttributes;
    classAttributes.push_back( getAccessModifier( classAstNode.accessModifier ) );
    const auto classAttrs = boost::join( classAttributes, " " );

    const auto inheritance = doParentClass( classAstNode );

    doMembers( classAstNode.members );

    auto constBlock = boost::join( _consts, "\n" );
    boost::replace_all( constBlock, "\n", "\n" + _indent + _indent );

    auto propertyBlock = boost::join( _properties, "\n\n" );
    boost::replace_all( propertyBlock, "\n", "\n" + _indent + _indent );

    auto methodBlock = boost::join( _methods, "\n\n" );
    boost::replace_all( methodBlock, "\n", "\n" + _indent + _indent );

    const auto importBlock = doUsingList( astNamespace.usings );

    const auto templateCode = readTemplateCode();


    const auto parentConstructor = classAstNode.parentClass
                                       ? "super();"
                                       : "";


    auto result = boost::replace_all_copy( templateCode, GENERIC_TEMPLATE_HEADER, _outputHeader );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_IMPORT_BLOCK, importBlock );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_PACKAGE, _mainNamespace );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_CLASS_ATTRS, classAttrs );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_CLASS_NAME, classAstNode.name.name );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_CLASS_INHERITANCE, inheritance );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_PROPERTIES, propertyBlock );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_CONSTS, constBlock );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_METHODS, methodBlock );
    result = boost::replace_all_copy( result, AS3GEN_TEMPLATE_PARENT_CONSTRUCTOR, parentConstructor );

    return result;
}

string As3Generator::doProperty( PropertyNode const& astNode )
{
    auto result = astNode.arguments.size() == 0
                      ? doSingleProperty( astNode )
                      : doIndexerProperty( astNode );

    return result;
}

string As3Generator::doShortProperty( ShortPropertyNode const& shortPropertyNode )
{
    auto result = doSingleProperty( shortPropertyNode );

    return result;
}


string As3Generator::visitLiteralString( string const& value )
{
    const auto as3EscapedString = boost::replace_all_copy( value, "\"", "\\\"" );
    return "\"" + as3EscapedString + "\"";
}



string As3Generator::doSingleProperty( ShortPropertyNode const& shortPropertyNode )
{
    const auto type = getPropertyType( shortPropertyNode.propertyType );

    string result;

    const std::string accessModifier = "public";

    if ( _hasPropertyControl && isScalarType( ( type ) ) )
    {
        result = doSinglePropertyWithCall( shortPropertyNode, accessModifier, PropertyAccess::paReadWrite, "" );
    }
    else
    {
        result = doSinglePropertyWithField( shortPropertyNode, accessModifier, PropertyAccess::paReadWrite, "" );
    }

    return result;

}


string As3Generator::doSingleProperty( PropertyNode const& propertyNode )
{
    const auto type = getPropertyType( propertyNode.propertyType );
    string result;

    const auto accessModifier = getAccessModifier( propertyNode.accessModifier );
    const auto propertyAccess = propertyNode.access.get();

    const auto attributes = doAttributes( propertyNode.attributes );

    if ( _hasPropertyControl && isScalarType( ( type ) ) )
    {
        result = doSinglePropertyWithCall( propertyNode, accessModifier, propertyAccess, attributes );
    }
    else
    {
        result = doSinglePropertyWithField( propertyNode, accessModifier, propertyAccess, attributes );
    }

    return result;
}


std::string As3Generator::getPropertyType( PropertyType const& propertyType )
{
    auto result = translateType( SymbolTable::joinIdentifier( propertyType.type ) );

    string as3Vectors;
    const auto dimensions = propertyType.arrayNotationList.size();
    for ( unsigned int i = 0; i < dimensions; i++ )
    {
        as3Vectors += "Vector.<";
    }
    result = as3Vectors + result;

    for (unsigned int i = 0; i < dimensions; i++ )
    {
        result += ">";
    }


    return result;
}

string As3Generator::doIndexerProperty( PropertyNode const& propertyNode )
{
    const auto type = getPropertyType( propertyNode.propertyType );
    const auto name = toCamelString( propertyNode.name.name );

    auto const& argument1 = propertyNode.arguments.front();
    const auto argument1Name = argument1.name.name;
    const auto argument1Type = translateType( SymbolTable::joinIdentifier( argument1.type ) );

    const auto fieldName = fieldNameFromPropertyName( name );

    string getMethod;
    string setMethod;

    //there are no indexer or property with arguments in as3... so we use functions
    /*
    _formatter.indexerGetProperty = "\n%6% function get%2%( %5%: %4% ):%1% { return this.%3%[ %5% ]; }";
    _formatter.indexerSetProperty = "\n%6% function set%2%( %5%: %4%, %1%: value ):void { this.%3%[ %5% ] = value; }";
    */

    const auto accessModifier = getAccessModifier( propertyNode.accessModifier );

    const auto paAccess = propertyNode.access.get();
    if ( paAccess == PropertyAccess::paRead || paAccess == PropertyAccess::paReadWrite )
    {
        //_formatter.indexerGetProperty = "\n%6% function get %2%( %5%: %4% ):%1% { return this.%3%[ %5% ]; }";
        getMethod = ( boost::format( _formatter.indexerGetProperty )
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
        setMethod = ( boost::format( _formatter.indexerSetProperty )
            % type
            % name
            % fieldName
            % argument1Type
            % argument1Name
            % accessModifier ).str();
    }

    const auto propertyAttributes = doAttributes( propertyNode.attributes );
    auto result =
        indexerFieldType( type, fieldName, argument1Type ) + "\n\n" +
        propertyAttributes +
        getMethod +
        setMethod;

    return result;
}

string As3Generator::doArgument( ArgumentNode const& astNode )
{
    return astNode.name.name + ": " + translateType( SymbolTable::joinIdentifier( astNode.type ) );
}


string As3Generator::doMethod( MethodNode const& astNode )
{
    //In the case of AS3 most method should be "virtual" so 
    //you can override them in your project

    auto result = internalGenerateMethod(
        getAccessModifier( astNode.accessModifier ),
        translateType( SymbolTable::joinIdentifier( astNode.type ) ),
        astNode.name.name,
        "",
        astNode.arguments );

    return result;
}

string As3Generator::doConst( ConstNode const& constNode )
{
    const auto as3Type = translateType( SymbolTable::joinIdentifier( constNode.type ) );
    const auto value = boost::apply_visitor( _literalVisitor, constNode.value );

    auto formatter = boost::format( _formatter.constMember )
        % constNode.name
        % as3Type
        % value;

    auto result = formatter.str();

    return result;
}


bool As3Generator::outputClass( string const& classSource, ClassNode const& classAstNode )
{
    const auto className = classAstNode.name.name;
    const auto fullClassName = _mainNamespace + "." + className;

    const auto result = Generator::outputClass( classSource, fullClassName );

    return result;
}

fs::path As3Generator::getFileOutputFolder()
{
    return fs::path( _outputFolder ) / "as3/src" / boost::replace_all_copy( _mainNamespace, ".", "/" );
}



string As3Generator::internalGenerateMethod(
    string const& accessModifier,
    string const& as3type,
    string const& name,
    string const& virtuality,
    pam::pdl::ast::ArgumentList const& ArgumentList
)
{
    auto formatter = boost::format( "%1% %2% function %3%( %4% ): %5% { /*TODO Override*/ return %6%; }" )
        % accessModifier
        % virtuality
        % name
        % doArgumentList( ArgumentList )
        % as3type
        % getEmptyValueOfType( as3type );

    string result = formatter.str();

    return result;
}

Generator::AttributeInfo As3Generator::processAttributeName( FullIdentifierNode const& attrName )
{
    //custom attribute class is ignored in as3
    Generator::AttributeInfo result;
    result.classInfo = SymbolTable::parseFullClassName( attrName );
    result.name = result.classInfo.second;
    return result;
}
