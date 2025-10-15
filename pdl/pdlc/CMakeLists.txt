cmake_minimum_required(VERSION 3.20)

project(pdlc LANGUAGES CXX)

set(PDLC_SOURCES
    src/ast.cpp
    src/astJsonWriter.cpp
    src/codeChecker.cpp
    src/parser.cpp
    src/pdlc2.cpp
    src/pdlConfig.cpp
    src/symbols.cpp
    src/symbolTable.cpp
)

add_executable(pdlc ${PDLC_SOURCES})

set_target_properties(pdlc PROPERTIES
    CXX_STANDARD 20
    CXX_STANDARD_REQUIRED YES
    CXX_EXTENSIONS NO
)

target_include_directories(pdlc PRIVATE
    ${CMAKE_CURRENT_SOURCE_DIR}/src
)

target_compile_definitions(pdlc PRIVATE BOOST_ALL_NO_LIB)

find_package(Boost REQUIRED COMPONENTS filesystem date_time system)
target_include_directories(pdlc PRIVATE ${Boost_INCLUDE_DIRS})
target_link_libraries(pdlc PRIVATE ${Boost_LIBRARIES})

find_package(cpprestsdk REQUIRED)
target_link_libraries(pdlc PRIVATE cpprestsdk::cpprest)
