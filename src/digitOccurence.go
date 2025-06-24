package main

import (
	"fmt"
)

func validate_helper(num int, seriesType int) (bool) {
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
	} else {
		//needs error handling for seriesType != [1,2,3]
		return false
	}
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

func digitOccurence(seriesStart int64, seriesEnd int64, seriesIncrement int64, specifiedDigit int, seriesType int) (int, error) {
	if err := validateParams(seriesIncrement, specifiedDigit, seriesType); err != nil {
		return 0, err
	}

	if(validate_helper(1, 2)) {
		return 1, nil
	}
	return 0, nil
}


func main() {
	fmt.Println(digitOccurence(1, 11, 1, 1, 1))
}
