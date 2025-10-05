#include "stdafx.h"
#include "languageCommon.h"
#include "pdlConfig.h"
#include "skipper.h"
#include "symbols.h"
#include "ast.h"
#include "parserErrorHandler.h"
#include "symbolTable.h"
#include "codeChecker.h"
#include "annotation.h"
#include "parser.h"
#include "generators/generator.h"
#include "generators/csharpGenerator.h"
#include "generators/as3Generator.h"
#include "generators/CppGenerator.h"
#include "generators/phpGenerator.h"
#include "generators/jsGenerator.h"
#include "generators/javaGenerator.h"
#include "generators/goGenerator.h"

int main(int argc, char **argv)
{

    pam::pdl::config::PdlConfig config( argc, argv );

    if ( !config.isValid() )
    {
        std::cerr << "Invalid arguments" << std::endl;
        std::cerr << "pdlc <inputfile> <template-folder> <class-template> <output-folder> [<pdl.config.json file>]" << std::endl;
#ifdef _DEBUG
    std::cin.ignore();
#endif

        return 1;
    }

	auto fileName = config.getInputFileName();

    std::ifstream in( fileName, std::ios_base::in );

    if ( !in )
    {
        std::cerr << "Error: Could not open input file: " << fileName << std::endl;
#ifdef _DEBUG
    std::cin.ignore();
#endif

        return 1;
    }

	std::cout << "Compiling " << fileName << std::endl;

    in.unsetf(std::ios::skipws);

    std::string sourceCode;

    std::copy(
        std::istream_iterator<char>(in),
        std::istream_iterator<char>(),
        std::back_inserter(sourceCode));

    typedef std::string::const_iterator IteratorType;

    IteratorType iter = sourceCode.begin();
	const IteratorType end = sourceCode.end();
    pam::pdl::ErrorHandler<IteratorType> errorHandler( fileName.c_str(), iter, end );

    pam::pdl::PdlGrammar<IteratorType> grammar( errorHandler );
    pam::pdl::ast::NamespaceNode ast;
	const pam::pdl::skipper<IteratorType> skipper;

	const auto success = phrase_parse( iter, end, grammar, skipper, ast );

	auto exitCode = 0;

    if ( success && iter == end )
    {
#ifdef _DEBUG
        std::cout << "Parsing succeeded!\n";
#endif
	    auto isOk = false;

        pam::pdl::codechecker::Checker checker( errorHandler );
        if ( checker.visitNamespace( ast ) )
        {
            isOk = true;
            if ( ast.isPending )
            {
                isOk = checker.resolveNamespace( ast );
            }

            if ( isOk )
            {

				if (pam::pdl::codegen::Generator::isEnabled(config, "csharp"))
				{
#ifdef _DEBUG
					std::cout << "Generating C# Classes...\n";
#endif
					pam::pdl::codegen::CSharpGenerator csGenerator(config);
					csGenerator.doNamespace(ast);
				}

				if (pam::pdl::codegen::Generator::isEnabled(config, "as3"))
				{
#ifdef _DEBUG
					std::cout << "Generating ActionScript Classes...\n";
#endif
					pam::pdl::codegen::As3Generator as3Generator(config);
					as3Generator.doNamespace(ast);
				}
                
                //C/C++
                /*
                std::cout << "Generating C/C++ Classes...\n";
                pam::pdl::codegen::CppGenerator cppGenerator( config );
                cppGenerator.doNamespace( ast );
                */
                
                //PHP
#ifdef _DEBUG
                std::cout << "Generating PHP Classes...\n";
#endif
				if ( pam::pdl::codegen::Generator::isEnabled( config, "php"))
				{
					pam::pdl::codegen::PhpGenerator phpGenerator(config);
					phpGenerator.doNamespace(ast);
				}
                

                //Javascript
#ifdef _DEBUG
                std::cout << "Generating Javascript Classes...\n";
#endif
				if (pam::pdl::codegen::Generator::isEnabled(config, "js"))
				{
					pam::pdl::codegen::JsGenerator jsGenerator(config);
					jsGenerator.doNamespace(ast);
				}

				if (pam::pdl::codegen::Generator::isEnabled(config, "go"))
				{
					pam::pdl::codegen::GoGenerator goGenerator(config);
					goGenerator.doNamespace(ast);
				}

                //Java
                //std::cout << "Generating Java Classes...\n";
                //pam::pdl::codegen::JavaGenerator javaGenerator( config );
                //javaGenerator.doNamespace( ast );
                //
#ifdef _DEBUG
                std::cout << "Done!\n";
#endif
            }
            
        }

        if ( !isOk )
        {
            std::cerr << "Error(s) found!\n";
			exitCode = 1;
        }

    }
    else
    {
        std::cerr << "Parsing failed\n";
        exitCode = 1;
    }
	//std::cin.ignore();
#ifdef _DEBUG
    std::cin.ignore();
#endif

    return exitCode;
}

/*

    std::ostream& operator << (std::ostream& stream, const pam::pdl::AccessModifiers am)
    {
        stream << "Access Modifier: " << static_cast<int>(am) << std::endl;
        return stream;
    }

    std::ostream& operator << (std::ostream& stream, const pam::pdl::PropertyAccess pa)
    {
        stream << "Property Access: " << static_cast<int>(pa) << std::endl;
        return stream;
    }

*/

