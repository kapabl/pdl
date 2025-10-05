const fs = require( 'fs-extra' );
const path = require( 'path' );
const rimraf = require( 'rimraf' );

/**
 *
 * @param dir
 */
function cleanDir( dir ) {
    if ( fs.existsSync( dir ) )
    {
        var oldIndex = 1;
        var old = dir + '_old';
        while ( fs.existsSync( old ) )
        {
            try
            {
                rimraf.sync( old );
            }
            catch ( error )
            {
            }
            if ( fs.existsSync( old ) )
            {
                old = dir + '_old' + oldIndex.toString();
                oldIndex++;
            }
        }

        fs.renameSync( dir, old );
        rimraf.sync( old );
    }

    fs.ensureDirSync( dir );
}

/**
 *
 * @param message
 */
function outputError( message ) {
    process.stderr.write( message.red );
    throw Error( message );
}

/**
 *
 * @param {string} namespace
 * @return {string}
 */
function namespace2dir( namespace ) {
    var result = namespace.split( '.' ).join( path.sep );
    return result;
}

/**
 *
 * @param {string} namespace
 * @return {string}
 */
function namespace2php( namespace ) {
    var result = namespace.split( '.' ).map( word => {
        const result = word.charAt( 0 ).toUpperCase() + word.slice( 1 );
        return result;
    } );
    result = result.join( '\\' );
    return result;
}

/**
 *
 * @returns {string|*}
 */
function getTemplatesDir() {

    var result = 'default';

    for ( var i = 0; i < arguments.length; ++i )
    {
        if ( arguments[ i ] )
        {
            result = arguments[ i ];
            break;
        }

    }

    if ( result === 'default' )
    {
        result = path.join( __dirname, 'templates' );
    }

    return result;
}

var doVerbose = false;

/**
 *
 * @param value
 */
function verboseLog( value ) {
    if ( doVerbose )
    {
        process.stdout.write( value.yellow + '\n' );
    }
}

module.exports = {
    cleanDir: cleanDir,
    outputError: outputError,
    namespace2dir: namespace2dir,
    verboseLog: verboseLog,
    getTemplatesDir: getTemplatesDir,
    namespace2php: namespace2php,
    setVerbose: function( verbose ) {
        doVerbose = verbose;
    }
};
