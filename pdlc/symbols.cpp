#include "stdafx.h"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"

using namespace pam::pdl::ast;
using namespace pam::pdl::symbols;

ContainerSymbol::ContainerSymbol( std::string const& name ):
    Symbol( name ),
    _symbolTable( SymbolTablePtr( new SymbolTable() ) )
{
}

//void ClassMemberSymbol::addArgument( ArgumentNode const& astNode )
void ClassMemberSymbol::addArgument( std::string const& name, std::string const& fullType )
{
    VarSymbolPtr varSymbol( new VarSymbol( name, fullType, ClassMemberSymbolPtr(this) ) );
    _symbolTable->add( varSymbol );
    _arguments[ name ] = varSymbol;
}


