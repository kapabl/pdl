#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"
#include "goGenerator.h"

using namespace pam::pdl;
using namespace pam::pdl::symbols;
using namespace pam::pdl::codegen;
using namespace pam::pdl::ast;
using namespace pam::pdl::config;

using std::string;
namespace fs = boost::filesystem;

GoGenerator::GoGenerator( PdlConfig const& config ) :
    Generator( config, "go" )
{
    _fileExt = ".go";
    _classTemplate += _fileExt;

    _nullObjectValue = "nil";

    _formatter.paramlessAttr.clear();
    _formatter.attr.clear();
    _formatter.singleField.clear();
    _formatter.indexerField.clear();
    _formatter.singleGetProperty.clear();
    _formatter.singleSetProperty.clear();
    _formatter.indexerGetProperty.clear();
    _formatter.indexerSetProperty.clear();
    _formatter.constMember.clear();

    _attributeSeparator = ", ";

    _typesMap[ ITYPE_BOOL ] = "bool";
    _typesMap[ ITYPE_STRING ] = "string";
    _typesMap[ ITYPE_INT ] = "int";
    _typesMap[ ITYPE_UINT ] = "uint";
    _typesMap[ ITYPE_DOUBLE ] = "float64";
    _typesMap[ ITYPE_ARRAY ] = "[]interface{}";
    _typesMap[ ITYPE_FUNCTION ] = "func()";
    _typesMap[ ROOT_OBJECT ] = "interface{}";

    _emptyValueOfType["bool"] = "false";
    _emptyValueOfType["string"] = "\"\"";
    _emptyValueOfType["int"] = "0";
    _emptyValueOfType["uint"] = "0";
    _emptyValueOfType["float64"] = "0";
    _emptyValueOfType["[]interface{}"] = "nil";
    _emptyValueOfType["interface{}"] = "nil";
    _emptyValueOfType["func()"] = "nil";
}

std::string GoGenerator::doUsingList( UsingList const& usingList )
{
    _imports.clear();

    for ( auto const& usingNode : usingList )
    {
        const auto fullIdentifier = SymbolTable::joinIdentifier( usingNode.className );
        if ( !fullIdentifier.empty() )
        {
            auto importPath = fullIdentifier;
            boost::replace_all( importPath, ".", "/" );
            _imports.insert( importPath );
        }
    }

    return buildImportsBlock( _imports );
}

std::string GoGenerator::doClass( NamespaceNode const& astNamespace, ClassNode const& classNode )
{
    const auto importsBlock = doUsingList( astNamespace.usings );

    _methods.clear();
    _properties.clear();
    _consts.clear();

    _currentStructName = goExportName( classNode.name.name );

    const auto embedded = doParentClass( classNode );

    doMembers( classNode.members );

    std::ostringstream source;
    source << "package " << goPackageName() << "\n\n";

    if ( !importsBlock.empty() )
    {
        source << importsBlock << "\n";
    }

    const auto constBlock = buildConstBlock();
    if ( !constBlock.empty() )
    {
        source << constBlock << "\n";
    }

    source << buildStructBlock( _currentStructName, embedded, _properties ) << "\n";

    const auto methodsBlock = buildMethodsBlock();
    if ( !methodsBlock.empty() )
    {
        source << methodsBlock;
    }

    return source.str();
}

std::string GoGenerator::doParentClass( ClassNode const& classAstNode )
{
    std::string result;
    if ( classAstNode.parentClass )
    {
        const auto parsed = SymbolTable::parseFullClassName( classAstNode.parentClass.get() );
        result = goExportName( parsed.second );
    }
    return result;
}

std::string GoGenerator::doArgument( ArgumentNode const& astNode )
{
    auto argumentName = Generator::toCamelString( astNode.name.name );
    auto argumentType = translateFullIdentifier( astNode.type );
    return argumentName + " " + argumentType;
}

std::string GoGenerator::doConst( ConstNode const& constNode )
{
    const auto name = goExportName( constNode.name.name );
    const auto type = translateFullIdentifier( constNode.type );
    const auto value = boost::apply_visitor( _literalVisitor, constNode.value );
    return name + " " + type + " = " + value;
}

std::string GoGenerator::doMethod( MethodNode const& astNode )
{
    std::vector<std::string> arguments;
    arguments.reserve( astNode.arguments.size() );
    for ( auto const& argument : astNode.arguments )
    {
        arguments.push_back( doArgument( argument ) );
    }

    const auto receiverName = Generator::toCamelString( _currentStructName.empty() ? std::string("receiver") : _currentStructName );
    const auto exportedName = goExportName( astNode.name.name );
    const auto returnType = translateFullIdentifier( astNode.type );

    std::ostringstream method;
    method << "func (" << receiverName << " *" << _currentStructName << ") " << exportedName << "(";
    method << boost::algorithm::join( arguments, ", " ) << ")";

    const bool hasReturn = !returnType.empty() && returnType != ITYPE_VOID;

    if ( hasReturn )
    {
        method << " " << returnType;
    }

    method << " {\n";

    if ( hasReturn )
    {
        method << _indent << "return " << defaultReturnValue( returnType ) << "\n";
    }

    method << "}\n";

    return method.str();
}

std::string GoGenerator::doProperty( PropertyNode const& astNode )
{
    if ( astNode.arguments.empty() )
    {
        return buildPropertyField( astNode.name.name, astNode.propertyType );
    }

    return string("// TODO: indexer property for ") + goExportName( astNode.name.name );
}

std::string GoGenerator::doShortProperty( ShortPropertyNode const& astNode )
{
    return buildPropertyField( astNode.name.name, astNode.propertyType );
}

Generator::AttributeInfo GoGenerator::processAttributeName( FullIdentifierNode const& attrName )
{
    AttributeInfo info;
    info.classInfo = SymbolTable::parseFullClassName( attrName );
    info.name = SymbolTable::joinIdentifier( attrName );
    return info;
}

std::string GoGenerator::visitLiteralString( std::string const& value )
{
    auto escaped = value;
    boost::replace_all( escaped, "\"", "\\\"" );
    return string("\"") + escaped + "\"";
}

bool GoGenerator::outputClass( std::string const& classSource, ClassNode const& astClass )
{
    return Generator::outputClass( classSource, astClass.name.name );
}

boost::filesystem::path GoGenerator::getFileOutputFolder()
{
    return fs::path( _outputFolder ) / "go";
}

std::string GoGenerator::buildImportsBlock( std::set<std::string> const& imports ) const
{
    if ( imports.empty() )
    {
        return {};
    }

    std::ostringstream out;
    out << "import (\n";
    const std::string indent = "\t";
    for ( auto const& importPath : imports )
    {
        out << indent << '"' << importPath << '"' << '\n';
    }
    out << ")";
    return out.str();
}

std::string GoGenerator::buildConstBlock() const
{
    if ( _consts.empty() )
    {
        return {};
    }

    std::ostringstream out;
    if ( _consts.size() == 1 )
    {
        out << "const " << _consts.front();
    }
    else
    {
        out << "const (\n";
        for ( auto const& entry : _consts )
        {
            out << _indent << entry << "\n";
        }
        out << ")";
    }
    return out.str();
}

std::string GoGenerator::buildStructBlock( std::string const& structName, std::string const& embedded, std::vector<std::string> const& fields ) const
{
    std::ostringstream out;
    out << "type " << structName << " struct {\n";
    if ( !embedded.empty() )
    {
        out << _indent << embedded << "\n";
    }
    for ( auto const& field : fields )
    {
        out << _indent << field << "\n";
    }
    out << "}";
    return out.str();
}

std::string GoGenerator::buildMethodsBlock() const
{
    if ( _methods.empty() )
    {
        return {};
    }

    std::ostringstream out;
    for ( size_t i = 0; i < _methods.size(); ++i )
    {
        out << _methods[i];
        if ( i + 1 < _methods.size() )
        {
            out << "\n";
        }
    }
    return out.str();
}

std::string GoGenerator::goExportName( std::string const& identifier ) const
{
    if ( identifier.empty() )
    {
        return identifier;
    }

    auto result = identifier;
    result[ 0 ] = static_cast<char>( toupper( result[ 0 ] ) );
    return result;
}

std::string GoGenerator::goFieldTag( std::string const& name ) const
{
    if ( name.empty() )
    {
        return {};
    }

    return string(" `json:\"") + name + "\"`";
}

std::string GoGenerator::goPackageName() const
{
    if ( _mainNamespace.empty() )
    {
        return "main";
    }

    auto package = _mainNamespace;
    auto const dotPos = package.find_last_of( '.' );
    if ( dotPos != std::string::npos )
    {
        package = package.substr( dotPos + 1 );
    }

    std::transform( package.begin(), package.end(), package.begin(), []( unsigned char c ) { return static_cast<char>( std::tolower( c ) ); } );
    return package;
}

std::string GoGenerator::translateFullIdentifier( FullIdentifierNode const& identifier )
{
    auto const fullIdentifier = SymbolTable::joinIdentifier( identifier );
    auto translated = translateType( fullIdentifier );
    if ( translated == ITYPE_VOID )
    {
        translated.clear();
    }
    return translated;
}

std::string GoGenerator::buildPropertyField( std::string const& name, PropertyType const& type )
{
    const auto fieldName = goExportName( name );
    const auto typeName = getPropertyType( type );
    const auto tag = goFieldTag( name );
    return fieldName + " " + typeName + tag;
}

std::string GoGenerator::defaultReturnValue( std::string const& type ) const
{
    auto defaultValue = _nullObjectValue;
    const auto it = _emptyValueOfType.find( type );
    if ( it != _emptyValueOfType.end() )
    {
        defaultValue = it->second;
    }
    return defaultValue;
}
