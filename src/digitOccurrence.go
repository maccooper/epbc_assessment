//digitOccurence.go - Contains core logic to count how many times a specific digit appears within a given range of intervals
package main

import (
	"fmt"
	"strconv"
)

func validate_helper(num int64, seriesType int) (bool) {
	//validates seriesType of the form [1,2,3] & handles corresponding odd-even matching
	if seriesType == 1 {
		return true
	} else if (seriesType == 2) && (num%2 == 0) {
		//wanted evens, got evens.
		return true
	} else if (seriesType == 2) && (num%2 != 0) {
		//wanted evens, got odds.
		return false
	} else if (seriesType == 3) && (num%2 == 0) {
		//wanted odds, got evens.
		return false
	} else if (seriesType == 3) && (num%2 != 0) {
		//wanted odds, got odds.
		return true
	}
	return false
}

func validateParams(seriesIncrement int64, specifiedDigit int, seriesType int) error {
	//Validates specifiedDigit is (0,9), seriesType = [1 || 2 || 3]

	if specifiedDigit < 0 || specifiedDigit > 9 {
		return fmt.Errorf("specifiedDigit must be between 0 and 9")
	}
	if seriesType < 1 || seriesType > 3 {
		return fmt.Errorf("seriesType must be 1,2, or 3")
	}
	return nil
}

func DigitOccurrence(seriesStart int64, seriesEnd int64, seriesIncrement int64, specifiedDigit int, seriesType int) (int, error) {
	//Input: range of values, increment number + odd/even modifier.
	//Output: number of char digit occurences for all numbers in series.

	if err := validateParams(seriesIncrement, specifiedDigit, seriesType); err != nil {
		return 0, err
	}

	if seriesIncrement < 0 {
		//Handles edge-case for descending seriesIncrement
		seriesStart, seriesEnd = seriesEnd, seriesStart
		seriesIncrement = -seriesIncrement
	
	}

	var count int = 0;

	for i := seriesStart; i <= seriesEnd; i += seriesIncrement {
		var temp int64 = i
		if(validate_helper(i, seriesType)) {
			//proceed
			numStr := strconv.FormatInt(temp, 10)
			for j := 0; j < len(numStr); j++ {
				if numStr[j] == (byte(specifiedDigit) + '0') {
					//fmt.Printf("%d @ %s\n",numStr[j], numStr)
					count++
				}
			}
		}

	}

	return count, nil
}
