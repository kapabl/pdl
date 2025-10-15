#include "commonHeaders.hpp"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"

using namespace io::pdl::ast;

std::string operator+(const char* left, Identifier const& right )
{
    return left + right.name;
}

std::string operator+(Identifier const& left, const char* right )
{
    return left.name + right;
}

std::string operator+(const char* left, FullIdentifierList const& right)
{
    std::string fullIdentifier = io::pdl::symbols::SymbolTable::joinIdentifier( right );
    return left + fullIdentifier;
}

std::string operator+(FullIdentifierList const& left, const char* right)
{
    std::string fullIdentifier = io::pdl::symbols::SymbolTable::joinIdentifier( left );
    return fullIdentifier + right;
}
