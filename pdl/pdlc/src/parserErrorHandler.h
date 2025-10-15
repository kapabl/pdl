#pragma once

namespace io { namespace pdl
{
    template <typename Iterator>
    struct ErrorHandler
    {

        const char* _fileName;

        template <typename, typename, typename>
        struct result { typedef void type; };

        ErrorHandler(const char* fileName, Iterator first, Iterator last) : 
            first( first ), 
            last( last ),
            _fileName( fileName )
        {}

        template <typename Message, typename What>
        void operator()(
            Message const& message,
            What const& what,
            Iterator errorPosition) const
        {
            int line;
            Iterator lineStart = getErrorPosition( errorPosition, line );
            if ( errorPosition != last )
            {
                auto lineOffset = errorPosition - lineStart + 1;

                auto visualStudioError =  boost::format( std::string( "%1%(%2%,%3%): error: %4% %5%" ) ) 
                    % std::string( _fileName )
                    //% boost::lexical_cast<std::string>(line)
                    % std::to_string( line )
                    //% boost::lexical_cast<std::string>(lineOffset)
                    % std::to_string( lineOffset )
                    % message
                    % what;

                //std::cout << visuaStudioError.str() << std::endl;
                std::cerr << visualStudioError.str() << std::endl;
                /*
                std::cout << message << what << " line " << line << ':' << std::endl;
                std::cout << get_line(lineStart) << std::endl;
                for (; lineStart != errorPosition; ++lineStart)
                    std::cout << ' ';
                std::cout << '^' << std::endl;
                */
            }
            else
            {
                auto visuaStudioError =  boost::format( std::string( "%1%(%2%,%3%): error: unexpected end of file: %4% %5%"  ) )
                    % _fileName
                    % line
                    % "0"
                    % message
                    % what;

                //std::cout << visuaStudioError.str() << std::endl;
                std::cerr << visuaStudioError.str() << std::endl;

                /*
                std::cout << "Unexpected end of file. ";
                std::cout << message << what << " line " << line << std::endl;
                */
            }
        }

        Iterator getErrorPosition( Iterator errorPosition, int& line ) const
        {
            line = 1;
            Iterator i = first;
            Iterator lineStart = first;

            while (i != errorPosition )
            {
                bool eol = false;

                if ( i != errorPosition && *i == '\r' ) // CR
                {
                    eol = true;
                    lineStart = ++i;
                }

                if ( i != errorPosition && *i == '\n') // LF
                {
                    eol = true;
                    lineStart = ++i;
                }

                if ( eol )
                {
                    ++line;
                }
                else
                {
                    ++i;
                }
            }
            return lineStart;
        }

        std::string get_line(Iterator errorPosition) const
        {
            Iterator i = errorPosition;
            // position i to the next EOL
            while (i != last && (*i != '\r' && *i != '\n'))
                ++i;
            return std::string(errorPosition, i);
        }

        Iterator first;
        Iterator last;
        std::vector<Iterator> iters;
    };

}}

