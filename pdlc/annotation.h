#pragma once

#include "symbols.h"
#include "ast.h"

namespace pam { namespace pdl
{
    ///////////////////////////////////////////////////////////////////////////////
    //  The annotation handler links the AST to a map of iterator positions
    //  for the purpose of subsequent semantic error handling when the
    //  program is being compiled.
    ///////////////////////////////////////////////////////////////////////////////
    template <typename Iterator>
    struct Annotation
    {
        template <typename, typename>
        struct result { typedef void type; };

        std::vector<Iterator>& iters;

        Annotation( std::vector<Iterator>& iters )
          : iters( iters ) {}

        struct setId
        {
            typedef void result_type;

            int id;
            setId(int id) : id(id) {}

            /*
            void operator()(ast::function_call& astNode ) const
            {
                x.function_name.id = id;
            }*/

            void operator()(ast::Identifier& astNode) const
            {
                astNode.id = id;
            }

            void operator()(ast::MethodNode& astNode) const
            {
                astNode.name.id = id;
            }

            void operator()(ast::PropertyNode& astNode) const
            {
                astNode.name.id = id;
            }

            template <typename T>
            void operator()(T& x) const
            {
                // no-op
            }
        };

        void operator()(ast::MemberNode& astNode, Iterator pos) const
        {
            int id = iters.size();
            iters.push_back( pos );

            boost::apply_visitor( setId( id ), astNode );
        }

        void operator()(ast::ClassNode& astNode, Iterator pos) const
        {
            int id = iters.size();
            iters.push_back( pos );
            astNode.name.id = id;
        }

        void operator()(ast::NamespaceNode& astNode, Iterator pos) const
        {
            int id = iters.size();
            iters.push_back( pos );

            if ( !astNode.name.empty() )
            {
                astNode.name.front().id = id;
            }
        }


        void operator()(ast::ArgumentNode& astNode, Iterator pos) const
        {
            int id = iters.size();
            iters.push_back(pos);
            if ( !astNode.type.empty() )
            {
                astNode.type.front().id = id;
            }
        }

        void operator()(ast::UsingNode& astNode, Iterator pos) const
        {
            int id = iters.size();
            iters.push_back(pos);
            if ( !astNode.className.empty() )
            {
                astNode.className.front().id = id;
            }
        }

        void operator()(ast::Identifier& ast, Iterator pos) const
        {
            const auto id = iters.size();
            iters.push_back( pos );
            ast.id = static_cast< int >(id);
        }
    };
}}
