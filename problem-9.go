package main

import (
	"fmt"
)

func numberToWord(number int) string {
	if number < 10 || number >= 1000 {
		return "Error: Number is out of the valid range (10-999)."
	}

	units := []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	teens := []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	if number < 20 {
		return teens[number-10]
	} else if number < 100 {
		return tens[number/10] + " " + units[number%10]
	} else {
		hundredPart := units[number/100] + " hundred"
		remainder := number % 100

		if remainder == 0 {
			return hundredPart
		} else if remainder < 10 {
			return hundredPart + " and " + units[remainder]
		} else if remainder < 20 {
			return hundredPart + " and " + teens[remainder-10]
		} else {
			return hundredPart + " and " + tens[remainder/10] + " " + units[remainder%10]
		}
	}
}

func main() {
	var number int

	fmt.Print("Enter a number (10-999): ")
	fmt.Scanln(&number)

	result := numberToWord(number)
	fmt.Println(result)
}
