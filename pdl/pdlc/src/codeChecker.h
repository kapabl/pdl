#pragma once


namespace io { namespace pdl { namespace codechecker
{
	using namespace io::pdl::symbols;

	struct Checker
	{
		GlobalSymbolTablePtr _symbolTable;

		template <typename ErrorHandler>
		explicit Checker( ErrorHandler& errorHandler ) :
			_symbolTable( new GlobalSymbolTable() )
		{
			//using namespace boost::phoenix::arg_names;
			namespace phx = boost::phoenix;
			using boost::phoenix::function;

			_errorHandler = function<ErrorHandler>( errorHandler )(
				"Error! ", phx::arg_names::_2, phx::cref( errorHandler.iters )[phx::arg_names::_1] );

			addIntrinsicType( ITYPE_VOID );
			addIntrinsicType( ITYPE_BOOL );
			addIntrinsicType( ITYPE_ARRAY );
			addIntrinsicType( ITYPE_INT );
			addIntrinsicType( ITYPE_UINT );
			addIntrinsicType( ITYPE_DOUBLE );
			addIntrinsicType( ITYPE_FUNCTION );
			addIntrinsicType( ITYPE_STRING );
			addIntrinsicType( ROOT_OBJECT );
		}


		bool visitNamespace( ast::NamespaceNode& astNamespace );
		bool visitClass( ast::ClassNode& astClass );
		bool visitClassList( ast::ClassList& astClassList );
		bool visitUsing( ast::UsingNode& astUsingN );
		bool visitUsingList( ast::UsingList& astUsingList );
		bool visitMembers( ast::MemberList& astMemberList );
		bool visitArgument( ast::ArgumentNode& astArgument );
		bool visitArgumentList( ast::ArgumentList& astArgumentList );

		bool visitProperty( ast::PropertyNode& astProperty );
		bool visitShortProperty( ast::ShortPropertyNode& shortPropertyNode );

		bool visitPropertyType( ast::PropertyNode& propertyNode );
		bool visitPropertyType( ast::ShortPropertyNode& shortPropertyNode );

		bool visitMethod( ast::MethodNode& astMethod );
		bool visitConst( ast::ConstNode& constNode );
		bool visitParentClass( ast::ClassNode& astClass );
		bool visitAttributeList( ast::AttributeListNode& astAttributeList );
		bool visitAttribute( ast::AttributeNode& astAttribute );
		static bool visitPropertyAccess( boost::optional<io::pdl::PropertyAccess>& access );


		struct MemberVisitor
		{
			typedef bool result_type;

			Checker& _checker;

			ast::MemberList& _members;

			MemberVisitor( Checker& checker, ast::MemberList& members ) :
				_checker( checker ),
				_members( members )
			{
			}

			bool operator()( ast::MemberNode& astMember ) { return boost::apply_visitor( *this, astMember ); }

			bool operator()( ast::PropertyNode& astProperty ) const
			{
				const auto result = _checker.visitProperty( astProperty );
				_members.isPending = _members.isPending || astProperty.isPending;
				return result;
			}

			bool operator()(ast::ShortPropertyNode& shortPropertyNode) const
			{
				const auto result = _checker.visitShortProperty(shortPropertyNode);
				_members.isPending = _members.isPending || shortPropertyNode.isPending;
				return result;
			}

			bool operator()( ast::MethodNode& astMethod ) const
			{
				const auto result = _checker.visitMethod( astMethod );
				_members.isPending = _members.isPending || astMethod.isPending;
				return result;
			}

			bool operator()( ast::ConstNode & constNode ) const
			{
				const auto result = _checker.visitConst( constNode );
				_members.isPending = _members.isPending || constNode.isPending;
				return result;
			}
		};

		bool resolveNamespace( ast::NamespaceNode& astNamespace );
		bool resolveClass( ast::ClassNode& astClass );
		bool resolveClassList( ast::ClassList& astClassList );
		bool resolveMemberList( ast::MemberList& astMemberList );
		bool resolveArgument( ast::ArgumentNode& astArgument );
		bool resolveArgumentList( ast::ArgumentList& astArgumentList );
		bool resolvePropertyType( ast::PropertyType& propertyType );
		bool resolveProperty( ast::PropertyNode& astProperty );
		bool resolveProperty( ast::ShortPropertyNode& astProperty );
		bool resolveConst( ast::ConstNode& constNode );
		bool resolveMethod( ast::MethodNode& methodNode );
		bool resolveParentClass( ast::ClassNode& astClass );
		bool resolveAttributeList( ast::AttributeListNode& astAttributeList );
		bool resolveAttribute( ast::AttributeNode& astAttribute );

		ClassSymbolPtr useClass( ast::FullIdentifierNode const& className, bool excludeCurrentClass );

		SymbolPtr useType( ast::FullIdentifierNode const& className );

		struct ResolveMemberVisitor
		{
			typedef bool result_type;

			Checker& checker;
			ast::MemberList& members;

			ResolveMemberVisitor( Checker& checker, ast::MemberList& members ) :
				checker( checker ),
				members( members )
			{
			}

			//boost::variant visitors
			bool operator()( ast::MemberNode& memberNode ) { return boost::apply_visitor( *this, memberNode ); }

			bool operator()( ast::PropertyNode& propertyNode ) const
			{
				auto result = true;
				if ( propertyNode.isPending )
				{
					result = checker.resolveProperty( propertyNode );
				}
				return result;
			}

			bool operator()(ast::ShortPropertyNode& propertyNode) const
			{
				auto result = true;
				if (propertyNode.isPending)
				{
					result = checker.resolveProperty(propertyNode);
				}
				return result;
			}

			bool operator()( ast::MethodNode& methodNode ) const
			{
				auto result = true;
				if ( methodNode.isPending )
				{
					result = checker.resolveMethod( methodNode );
				}
				return result;
			}

            bool operator()( ast::ConstNode& constNode ) const
            {
                auto result = true;
                if ( constNode.isPending )
                {
                    result = checker.resolveConst( constNode );
                }
                return result;
            }
		};


		bool classExists( ast::FullIdentifierNode const& className );

		static bool propertyExists( ClassSymbolPtr classSymbol, ast::PropertyNode& astProperty );
		static bool propertyExists( ClassSymbolPtr classSymbol, ast::ShortPropertyNode& astProperty );
		static bool methodExists( ClassSymbolPtr classSymbol, ast::MethodNode& astMethod );
		static bool constExists( ClassSymbolPtr classSymbol, ast::ConstNode& constNode );

		static bool argumentExists( ClassMemberSymbolPtr memberSymbol, ast::ArgumentNode& astArgument );

		bool isIntrinsicType( ast::FullIdentifierNode const& astFullIdentifier );

		ClassSymbolPtr getClassSymbol( ast::FullIdentifierNode const& astFullIdentifier, bool excludeCurrentClass );
		bool isNamespace( std::string const& fullNamespace );
		static bool isFullyQualified( std::string const& name ) { return name.find( '.' ) != std::string::npos; }

		void enterNamespaceScope( NamespaceSymbolPtr symbol ) { _currentNamespace = symbol; }
		void exitNamespaceScope() { _currentNamespace.reset(); }

		void enterClassScope( ClassSymbolPtr classSymbol )
		{
		    _currentClassSymbol = classSymbol;
            _currentFullClassName = classSymbol->getNamespace()->name() + "." + classSymbol->name();
		}

		void exitClassScope()
		{
		    _currentClassSymbol.reset();
            _currentFullClassName.clear();
		}

		void enterClassMemberScope( ClassMemberSymbolPtr symbol ) { _currentMember = symbol; }
		void exitClassMemberScope() { _currentMember.reset(); }

	private:

		boost::function<void( int tag, std::string const& what )> _errorHandler;

		NamespaceSymbolPtr _currentNamespace;
		ClassSymbolPtr _currentClassSymbol;
        std::string _currentFullClassName;
		ClassMemberSymbolPtr _currentMember;

		std::unordered_set<std::string> _intrinsicTypes;
		std::unordered_map<std::string, ClassSymbolPtr> _classes;
		std::unordered_map<std::string, NamespaceSymbolPtr> _namespaces;

		void addIntrinsicType( std::string const& type )
		{
			_intrinsicTypes.insert( type );
			_symbolTable->addIntrinsicType( type );
		}
	};
}}}
