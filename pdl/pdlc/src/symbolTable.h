#pragma once

#include "symbols.h"

namespace io { namespace pdl { namespace symbols
{


class SymbolTable
{
private:

    std::unordered_map<std::string, SymbolPtr> _table;

public:
	SymbolTable() {}
	~SymbolTable() 
    {
    }   

    typedef std::pair<std::string, std::string> ParsedClassname;

    //void add( Symbol& symbol )
    void add( SymbolPtr symbol )
    {
        _table[ symbol->name() ] = symbol;
    }

    bool exists( const std::string& name )
    {
	    const auto result = _table.find( name ) != _table.end();
        return result;
    }

    template<typename T>
    T getSymbol( const std::string& name )
    {
        T result;
        const auto iter = _table.find( name );
        if ( iter != _table.end() )
        {
            result = boost::static_pointer_cast<typename T::element_type>( iter->second );
        }
        return result;
    }

    NamespaceSymbolPtr addNamespace( std::string const& fullName );
    NamespaceSymbolPtr addNamespace( ast::NamespaceNode const& astNode );

    ClassSymbolPtr addClass( ast::ClassNode const& astNode, NamespaceSymbolPtr parent );
    ClassSymbolPtr addUsingClass( ast::UsingNode const& astNode );

    MethodSymbolPtr addMethod( ast::MethodNode const& astNode, ClassSymbolPtr parent );
    ConstSymbolPtr addConst( ast::ConstNode const& astNode, ClassSymbolPtr parent );
    PropertySymbolPtr addProperty( ast::PropertyNode const& astNode, ClassSymbolPtr parent );
    PropertySymbolPtr addProperty( ast::ShortPropertyNode const& astNode, ClassSymbolPtr parent );

    bool classExists( std::string const& className );

    bool namespaceExists( std::string const& fullName );
    NamespaceSymbolPtr getNamespace( std::string const& fullName );

    static std::string joinIdentifier( ast::FullIdentifierList const& identifierList )
    {
        std::vector<std::string> vString;
        for( auto const& identifier : identifierList )
        {
            vString.push_back( identifier.name );
        }

	    auto result = boost::algorithm::join( vString, "." );

        return result;
    }

    static ParsedClassname parseFullClassName( ast::FullIdentifierNode const& fullIdentifier )
    {
        std::vector<std::string> vString;

        for(auto const& identifier : fullIdentifier )
        {
            vString.push_back( identifier.name );
        }

	    auto className = vString.back();

        vString.pop_back();
	    auto namespaceName = boost::algorithm::join( vString, "." );

		ParsedClassname result( namespaceName, className );

        return result;
    }

};

class GlobalSymbolTable: public SymbolTable
{
public:
    void addIntrinsicType( std::string const& type )
    {
        SymbolPtr symbol( new Symbol( type ) );
        symbol->_isIntrinsicType = true;
        add( symbol );
    }
};




}}}

