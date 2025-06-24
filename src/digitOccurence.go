package main

import (
	"fmt"
	"strconv"
)

func validate_helper(num int64, seriesType int) (bool) {
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

	if specifiedDigit < 0 || specifiedDigit > 9 {
		return fmt.Errorf("specifiedDigit must be between 0 and 9")
	}
	if seriesType < 1 || seriesType > 3 {
		return fmt.Errorf("seriesType must be 1,2, or 3")
	}
	return nil
}

func digitOccurrence(seriesStart int64, seriesEnd int64, seriesIncrement int64, specifiedDigit int, seriesType int) (int, error) {
	//Input: range of values, increment number + odd/even modifier.
	//Output: number of char digit occurences for all numbers in series.

	if err := validateParams(seriesIncrement, specifiedDigit, seriesType); err != nil {
		return 0, err
	}

	if seriesIncrement < 0 {
		//Swap start and end
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
					fmt.Printf("%d @ %s\n",numStr[j], numStr)
					count++
				}
			}
		}

	}

	return count, nil
}


func main() {
	seriesStart := int64(0)
	seriesEnd := int64(100)
	seriesIncrement := int64(1)
	specifiedDigit := int(0)
	seriesType := int(1)

	result, err := digitOccurrence(seriesStart, seriesEnd, seriesIncrement, specifiedDigit, seriesType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Number of occurrences of digit %d between %d and %d: %d\n",
		specifiedDigit, seriesStart, seriesEnd, result)
}
