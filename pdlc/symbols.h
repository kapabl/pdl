#pragma once

namespace pam { namespace pdl { namespace symbols
{

#define DECLARE_SYMBOL_CLASS( className ) \
    class className; \
    typedef boost::shared_ptr<className> className ## Ptr;

    
DECLARE_SYMBOL_CLASS( SymbolTable )
DECLARE_SYMBOL_CLASS( GlobalSymbolTable )
DECLARE_SYMBOL_CLASS( Symbol )
DECLARE_SYMBOL_CLASS( NamespaceSymbol )
DECLARE_SYMBOL_CLASS( ClassSymbol )
DECLARE_SYMBOL_CLASS( ClassMemberSymbol )
DECLARE_SYMBOL_CLASS( MethodSymbol )
DECLARE_SYMBOL_CLASS( ConstSymbol )
DECLARE_SYMBOL_CLASS( PropertySymbol )
DECLARE_SYMBOL_CLASS( ShortPropertySymbol )
DECLARE_SYMBOL_CLASS( VarSymbol )

#undef DECLARE_SYMBOL_CLASS

class Symbol
{

    friend class GlobalSymbolTable;
    

protected:
    std::string _name;
    bool _isClass;
    bool _isIntrinsicType;
    int _useCount;

public:
    Symbol( std::string const& name ):
        _name( name ),
        _isClass( false ),
        _isIntrinsicType( false ),
        _useCount( 0 )
    {
    }

    ~Symbol()
    {
    }

    std::string& name() { return _name; }

    bool isClass() const
    { return _isClass; }

    bool isIntrinsicType() const
    { return _isIntrinsicType; }

	/*
    bool setAsIntrinsicType() 
    { 
        _isClass = false;
        _isIntrinsicType = true; 
    }

    bool setAsClass() 
    { 
        _isClass = true;
        _isIntrinsicType = false; 
    }*/

    bool isType() { return _isIntrinsicType || _isClass; }

    int getUseCount() { return _useCount; }
    void incUseCount() { _useCount++; }
    void decUseCount() { _useCount--; }

};

class ContainerSymbol: public Symbol
{
protected:
    SymbolTablePtr _symbolTable;

public:
    ContainerSymbol( std::string const& name );
    ~ContainerSymbol()
    {
    }

    SymbolTablePtr symbolTable() { return _symbolTable; }

};



class NamespaceSymbol: public ContainerSymbol
{
public:
    NamespaceSymbol( std::string const& name ):
        ContainerSymbol( name )
    {
    }

    ~NamespaceSymbol()
    {
    }
};

class ClassSymbol: public ContainerSymbol
{
protected:
    NamespaceSymbolPtr _namespace;
    ClassSymbolPtr _parentClass;
    bool _isExternalClass;
    bool _isAttributeClass;
    bool _hasIndexer;
    bool _hasPropertyControl;

public:
    ClassSymbol( std::string const& name, NamespaceSymbolPtr namespaceSymbol ):
        ContainerSymbol( /*namespaceSymbol->symbolTable(), */name ),
        _namespace( namespaceSymbol ),
        _isExternalClass( false ),
        _isAttributeClass( false ),
        _hasIndexer( false ),
        _hasPropertyControl( false )
    {
        _isClass = true;
    }

    ~ClassSymbol()
    {
    }

    void SetAsExternalClass() { _isExternalClass = true; }
    bool IsExternalClass() const
    { return _isExternalClass; }

    void setAsAttributeClass() { _isAttributeClass = true; }
    bool isAttributeClass() const
    { return _isAttributeClass; }

    void setIndexer() { _hasIndexer = true; }
    bool hasIndexer() const
    { return _hasIndexer; }

    void setPropertyControl() { _hasPropertyControl = true; }
    bool hasPropertyControl() const
    {
        return _hasPropertyControl;
    }

    NamespaceSymbolPtr getNamespace() 
    { 
        return _namespace; 
    }

    ClassSymbolPtr getParentClass() 
    { 
        return _parentClass; 
    }

    void setParentClass( ClassSymbolPtr parentClass ) 
    { 
        _parentClass = parentClass; 
    }

};

class ClassMemberSymbol: public ContainerSymbol
{
protected:
    ClassSymbolPtr _classSymbol;
    std::unordered_map<std::string, SymbolPtr> _arguments;

public:
    ClassMemberSymbol( std::string const& name, ClassSymbolPtr classSymbol ):
        ContainerSymbol( name ),
        _classSymbol( classSymbol )
    {
    }

    ~ClassMemberSymbol()
    {
    }

    std::unordered_map<std::string, SymbolPtr> const& getArguments() { return _arguments; }
    void addArgument( std::string const& name, std::string const& fullType );
};

class MethodSymbol: public ClassMemberSymbol
{
public:
    MethodSymbol( std::string const& name, ClassSymbolPtr classSymbol ):
        ClassMemberSymbol( name, classSymbol )
    {
    }

    ~MethodSymbol()
    {
    }
};

class PropertySymbol: public ClassMemberSymbol
{
private:
    SymbolPtr _typeSymbol;

public:
    PropertySymbol( std::string const& name, ClassSymbolPtr classSymbol ):
        ClassMemberSymbol( name, classSymbol )
    {
    }

    ~PropertySymbol()
    {
    }

    void setTypeSymbol( SymbolPtr symbol ) { _typeSymbol = symbol; }
    SymbolPtr getTypeSymbol() { return _typeSymbol; }
};

class ConstSymbol : public ClassMemberSymbol
{
private:
    SymbolPtr _typeSymbol;

public:
    ConstSymbol( std::string const& name, ClassSymbolPtr classSymbol ) :
        ClassMemberSymbol( name, classSymbol )
    {
    }

    ~ConstSymbol()
    {
    }

    void setTypeSymbol( SymbolPtr symbol ) { _typeSymbol = symbol; }
    SymbolPtr getTypeSymbol() { return _typeSymbol; }
};


class VarSymbol: public Symbol
{
private:
    ClassMemberSymbolPtr _memberSymbol;
    std::string _fullType;
public:
    VarSymbol( std::string const& name, std::string const& fullType, ClassMemberSymbolPtr memberSymbol ):
        Symbol( name ),
        _memberSymbol( memberSymbol ),
        _fullType( fullType )
    {
    }

    ~VarSymbol()
    {
    }

    std::string const& getType() const { return _fullType; }
};



}}}

