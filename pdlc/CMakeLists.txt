set(PDLC_SOURCES
    ast.cpp
    codeChecker.cpp
    parser.cpp
    pdlc2.cpp
    pdlConfig.cpp
    symbols.cpp
    symbolTable.cpp
    stdafx.cpp
    generators/as3Generator.cpp
    generators/CppGenerator.cpp
    generators/csharpGenerator.cpp
    generators/generator.cpp
    generators/javaGenerator.cpp
    generators/jsGenerator.cpp
    generators/phpGenerator.cpp
    generators/goGenerator.cpp
)

add_executable(pdlc ${PDLC_SOURCES})

set_target_properties(pdlc PROPERTIES
    CXX_STANDARD 17
    CXX_STANDARD_REQUIRED YES
    CXX_EXTENSIONS NO
)

target_include_directories(pdlc PRIVATE
    ${CMAKE_CURRENT_SOURCE_DIR}
)

target_compile_definitions(pdlc PRIVATE BOOST_ALL_NO_LIB)

find_package(Boost CONFIG REQUIRED COMPONENTS filesystem date_time system format)
target_link_libraries(pdlc PRIVATE Boost::filesystem Boost::date_time Boost::system Boost::format)

find_package(cpprestsdk REQUIRED)
target_link_libraries(pdlc PRIVATE cpprestsdk::cpprest)
