#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "jsGenerator.h"

using namespace pam::pdl;
using namespace symbols;
using namespace config;
using namespace codegen;
using namespace ast;


using std::string;
namespace fs = boost::filesystem;

JsGenerator::JsGenerator( PdlConfig const& config ) :
    Generator( config, "js" )
{
    _fileExt = ".js";
    _classTemplate += _fileExt;
    _generateAsObject = PdlConfig::as_bool( _languageConfig, "generateAsObject" );
    _variableDeclarationKeyword = PdlConfig::as_bool( _languageConfig, "useLet" ) ? "let" : "var";


    //	_formatter.attr =
    //		"{\n" +
    //		_indent + "name: '%1%',\n" +
    //		_indent + "values: {\n" +
    //		_indent + _indent + "%2%" + "\n" +
    //		_indent + "}\n" +
    //		"}";

    _formatter.attr = R"attr(
{
    name: '%1%',
    values: {
    %2%
    }
})attr";

    _formatter.paramlessAttr = "{ name: '%1%' }";
    _formatter.constMember = _generateAsObject
                                 ? "%3%\n%1%: %2%"
                                 : "%3%\nthis.%1% = %2%;";

    //_formatter.singleField = "this.%1% = %2%;\n";
    _jsFormatter.jsDocType = "/** @type {%1%} */";
    _jsFormatter.jsDocVar = "/** @var {%1%} %2% */";
    _jsFormatter.jsDocParam = "/** @param {%1%} %2% */";

    _jsFormatter.jsDocClass = " * @class {%1%}%2%";

    _formatter.singleField = "%3%\nthis.%1% = %2%;\n";
    _formatter.indexerField = "%2%\nthis.%1% = {};\n";


    _typesMap[ "object" ] = "Object";
    _typesMap[ "function" ] = "Function";
    _typesMap[ "string" ] = "String";
    _typesMap[ "bool" ] = "Boolean";
    _typesMap[ "double" ] = "Number";
    _typesMap[ "int" ] = "int";
    _typesMap[ "uint" ] = "int";
    _typesMap[ "array" ] = "[]";

    addScalarType( _typesMap[ "bool" ] );
    addScalarType( _typesMap[ "double" ] );
    addScalarType( _typesMap[ "uint" ] );
    addScalarType( _typesMap[ "string" ] );

    _emptyValueOfType[ "Boolean" ] = "false";
    _emptyValueOfType[ "String" ] = "''";
    _emptyValueOfType[ "int" ] = "0";
    _emptyValueOfType[ "double" ] = "0.0";
    _emptyValueOfType[ "[]" ] = "null";

    _attributeSeparator = ",\n";
}

string JsGenerator::doNamespace( NamespaceNode const& astNamespaceNode )
{
    _mainNamespace = SymbolTable::joinIdentifier( astNamespaceNode.name );

    string leftSide;
    for ( const auto& item : astNamespaceNode.name )
    {
        leftSide += item.name;
        _namespaceLines.push_back( ( boost::format( "%1% = %1% || {};" ) % leftSide ).str() );
        leftSide += '.';
    }

    //string result = generator::doNamespace( astNamespaceNode );

    for ( auto const& classAstNode : astNamespaceNode.classes )
    {
        _methods.clear();
        _properties.clear();
        _outputHeader = createFileHeader( _mainNamespace + "." + classAstNode.name.name );

        const auto classResult = doClass( astNamespaceNode, classAstNode );

        prepareOutputFolder( astNamespaceNode, classAstNode );
        outputClass( classResult, classAstNode );
    }
    return "";
}


string JsGenerator::doUsingList( UsingList const& astNode )
{
    for ( auto const& usingNode : astNode )
    {
        _usingLines.insert( _indent + string( "using " ) + SymbolTable::joinIdentifier( usingNode.className ) + ";" );
    }
    auto result = boost::join( _usingLines, "\n" );
    return result;
}


string JsGenerator::doParentClass( ClassNode const& classAstNode )
{
	string result = "";
    //	if ( classAstNode.parentClass )
    //	{
    //		auto parentClass = classAstNode.symbol->getParentClass();
    //
    //		_fullParentClass = parentClass->getNamespace()->name() + "." + parentClass->name();
    //
    //		result = ( boost::format( "%1% = $.extend( true, {}, %2%, %1% );" )
    //			% _classInNamespace % _fullParentClass ).str();
    //	}

    return result;
}

std::set< std::string > JsGenerator::readNamespaceFile( std::string namespaceFile )
{
    std::set< std::string > result;
    auto outputPath = _outputFolder;

    if ( fs::exists( namespaceFile ) )
    {
        std::ifstream file( namespaceFile );
        std::string line;

        while ( std::getline( file, line ) )
        {
            result.insert( line );
        }
    }

    return result;
}

void JsGenerator::addNamespaceToFile( std::string const& namespaceFile )
{
    const auto fullNamespaceFile = ( fs::path( _outputFolder ) / "js" / namespaceFile ).string();
    auto fileNamespaceLines = readNamespaceFile( fullNamespaceFile );

    std::copy( _namespaceLines.begin(), _namespaceLines.end(),
               std::inserter( fileNamespaceLines, fileNamespaceLines.end() ) );

    std::ofstream file( fullNamespaceFile, std::ios_base::trunc );

    const std::ostream_iterator< std::string > iterator( file, "\n" );
    std::copy( fileNamespaceLines.begin(), fileNamespaceLines.end(), iterator );
}

void JsGenerator::processNamespaces()
{
    _namespaceSimulation = boost::join( _namespaceLines, "\n" );

    const auto namespaceFile = PdlConfig::as_string( _languageConfig, "namespaceFile" );
    if ( !namespaceFile.empty() )
    {
        addNamespaceToFile( namespaceFile );
    }
}

std::string JsGenerator::generateJsDocProperties() const
{
    const auto result = boost::join( _jsDocProperties, "\n" );
    return result;
}

/**
 *
 */
std::string JsGenerator::generateJsDocClass( ClassNode const& classNode ) const
{
    std::string jsDocExtends;
    if ( classNode.parentClass )
    {
        jsDocExtends = ( boost::format( "\n * @extends {%1%}" ) % _fullParentClass ).str();
    }

    auto result = ( boost::format( _jsFormatter.jsDocClass )
        % _classInNamespace
        % jsDocExtends
    ).str();

    return result;
}


string JsGenerator::doClass( NamespaceNode const& astNamespace, ClassNode const& classNode )
{
    _usingLines.clear();

    std::vector< std::string > classAttributes;
    classAttributes.push_back( getAccessModifier( classNode.accessModifier ) );
    const auto classAttrs = boost::join( classAttributes, " " );

    _classInNamespace = _mainNamespace + "." + classNode.name.name;
    //const auto inheritance = doParentClass( classAstNode );

    _fullParentClass = "null";
    if ( classNode.parentClass )
    {
        auto parentClass = classNode.symbol->getParentClass();
        _fullParentClass = parentClass->getNamespace()->name() + "." + parentClass->name();
    }

    doMembers( classNode.members );


    auto constBlock = _generateAsObject
                          ? boost::join( _consts, ",\n" )
                          : boost::join( _consts, "\n" );

    boost::replace_all( constBlock, "\n", "\n" + _indent );
    constBlock = _indent + constBlock;

    auto propertyBlock = boost::join( _properties, "\n" );
    boost::replace_all( propertyBlock, "\n", "\n" + _indent );
    propertyBlock = _indent + propertyBlock;

    auto methodBlock = boost::join( _methods, "\n" );
    boost::replace_all( methodBlock, "\n", "\n" + _indent );
    methodBlock = _indent + methodBlock;

    const auto usingBlock = doUsingList( astNamespace.usings );

    //_classInNamespace = "root." + _mainNamespace + "." + classAstNode.name.name;


    processNamespaces();

    if ( !_propertyAttributes.empty() )
    {
        _propertyAttributes.insert( _propertyAttributes.begin(), _variableDeclarationKeyword + " propertyAttributes = {};" );
        const auto assignAttributes = ( boost::format( "%1%.__propertyAttributes = propertyAttributes;" )
            % _classInNamespace ).str();

        _propertyAttributes.push_back( assignAttributes );
    }
    const auto propertyAttributes = boost::join( _propertyAttributes, "\n" );

    const auto jsDocClass = generateJsDocClass( classNode );

    const auto jsDocProperties = generateJsDocProperties();

    const auto templateCode = readTemplateCode();
    auto result = boost::replace_all_copy( templateCode, GENERIC_TEMPLATE_HEADER, _outputHeader );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_USING_BLOCK, usingBlock );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_NAMESPACE, _mainNamespace );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_CLASS_ATTRS, classAttrs );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_JSDOC_CLASS, jsDocClass );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_JSDOC_PROPERTIES, jsDocProperties );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_CLASS_NAME, classNode.name.name );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_FULL_PARENT_CLASS, _fullParentClass );

    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_CONSTS, constBlock );

    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_PROPERTIES, propertyBlock );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_PROPERTIES_ATTRIBUTES, propertyAttributes );
    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_METHODS, methodBlock );

    result = boost::replace_all_copy( result, JSGEN_TEMPLATE_CLASS_DECLARATION, _classInNamespace );

    return result;
}

string JsGenerator::doProperty( PropertyNode const& propertyNode )
{
    auto result = propertyNode.arguments.size() == 0
                      ? doSingleProperty( propertyNode )
                      : doIndexerProperty( propertyNode );

    return result;
}

std::string JsGenerator::doShortProperty( ShortPropertyNode const& shortPropertyNode )
{
    auto result = doSingleProperty( shortPropertyNode );
    return result;
}


string JsGenerator::visitLiteralString( string const& value )
{
    const auto jsScapedString = boost::replace_all_copy( value, "\"", "\\\"" );
    return "\"" + jsScapedString + "\"";
}


void JsGenerator::generateJsDocProperty( const std::string& type, const std::string& name )
{
    const auto jsDocProperty = ( boost::format( " * @property {%1%} %2%" )
        % type
        % name
    ).str();

    _jsDocProperties.push_back( jsDocProperty );
}

string JsGenerator::doSingleProperty( PropertyNode const& propertyNode )
{
    //const auto type = getPropertyType( propertyNode.propertyType, propertyNode.symbol->getTypeSymbol() );

    const auto name = propertyNode.name.name;

    if ( propertyNode.attributes.size() > 0 )
    {
        const auto attributes = boost::
            replace_all_copy( doAttributes( propertyNode.attributes ), "\n", "\n" + _indent );
        _propertyAttributes.push_back(
            ( boost::format( "propertyAttributes.%1% = [\n" + _indent + attributes ) % name ).str() + "\n];" );
    }

    const auto type = getJsPropertyType( propertyNode );

    const auto jsDocType = ( boost::format( _jsFormatter.jsDocType )
        % type
    ).str();

    generateJsDocProperty( type, name );

    auto result = ( boost::format( _formatter.singleField )
        % name
        % getEmptyValueOfType( type )
        % jsDocType
        //% propertyAttributes*/ 
    ).str();

    return result;
}

string JsGenerator::doSingleProperty( ShortPropertyNode const& shortPropertyNode )
{
    const auto type = getJsPropertyType( shortPropertyNode );

    const auto name = shortPropertyNode.name.name;

    const auto jsDocType = ( boost::format( _jsFormatter.jsDocType )
        % type
    ).str();

    generateJsDocProperty( type, name );

    auto result = ( boost::format( _formatter.singleField )
        % name
        % getEmptyValueOfType( type )
        % jsDocType
    ).str();

    return result;
}

string JsGenerator::doIndexerProperty( PropertyNode const& propertyNode )
{
    const auto type = getJsPropertyType( propertyNode );
    //const auto type = getPropertyType( propertyNode.propertyType, propertyNode.symbol->getTypeSymbol() );
    const auto name = propertyNode.name.name;

    auto const& argument1 = propertyNode.arguments.front();
    auto argument1Name = argument1.name.name;
    auto argument1Type = translateType( SymbolTable::joinIdentifier( argument1.type ) );

    if ( !propertyNode.attributes.empty() )
    {
        auto attributes = ( boost::format( "propertyAttributes[ '%1%' ] = \n{\n%2%\n}" )
            % name
            % doAttributes( propertyNode.attributes )
        ).str();
        _propertyAttributes.push_back( attributes );
    }

    const auto jsDocType = ( boost::format( _jsFormatter.jsDocType )
        % type
    ).str();

    generateJsDocProperty( type, name );

    //_formatter.indexerField = "%2%\nthis.%1% = {};\n";
    auto result = ( boost::format( _formatter.indexerField )
        % name
        % jsDocType
        //% argument1Type
        //% propertyAttributes 
        //% getEmptyValueOfType( type ) 
    ).str();

    return result;
}


string JsGenerator::doArgument( ArgumentNode const& astNode )
{
    return "/*" + translateType( SymbolTable::joinIdentifier( astNode.type ) ) + "*/" +
        astNode.name.name;
}


bool JsGenerator::outputClass( string const& classSource, ClassNode const& classAstNode )
{
    const auto fullClassName = _mainNamespace + "." + classAstNode.name.name;
    const auto result = Generator::outputClass( classSource, fullClassName );
    return result;
}

fs::path JsGenerator::getFileOutputFolder()
{
    return fs::path( _outputFolder ) / "js" / boost::replace_all_copy( _mainNamespace, ".", "/" );
    //return fs::path( _outputFolder ) / "js";
}

string JsGenerator::doMethod( MethodNode const& astNode )
{
    const string format = "this.%2% = function( %3% )/*ret type:%1%*/ { return %4%; };";

    const auto returnType = translateType( SymbolTable::joinIdentifier( astNode.type ) );

    auto formatter = boost::format( format )
        % returnType
        % astNode.name.name
        % doArgumentList( astNode.arguments )
        % getEmptyValueOfType( returnType );

    auto result = formatter.str();

    return result;
}

string JsGenerator::doConst( ConstNode const& constNode )
{
    const auto value = boost::apply_visitor( _literalVisitor, constNode.value );
    const auto type = translateType( SymbolTable::joinIdentifier( constNode.type ) );

    const auto jsDocType = ( boost::format( _jsFormatter.jsDocType )
        % type
    ).str();

    auto formatter = boost::format( _formatter.constMember )
        % constNode.name
        % value
        % jsDocType;

    auto result = formatter.str();

    return result;
}

Generator::AttributeInfo JsGenerator::processAttributeName( FullIdentifierNode const& attrName )
{
    //same as as3
    AttributeInfo result;
    result.classInfo = SymbolTable::parseFullClassName( attrName );
    result.name = result.classInfo.second;
    return result;
}


string JsGenerator::visitAttrRequiredParams( AttrRequiredParams const& astNode )
{
    std::vector< string > paramVector;
    auto index = 1;
    for ( auto const& param : astNode )
    {
        paramVector.push_back(
            "default" + std::to_string( index ) + ": " + boost::apply_visitor( _literalVisitor, param ) );
        index++;
    }
    auto result = boost::join( paramVector, ",\n" + _indent + _indent );
    return result;
}

string JsGenerator::visitAttrRequiredAndOptionals( AttrRequiredAndOptionals const& astNode )
{
    auto result = visitAttrRequiredParams( astNode.required ) + ",\n" + _indent + _indent + visitAttrOptionalParams(
        astNode.optionals );
    return result;
}

string JsGenerator::getTypeFromAttribute( const ast::AttributeListNode& attributes )
{
    string result;
    for ( auto const& attribute : attributes )
    {
        const auto attributeInfo = processAttributeName( attribute.name );
        if ( attributeInfo.name == JSGEN_ATTRI_JSTYPE )
        {
            const auto params = boost::get< AttrRequiredParams >( attribute.params );
            const auto firstParam = params.front();
            result = boost::get< string >( firstParam );
            break;
        }
    }
    return result;
}

string JsGenerator::visitAttrOptionalParams( AttrOptionalParams const& astNode )
{
    std::vector< string > paramVector;
    for ( auto const& param : astNode )
    {
        paramVector.push_back( param.name.name + ": " + boost::apply_visitor( _literalVisitor, param.value ) );
    }
    auto result = boost::join( paramVector, ",\n" + _indent + _indent );
    return result;
}


string JsGenerator::getJsPropertyType( ast::PropertyNode const& propertyNode )
{
    auto result = getTypeFromAttribute( propertyNode.attributes );

    if ( result.empty() )
    {
        result = internalGetJsPropertyType( propertyNode.symbol, propertyNode.propertyType );
    }

    return result;
}

std::string JsGenerator::internalGetJsPropertyType( PropertySymbolPtr const& propertySymbol,
                                                    PropertyType const& propertyType )
{
    std::string result;
    auto typeSymbol = propertySymbol->getTypeSymbol();
    if ( typeSymbol->isClass() )
    {
        auto& classSymbol = *static_cast< ClassSymbol* >(typeSymbol.get());
        result = classSymbol.getNamespace()->name() + '.' + classSymbol.name();
    }
    else
    {
        result = translateType( SymbolTable::joinIdentifier( propertyType.type ) );
    }

    result += arrayNotationList2Brackets( propertyType.arrayNotationList );

    return result;
}

std::string JsGenerator::getJsPropertyType( ShortPropertyNode const& shortPropertyNode )
{
    auto result = internalGetJsPropertyType( shortPropertyNode.symbol, shortPropertyNode.propertyType );
    return result;
}
