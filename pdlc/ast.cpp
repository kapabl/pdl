#include "stdafx.h"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"

using namespace pam::pdl::ast;

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
    std::string fullIdentifier = pam::pdl::symbols::SymbolTable::joinIdentifier( right );
    return left + fullIdentifier;
}

std::string operator+(FullIdentifierList const& left, const char* right)
{
    std::string fullIdentifier = pam::pdl::symbols::SymbolTable::joinIdentifier( left );
    return fullIdentifier + right;
}
