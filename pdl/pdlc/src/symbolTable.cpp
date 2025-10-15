#include "commonHeaders.hpp"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"

using namespace io::pdl;
using namespace io::pdl::symbols;



NamespaceSymbolPtr SymbolTable::addNamespace( std::string const& fullName )
{
    NamespaceSymbolPtr result( new NamespaceSymbol( /*SymbolTablePtr(this), */fullName ) );
    add( result );
    return result;
}

NamespaceSymbolPtr SymbolTable::addNamespace( ast::NamespaceNode const& astNode )
{
	auto result( addNamespace( joinIdentifier( astNode.name ) ) );
    return result;
}

ClassSymbolPtr SymbolTable::addClass( ast::ClassNode const& astNode, NamespaceSymbolPtr parent )
{
    //std::string fullClassName = parent->name() + "." + astNode.name.name;

	auto const& className = astNode.name.name;
    assert( !classExists( className ) );

    ClassSymbolPtr result( new ClassSymbol( className, parent ) );
    add( result );
    //_table[ className ] = result;

    //TODO incorporate more info from the astNode

    return result;
}

ClassSymbolPtr SymbolTable::addUsingClass( ast::UsingNode const& astNode )
{
    auto parsedClassName = SymbolTable::parseFullClassName( astNode.className );

	auto& fullNamespace = parsedClassName.first;
	const auto parentNamespace( getNamespace( fullNamespace ) );

	auto& className = parsedClassName.second;
    assert( !classExists( className ) );

    ClassSymbolPtr result( new ClassSymbol( className, parentNamespace ) );
    add( result );

    result->SetAsExternalClass();

    //TODO incorporate more info from the astNode

    return result;
}

MethodSymbolPtr SymbolTable::addMethod( ast::MethodNode const& astNode, ClassSymbolPtr parent )
{
    MethodSymbolPtr result( new MethodSymbol( astNode.name.name, parent ) );
    add( result );
    //TODO incorporate more info from the astNode
    return result;
}

ConstSymbolPtr SymbolTable::addConst( ast::ConstNode const& constNode, ClassSymbolPtr parent )
{
    ConstSymbolPtr result( new ConstSymbol( constNode.name.name, parent ) );
    add( result );
    //TODO incorporate more info from the astNode
    return result;
}

PropertySymbolPtr SymbolTable::addProperty( ast::PropertyNode const& propertyNode, ClassSymbolPtr parent )
{
    PropertySymbolPtr result( new PropertySymbol(propertyNode.name.name, parent ) );
    add( result );
    return result;
}

PropertySymbolPtr SymbolTable::addProperty(ast::ShortPropertyNode const& shortPropertyNode, ClassSymbolPtr parent)
{
	PropertySymbolPtr result(new PropertySymbol(shortPropertyNode.name.name, parent));
	add(result);
	return result;
}


bool SymbolTable::classExists( std::string const& className )
{
	const auto result = nullptr != getSymbol<ClassSymbolPtr>( className ).get();

    return result;
}

bool SymbolTable::namespaceExists( std::string const& fullName )
{
	const auto result = nullptr != getSymbol<NamespaceSymbolPtr>( fullName ).get();

    return result;
}

NamespaceSymbolPtr SymbolTable::getNamespace( std::string const& fullName )
{
	auto result = getSymbol<NamespaceSymbolPtr>( fullName );

    if ( !result )
    {
        result = addNamespace( fullName );
    }

    return result;
}

