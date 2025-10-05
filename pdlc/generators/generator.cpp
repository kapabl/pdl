#include "stdafx.h"
#include "pdlConfig.h"
#include "languageCommon.h"
#include "symbols.h"
#include "ast.h"
#include "symbolTable.h"
#include "generator.h"

using namespace pam::pdl;
using namespace config;
using namespace symbols;
using namespace codegen;
using namespace ast;

using std::string;
namespace fs = boost::filesystem;


Generator::Generator( PdlConfig const& config,  string const& language ) : 
    _config( config ),
    _language( language ),
    _literalVisitor( *this ),
    _memberVisitor( *this ),
    _attrParamsVisitor( *this ),
    _hasPropertyControl( false )

{
	_isSameSource = false;
    _accessModifierMap[ AccessModifiers::amInternal ] = ACCESS_MOD_INTERNAL_STR;
    _accessModifierMap[ AccessModifiers::amProtected ] = ACCESS_MOD_PROTECTED_STR;
    _accessModifierMap[ AccessModifiers::amPrivate ] = ACCESS_MOD_PRIVATE_STR;
    _accessModifierMap[ AccessModifiers::amPublic ] = ACCESS_MOD_PUBLIC_STR;


    addScalarType( ITYPE_INT );
    addScalarType( ITYPE_UINT );
    addScalarType( ITYPE_DOUBLE );
    addScalarType( ITYPE_STRING );

    _nullObjectValue = "null";

    readConfig();

	const fs::path outputFolder( _outputFolder );
    if ( !fs::is_directory( _outputFolder ) )
    {
        fs::create_directories(  outputFolder );
    }

    _formatter.attr = "[%1%(%2%)]";
    _formatter.paramlessAttr = "[%1%]";
    _attributeSeparator = "\n";

}

void Generator::readConfig()
{
    auto jsonConfig = _config.getJsonConfig();

    _indent = PdlConfig::to_string( jsonConfig[U("indent")].as_string() );
    _outputFolder = PdlConfig::to_string( jsonConfig[U("out")].as_string() );
    _inputFolder = PdlConfig::to_string( jsonConfig[U("in")].as_string() );
    _classTemplate = PdlConfig::to_string( jsonConfig[U("classTemplate")].as_string() );
	const auto configSectionName = PdlConfig::to_wstring( _language );

    if ( jsonConfig[U("file")].has_field( configSectionName ) )
    {
        _languageConfig = jsonConfig[U("file")][ configSectionName ];
    }

}

string Generator::getTargetLanguageNamespace( NamespaceNode const& astNamespace )
{
    auto result = SymbolTable::joinIdentifier( astNamespace.name );
    return result;
}

string Generator::doNamespace( NamespaceNode const& astNamespace )
{

    _mainNamespace = getTargetLanguageNamespace( astNamespace );
	_pdlMainNamespace = SymbolTable::joinIdentifier( astNamespace.name );

    for(auto const& astClass : astNamespace.classes )
    {
        _methods.clear();
        _properties.clear();
        _outputHeader = createFileHeader( _mainNamespace + "." + astClass.name.name );

	    const auto classResult = doClass( astNamespace, astClass );

        prepareOutputFolder( astNamespace, astClass );

        outputClass( classResult, astClass );

    }

    return "";
}

bool Generator::isEnabled( PdlConfig& pdlConfig, string const& language )
{
	auto result = true;
	auto jsonConfig = pdlConfig.getJsonConfig();

	const auto languageSectionKey = PdlConfig::to_wstring( language );

	if (jsonConfig[U("file")].has_field(languageSectionKey))
	{
		auto languageSection = jsonConfig[U("file")][languageSectionKey];
		result = PdlConfig::as_bool(languageSection, "enabled", true);
	}
	return result;
}

void Generator::prepareOutputFolder( NamespaceNode const& astNamespace, ClassNode const& astClass )
{
	const auto outputPath = getFileOutputFolder();

    if ( !is_directory( outputPath ) )
    {
        create_directories(  outputPath );
    }

    _outputFileName = (outputPath / ( astClass.name.name  + _fileExt )).string();
}

string Generator::doArgumentList( ArgumentList const& astArgumentList )
{
    std::vector<string> argumentVector;

    for(auto const& argument : astArgumentList )
    {
        argumentVector.push_back(doArgument( argument ) );
    }

	auto result = boost::join( argumentVector, ", " );

    return result;
}
    
string Generator::getAccessModifier( AccessModifiers const accessModifier )
{
    return _accessModifierMap[ accessModifier ];
}


string Generator::readTemplateCode()
{
    auto templateFile = fs::path( _inputFolder ) / ( _classTemplate );

    std::ifstream in( templateFile.string(), std::ios_base::in );
    in.unsetf( std::ios::skipws );

    string result;
    std::copy( std::istream_iterator<char>(in), std::istream_iterator<char>(), std::back_inserter( result ) );

    return result;
}


string Generator::doAttributes( AttributeListNode const& astAttributeList )
{
    std::vector<string> attributeVector;

    for(auto const& attribute : astAttributeList )
    {
	    const string params = boost::apply_visitor( _attrParamsVisitor, attribute.params );
	    const auto attrInfo = processAttributeName( attribute.name );
        if ( params == "" )
        {
            
            attributeVector.push_back( 
                (boost::format( _formatter.paramlessAttr ) 
                % attrInfo.name
                ).str() );
        }
        else
        {
            attributeVector.push_back( 
                (boost::format( _formatter.attr ) 
                    % attrInfo.name
                    % params).str() );
        }
    }

	auto result = boost::join( attributeVector, _attributeSeparator );

    return result;
}

string Generator::visitAttrRequiredParams( AttrRequiredParams const& astAttrRequireParams )
{
    std::vector<string> paramVector;
    for(auto const& param : astAttrRequireParams )
    {
        paramVector.push_back( boost::apply_visitor( _literalVisitor, param ) );
    }
	auto result = boost::join( paramVector, ", " );
    return result;
}

string Generator::visitAttrRequiredAndOptionals( AttrRequiredAndOptionals const& astAttrRequiresAndOptionals )
{
	auto result = visitAttrRequiredParams( astAttrRequiresAndOptionals.required ) + "," + visitAttrOptionalParams( astAttrRequiresAndOptionals.optionals );
    return result;
}

string Generator::visitAttrOptionalParams( AttrOptionalParams const& astAttrOptionalParams )
{
    std::vector<string> paramVector;
    for(auto const& param : astAttrOptionalParams )
    {
        paramVector.push_back( param.name.name + "=" + boost::apply_visitor( _literalVisitor, param.value ) );
    }
	auto result = boost::join( paramVector, ", " );
    return result;
}


void Generator::writeOutputFile(string const& source) const
{
//#ifdef _DEBUG
	std::cout << "Generating file: " << _outputFileName << "\n";
//#endif

	std::ofstream out( _outputFileName );
	out << source;
	out.close();
}

bool Generator::outputClass( string const& classSource, string const& fullClassName )
{
	const auto result = true;

	if ( !fs::exists( _outputFileName ) )
	{
		writeOutputFile(classSource);
	}
	else if ( !_isSameSource )
	{
		std::ifstream file( _outputFileName );

		const string originalClassSource( (std::istreambuf_iterator<char>(file)),
			std::istreambuf_iterator<char>());
		file.close();

		_isSameSource = originalClassSource == classSource;

		if ( !_isSameSource)
		{
			try
			{
				fs::remove(_outputFileName);
			}
			catch (boost::filesystem::filesystem_error const & e)
			{
				std::cerr << "exception: " << e.what() << "\n";
			}
			writeOutputFile( classSource );
		}
		else
		{
			
#ifdef _DEBUG
			std::cout << "File not modified: " << _outputFileName << "\n";
#endif			
		}
    }
	else
	{

#ifdef _DEBUG
		std::cout << "File not modified: " << _outputFileName << "\n";
#endif
	}

    return result;
}

string Generator::createFileHeader( string const& fullClassName )
{
    std::ostringstream out;

    out << "/**\n";  
    out << "* PDL compiler generated code\n";
    out << "*      class " << fullClassName << "\n";
//    out << "*      Generated on ";
//
//    auto *facet = new boost::posix_time::time_facet("%d-%b-%Y %H:%M:%S");
//    out.imbue( std::locale( out.getloc(), facet) );
//    out << boost::posix_time::second_clock::local_time() << "\n";

    out << "*/\n";

    return out.str();
}


string Generator::arrayNotationList2Brackets( ArrayNotationList const& arrayNotationList)
{
    string result;
    const auto dimensions = arrayNotationList.size();
    for (unsigned int i = 0; i < dimensions; i++)
    {
        result += "[]";
    }

    return result;
}

string Generator::getPropertyType( PropertyType const& propertyType )
{
	auto result = translateType(SymbolTable::joinIdentifier(propertyType.type));

    result += arrayNotationList2Brackets(propertyType.arrayNotationList);

	return result;

}


string Generator::fieldNameFromPropertyName( string name )
{
	auto result = "_" + name;
    result[ 1 ] = tolower( name[ 0 ] );
    return result;
}

string Generator::singleFieldType( std::string const& fieldType, std::string const& fieldName )
{
    return ( boost::format( _formatter.singleField ) % fieldType % fieldName).str();
}

string Generator::indexerFieldType( std::string const& fieldType, std::string const& fieldName, std::string const& indexerType )
{
    return ( boost::format( _formatter.indexerField ) % fieldType % fieldName % indexerType ).str();
}

string Generator::translateType( std::string const& pdlType )
{
	auto result = pdlType;
    if ( _typesMap.find( pdlType ) != _typesMap.end() )
    {
        result = _typesMap[ pdlType ];
    }
    return result;
}

string Generator::getEmptyValueOfType( std::string const& type )
{
	auto result = _nullObjectValue;
    if ( _emptyValueOfType.find( type ) != _emptyValueOfType.end() )
    {
        result = _emptyValueOfType[ type ];
    }
    return result;
}

void Generator::doMembers( MemberList const& astMember )
{
    for(auto const& member: astMember )
    {
        std::string memberStr = boost::apply_visitor( _memberVisitor, member );
    }
}

string Generator::visitProperty( PropertyNode const& astProperty ) 
{
	auto result = doProperty( astProperty ); 
    _properties.push_back( result );
    return result;
}

string Generator::visitShortProperty(ShortPropertyNode const& shortPropertyNode)
{
	auto result = doShortProperty(shortPropertyNode);
	_properties.push_back(result);
	return result;
}

string Generator::visitMethod( MethodNode const& astMethod )  
{
	auto result = doMethod( astMethod );
    _methods.push_back( result );
    return result;
}

string Generator::visitConst( ConstNode const& constNode )
{
    auto result = doConst( constNode );
    _consts.push_back( result );
    return result;
}
