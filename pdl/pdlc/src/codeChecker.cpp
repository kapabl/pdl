#include "commonHeaders.hpp"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "codeChecker.h"

using namespace io::pdl::symbols;
using namespace io::pdl::ast;
using namespace io::pdl::codechecker;

bool Checker::visitNamespace( NamespaceNode& astNamespace )
{
	auto result = true;
	const auto namespaceSymbol = _symbolTable->addNamespace( astNamespace );
	astNamespace.symbol = namespaceSymbol;

	const auto fullNamespace = SymbolTable::joinIdentifier( astNamespace.name );
	_namespaces[ fullNamespace ] = namespaceSymbol;

	enterNamespaceScope( namespaceSymbol );

	result = visitUsingList( astNamespace.usings ) && result;
	result = visitClassList( astNamespace.classes ) && result;

	astNamespace.isPending = astNamespace.classes.isPending;

	exitNamespaceScope();

	return result;
}


bool Checker::visitUsingList( UsingList& usingList )
{
	auto result = true;

	for ( auto& usingLine : usingList )
	{
		result = visitUsing( usingLine ) && result;
	}
	return result;
}

bool Checker::visitUsing( UsingNode& usingNode )
{
	auto result = true;

	auto parsedClassName = SymbolTable::parseFullClassName( usingNode.className );

	auto& fullNamespace = parsedClassName.first;
	auto namespaceSymbol( _symbolTable->getNamespace( fullNamespace ) );
	_namespaces[ fullNamespace ] = namespaceSymbol;

	auto& className = parsedClassName.second;

	if ( !namespaceSymbol->symbolTable()->classExists( className ) )
	{
		usingNode.symbol = namespaceSymbol->symbolTable()->addUsingClass( usingNode );

		const auto fullClassName = fullNamespace + "." + className;
		_classes[ fullClassName ] = usingNode.symbol;
	}
	else
	{
		_errorHandler( usingNode.className.front().id, "Class redeclared: " + className );
		result = false;
	}
	return result;
}

bool Checker::visitClassList( ClassList& classList )
{
	auto result = true;

	for ( auto& astClassNode : classList )
	{
		result = visitClass( astClassNode ) && result;
		classList.isPending = classList.isPending || astClassNode.isPending;
	}
	return result;
}

SymbolPtr Checker::useType( FullIdentifierNode const& astFullIdentifier )
{
	auto result = isIntrinsicType( astFullIdentifier )
		? _symbolTable->getSymbol<SymbolPtr>( SymbolTable::joinIdentifier( astFullIdentifier ) )
		: useClass( astFullIdentifier, false );


	return result;
}

ClassSymbolPtr Checker::useClass( FullIdentifierNode const& className, bool excludeCurrentClass )
{
	auto classSymbol = getClassSymbol( className, excludeCurrentClass );

	if ( classSymbol )
	{
		classSymbol->incUseCount();
	}

	return classSymbol;
}

bool Checker::classExists( FullIdentifierNode const& className )
{
	const auto classSymbol = getClassSymbol( className, false );
	const auto result = static_cast<bool>( classSymbol );
	return result;
}

ClassSymbolPtr Checker::getClassSymbol( FullIdentifierNode const& className, bool excludeCurrentClass )
{
	ClassSymbolPtr result;

	auto parsedClassName = SymbolTable::parseFullClassName( className );
	auto& namespaceName = parsedClassName.first;
	auto& singleClassName = parsedClassName.second;

	if ( namespaceName.empty() )
	{
		for ( const auto namespaceItem : _namespaces )
		{
			const auto fullClassName = namespaceItem.first + "." + singleClassName;

            auto excluded = excludeCurrentClass && fullClassName == _currentFullClassName;

            if ( !excluded )
            {
                const auto iter = _classes.find( fullClassName );
                if ( iter != _classes.end() )
                {
                    result = iter->second;
                    break;
                }
            }
		}
	}
	else
	{
		const auto fullClassName = namespaceName + "." + singleClassName;

		const auto iter = _classes.find( fullClassName );
		if ( iter != _classes.end() )
		{
			result = iter->second;
		}
	}


	return result;
}

bool Checker::visitParentClass( ClassNode& classNode )
{
	auto result = true;

    const auto parentClassSymbol = useClass( classNode.parentClass.get(), true );
    if ( parentClassSymbol )
    {
        classNode.symbol->setParentClass( parentClassSymbol );
    }
    else
    {
        result = false;
    }

    classNode.parentClass.get().isPending = !result;

	return result;
}

bool Checker::visitClass( ClassNode& classNode )
{
	auto result = true;

	auto const& className = classNode.name.name;
	const auto fullClassName = _currentNamespace->name() + "." + className;

	//checkAccessModifier( astClass.accessModifier );

	if ( !_currentNamespace->symbolTable()->classExists( className ) )
	{
		auto classSymbol = _currentNamespace->symbolTable()->addClass( classNode, _currentNamespace );
		classSymbol->incUseCount();

		_classes[ fullClassName ] = classSymbol;
		classNode.symbol = classSymbol;

        if ( isNamespace( fullClassName ) )
        {
            _errorHandler( classNode.name.id, "Invalid class name, there is a namespace already with this name: " + className );
            result = false;
        }

        enterClassScope( classSymbol );
		if ( classNode.parentClass )
		{
            visitParentClass( classNode );
		}

		result = visitMembers( classNode.members ) && result;

		exitClassScope();
	}
	else
	{
		_errorHandler( classNode.name.id, "Class redeclared: " + className );
		result = false;
	}

	classNode.isPending = classNode.members.isPending || ( classNode.parentClass && classNode.parentClass.get().isPending );

	return result;
}

bool Checker::visitMembers( MemberList& astMemberList )
{
	auto result = true;

	Checker::MemberVisitor memberVisitor( *this, astMemberList );

	for ( auto& member : astMemberList )
	{
		result = memberVisitor( member ) && result;
	}

	return result;
}

bool Checker::visitAttributeList( AttributeListNode& astAttributeList )
{
	auto result = true;

	for ( auto& attribute : astAttributeList )
	{
		result = visitAttribute( attribute ) || result;
	}

	return result;
}

bool Checker::visitAttribute( AttributeNode& astAttribute )
{
	auto classSymbol = useClass( astAttribute.name, true );
	const auto result = classSymbol.get() != nullptr;
	if ( classSymbol )
	{
		classSymbol->setAsAttributeClass();
	}

	astAttribute.isPending = !result || astAttribute.isPending;

	return result;
}

bool Checker::resolveAttribute( AttributeNode& astAttribute )
{
	auto classSymbol = useClass( astAttribute.name, true );
	const auto result = classSymbol.get() != nullptr;
	if ( result )
	{
		classSymbol->setAsAttributeClass();
	}
	else
	{
		_errorHandler( astAttribute.name.front().id, "Attribute class not found: " + astAttribute.name );
	}

	return result;
}


//bool Checker::visitPropertyType( PropertyNode& propertyNode )
//{
//	auto result = false;
//
//	const auto typeSymbol = useType( propertyNode.propertyType.type );
//
//	if ( typeSymbol )
//	{
//		propertyNode.symbol->setTypeSymbol( typeSymbol );
//		result = true;
//	}
//
//	return result;
//}


bool Checker::visitProperty( PropertyNode& propertyNode )
{
	auto result = true;

	if ( !propertyExists( _currentClassSymbol, propertyNode ) )
	{
		visitPropertyAccess( propertyNode.access );

		const auto propertySymbol = _currentClassSymbol->symbolTable()->addProperty( propertyNode, _currentClassSymbol );
		propertyNode.symbol = propertySymbol;
		propertyNode.propertyType.isPending = !visitPropertyType( propertyNode );

		enterClassMemberScope( propertySymbol );
        visitAttributeList( propertyNode.attributes );

		//NOTE: this is a simplification to allow easier code generation
		//in multiple languages
		if ( propertyNode.arguments.size() > 1 )
		{
			_errorHandler( propertyNode.name.id, "Properties can't have more than 1 argument: " + propertyNode.name.name );
			result = false;
		}

		if ( propertyNode.arguments.size() > 0 )
		{
			if ( _currentClassSymbol->hasIndexer() )
			{
				_errorHandler( propertyNode.name.id,
					"One indexer or property with argument allowed per class(c# limitation): " + propertyNode.name.
					name );
				result = false;
			}
			_currentClassSymbol->setIndexer();
		}

		result = visitArgumentList( propertyNode.arguments ) && result;

        if ( propertyNode.name.name == PROPERTY_CONTROL )
        {
            _currentClassSymbol->setPropertyControl();
        }

		exitClassMemberScope();
	}
	else
	{
		_errorHandler( propertyNode.name.id, "Property or method redeclared: " + propertyNode.name.name );
		result = false;
	}

	propertyNode.isPending =
		propertyNode.propertyType.isPending ||
		propertyNode.arguments.isPending ||
		propertyNode.attributes.isPending;

	return result;
}

bool Checker::visitPropertyType( PropertyNode& propertyNode )
{
	auto result = false;

	const auto typeSymbol = useType( propertyNode.propertyType.type );

	if ( typeSymbol )
	{
		propertyNode.symbol->setTypeSymbol( typeSymbol );
		result = true;
	}

	return result;
}

bool Checker::visitPropertyType( ShortPropertyNode& shortPropertyNode )
{
	auto result = false;

	const auto typeSymbol = useType( shortPropertyNode.propertyType.type );

	if ( typeSymbol )
	{
		shortPropertyNode.symbol->setTypeSymbol( typeSymbol );
		result = true;
	}

	return result;
}


bool Checker::visitConst( ConstNode& constNode )
{
    auto result = true;

    if ( !constExists( _currentClassSymbol, constNode ) )
    {
        const auto constSymbol = _currentClassSymbol->symbolTable()->addConst( constNode, _currentClassSymbol );
        constNode.symbol = constSymbol;

        constNode.type.isPending = !useType( constNode.type );

    }
    else
    {
        _errorHandler( constNode.name.id, "Property, method or const re-declared: " + constNode.name );
        result = false;
    }

    constNode.isPending = constNode.type.isPending;

    return result;
}

bool Checker::visitMethod( MethodNode& astMethod )
{
	auto result = true;

	if ( !methodExists( _currentClassSymbol, astMethod ) )
	{
		const auto methodSymbol = _currentClassSymbol->symbolTable()->addMethod( astMethod, _currentClassSymbol );
		astMethod.symbol = methodSymbol;

		astMethod.type.isPending = !useType( astMethod.type );

		enterClassMemberScope( methodSymbol );

		result = visitArgumentList( astMethod.arguments ) && result;

		exitClassMemberScope();
	}
	else
	{
		_errorHandler( astMethod.name.id, "Property or method redeclared: " + astMethod.name );
		result = false;
	}

	astMethod.isPending = astMethod.type.isPending || astMethod.arguments.isPending;

	return result;
}


bool Checker::visitArgument( ArgumentNode& astArgument )
{
	auto result = true;

	if ( !argumentExists( _currentMember, astArgument ) )
	{
		const auto fullType = SymbolTable::joinIdentifier( astArgument.type );
		_currentMember->addArgument( astArgument.name.name, fullType );

		astArgument.type.isPending = !useType( astArgument.type );
	}
	else
	{
		_errorHandler( astArgument.name.id, "Argument name redeclared: " + astArgument.name );
		result = false;
	}

	astArgument.isPending = astArgument.type.isPending;

	return result;
}

bool Checker::visitArgumentList( ArgumentList& astArgumentList )
{
	auto result = true;

	for ( auto& argument : astArgumentList )
	{
		result = visitArgument( argument ) && result;
		astArgumentList.isPending = astArgumentList.isPending || argument.isPending;
	}
	return result;
}

bool Checker::methodExists( ClassSymbolPtr classSymbol, MethodNode& astMethod )
{
	auto classSymbols = classSymbol->symbolTable();

	const auto result = nullptr != classSymbols->getSymbol<MethodSymbolPtr>( astMethod.name.name ).get();

	return result;
}

bool Checker::constExists( ClassSymbolPtr classSymbol, ConstNode& constNode )
{
	auto classSymbols = classSymbol->symbolTable();

	const auto result = nullptr != classSymbols->getSymbol<ConstSymbolPtr>( constNode.name.name ).get();

	return result;
}

bool Checker::propertyExists( ClassSymbolPtr classSymbol, PropertyNode& astProperty )
{
	auto classSymbols = classSymbol->symbolTable();

	const auto result = nullptr != classSymbols->getSymbol<PropertySymbolPtr>( astProperty.name.name ).get();

	return result;
}

bool Checker::propertyExists( ClassSymbolPtr classSymbol, ShortPropertyNode& shortPropertyNode )
{
	auto classSymbols = classSymbol->symbolTable();

	const auto result = nullptr != classSymbols->getSymbol<PropertySymbolPtr>( shortPropertyNode.name.name ).get();

	return result;
}

bool Checker::argumentExists( ClassMemberSymbolPtr memberSymbol, ArgumentNode& astArgument )
{
	auto arguments = memberSymbol->getArguments();

	const auto result = arguments.find( astArgument.name.name ) != arguments.end();

	return result;
}

bool Checker::isIntrinsicType( FullIdentifierNode const& astFullIndentifier )
{
	auto result = false;

	auto fullType = SymbolTable::joinIdentifier( astFullIndentifier );

	if ( fullType.find( '.' ) == std::string::npos )
	{
		result = _intrinsicTypes.find( fullType ) != _intrinsicTypes.end();
	}

	return result;
}

bool Checker::isNamespace( std::string const& fullNamespace )
{
	return _namespaces.find( fullNamespace ) != _namespaces.end();
}

bool Checker::resolveNamespace( NamespaceNode& astNamespace )
{
	enterNamespaceScope( astNamespace.symbol );

	const auto result = resolveClassList( astNamespace.classes );

	exitNamespaceScope();

	return result;
}

bool Checker::resolveClassList( ClassList& astClassList )
{
	auto result = true;

	for ( auto& astClassNode : astClassList )
	{
		if ( astClassNode.isPending )
		{
			result = resolveClass( astClassNode ) && result;
		}
	}
	return result;
}

bool Checker::resolveClass( ClassNode& astClass )
{
	auto result = true;
	if ( astClass.parentClass && astClass.parentClass.get().isPending )
	{
		if ( resolveParentClass( astClass ) )
		{
			const auto classSymbol = useClass( astClass.parentClass.get(), true );
			astClass.symbol->setParentClass( classSymbol );
		}
		else
		{
			_errorHandler( astClass.parentClass.get().front().id, "Parent class not found: " + astClass.name.name );
			result = false;
		}
	}

	if ( astClass.members.isPending )
	{
		enterClassScope( astClass.symbol );
		result = resolveMemberList( astClass.members ) && result;
		exitClassScope();
	}

	return result;
}


bool Checker::resolveMemberList( MemberList& astMemberList )
{
	auto result = true;

	Checker::ResolveMemberVisitor memberVisitor( *this, astMemberList );

	for ( auto& member : astMemberList )
	{
		result = memberVisitor( member ) && result;
	}

	return result;
}


bool Checker::resolveArgument( ArgumentNode& astArgument )
{
	auto result = nullptr != useType( astArgument.type ).get();

	if ( !result )
	{
		_errorHandler( astArgument.type.front().id, "Argument type not found: " + astArgument.type );
		result = false;
	}
	return result;
}


bool Checker::resolveArgumentList( ArgumentList& astArgumentList )
{
	auto result = true;

	for ( auto& argument : astArgumentList )
	{
		if ( argument.isPending )
		{
			result = resolveArgument( argument ) && result;
		}
	}
	return result;
}


bool Checker::resolvePropertyType( PropertyType& propertyType )
{
	const auto result = nullptr != useType( propertyType.type ).get();
	if ( !result )
	{
		_errorHandler( propertyType.type.front().id,
			"Property type not found: " + propertyType.type );
	}

	return result;
}

bool Checker::resolveProperty( PropertyNode& propertyNode )
{
	auto result = true;
	if ( propertyNode.propertyType.isPending )
	{
		result = resolvePropertyType( propertyNode.propertyType );
	}


	if ( propertyNode.arguments.isPending )
	{
		enterClassMemberScope( propertyNode.symbol );

		result = resolveArgumentList( propertyNode.arguments );

		exitClassMemberScope();
	}

	if ( propertyNode.attributes.isPending )
	{
		resolveAttributeList( propertyNode.attributes );
	}

	return result;
}

bool Checker::resolveProperty( ShortPropertyNode& shortPropertyNode )
{
	auto result = true;
	if ( shortPropertyNode.propertyType.isPending )
	{
		result = resolvePropertyType( shortPropertyNode.propertyType );
	}

	return result;
}

bool Checker::visitPropertyAccess( boost::optional<io::pdl::PropertyAccess>& access )
{
	if ( !access )
	{
		access = io::pdl::PropertyAccess::paReadWrite;
	}
	return true;
}

bool Checker::visitShortProperty( ShortPropertyNode& shortPropertyNode )
{
	auto result = true;

	if ( !propertyExists( _currentClassSymbol, shortPropertyNode ) )
	{

		const auto propertySymbol = _currentClassSymbol->symbolTable()->addProperty( shortPropertyNode, _currentClassSymbol );
		shortPropertyNode.symbol = propertySymbol;
		shortPropertyNode.propertyType.isPending = !visitPropertyType( shortPropertyNode );

	}
	else
	{
		_errorHandler( shortPropertyNode.name.id, "Property or method redeclared: " + shortPropertyNode.name.name );
		result = false;
	}

	shortPropertyNode.isPending = shortPropertyNode.propertyType.isPending;

	return result;
}

bool Checker::resolveAttributeList( AttributeListNode& astAttributeList )
{
	auto result = true;

	for ( auto& attribute : astAttributeList )
	{
		if ( attribute.isPending )
		{
			result = resolveAttribute( attribute ) || result;
		}
	}

	return result;
}

bool Checker::resolveMethod( MethodNode& astMethod )
{
	auto result = true;
	if ( astMethod.type.isPending && !useType( astMethod.type ) )
	{
		_errorHandler( astMethod.type.front().id, "Method return type not found: " + astMethod.type );
		result = false;
	}

	if ( astMethod.arguments.isPending )
	{
		enterClassMemberScope( astMethod.symbol );

		result = resolveArgumentList( astMethod.arguments );

		exitClassMemberScope();
	}

	return result;
}

bool Checker::resolveConst( ConstNode& constNode )
{
    auto result = true;
    if ( constNode.type.isPending && !useType( constNode.type ) )
    {
        _errorHandler( constNode.type.front().id, "Const type not found: " + constNode.type );
        result = false;
    }

    return result;
}

bool Checker::resolveParentClass( ClassNode& astClass )
{
	const auto result = nullptr != useClass( astClass.parentClass.get(), true ).get();

	return result;
}



