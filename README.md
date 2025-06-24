q1) given a set of tables, design a rest api which supports CRUD operations + provide all rest endpoints that would be necessary for CRUD and the varius objects.

q2) write a function to count the number of occurences of specified digit in a series.
        Parameters: - SeriesStart:  beginning # of series [-(2^53-1), (2^53 - 1)]
                    - SeriesEnd:    end # of series.    [-(2^53-1), (2^53 - 1)]
                    - SeriesIncr:   determines incr between elements of the series
                    - SpecifiedDigit:   target digit to track through series
                    - SeriesType:   identifier for which items in the series to analyze
                                    * 1: analyze all elements in the series
                                    * 2: analyze only even numbered elements
                                    * 3: "" odd elements.


    ex: fn digit_counter(1, 11, 1, 1, 1) => 4
            //(Represented by the tally of elements [1,10,11]


# TODO:
    
    - Complete restful api implementation for q1 (api\_endpoints.md)
    - Unit tests? 
    - convert digi function to full restful api
    - write SQL queuries for q4 
    - Write deployment instructions
    - Clean README

