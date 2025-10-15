#include "commonHeaders.hpp"
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
#include "astJsonWriter.h"

int main(int argc, char **argv)
{

    io::pdl::config::PdlConfig config(argc, argv);

    if (!config.isValid())
    {
        std::cerr << "Invalid arguments" << std::endl;
        std::cerr << "pdlc <inputfile> <template-folder> <class-template> <output-folder> [<pdl.config.json file>]" << std::endl;
#ifdef _DEBUG
        std::cin.ignore();
#endif

        return 1;
    }

    auto fileName = config.getInputFileName();

    std::ifstream in(fileName, std::ios_base::in);

    if (!in)
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
    io::pdl::ErrorHandler<IteratorType> errorHandler(fileName.c_str(), iter, end);

    io::pdl::PdlGrammar<IteratorType> grammar(errorHandler);
    io::pdl::ast::NamespaceNode ast;
    const io::pdl::skipper<IteratorType> skipper;

    const auto success = phrase_parse(iter, end, grammar, skipper, ast);

    auto exitCode = 0;

    if (success && iter == end)
    {
#ifdef _DEBUG
        std::cout << "Parsing succeeded!\n";
#endif
        auto isOk = false;

        io::pdl::codechecker::Checker checker(errorHandler);
        if (checker.visitNamespace(ast))
        {
            isOk = true;
            if (ast.isPending)
            {
                isOk = checker.resolveNamespace(ast);
            }

            if (isOk)
            {
                io::pdl::astjson::AstJsonWriter astWriter(config);
                astWriter.write(ast);
            }
        }

        if (!isOk)
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
    // std::cin.ignore();
#ifdef _DEBUG
    std::cin.ignore();
#endif

    return exitCode;
}

/*

    std::ostream& operator << (std::ostream& stream, const io::pdl::AccessModifiers am)
    {
        stream << "Access Modifier: " << static_cast<int>(am) << std::endl;
        return stream;
    }

    std::ostream& operator << (std::ostream& stream, const io::pdl::PropertyAccess pa)
    {
        stream << "Property Access: " << static_cast<int>(pa) << std::endl;
        return stream;
    }

*/
