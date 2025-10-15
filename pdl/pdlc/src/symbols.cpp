#include "commonHeaders.hpp"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"

using namespace io::pdl::ast;
using namespace io::pdl::symbols;

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


