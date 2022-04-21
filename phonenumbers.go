package phonenumbers

import (
	"fmt"
	"strconv"
	"strings"
)

//Splits a range of numbers and returns two ints
//example "02080668800 - 02080668850"
func SplitRange(input string, deliminator string) (int, int) {

	numbers := strings.Split(input, deliminator)

	return ConverttoE164int(numbers[0]), ConverttoE164int(numbers[1])

}

func ListRange(firstnum int, secondnum int) {

	var divideby int
	var output string

	if secondnum < 9999 {

		fmt.Println("Second number is shorthand")

		if DigitCount(secondnum) == 2 {
			divideby = 100
		} else if DigitCount(secondnum) == 3 {
			divideby = 1000
		} else if DigitCount(secondnum) == 4 {
			divideby = 10000
		}

		if divideby == 0 {
			return
		}

		base := firstnum / divideby
		fmt.Println("Base number", base)

		startingdigits := RetrieveTrailingDigits(firstnum, DigitCount(secondnum))
		fmt.Println("Starting Number", startingdigits)

		for count := ConverttoE164int(startingdigits); count <= secondnum; count++ {
			output = fmt.Sprintf("%v%v", base, count)
			fmt.Println(output)
		}

	} else {
		fmt.Println("Second number is a full number")

		for count := firstnum; count <= secondnum; count++ {
			fmt.Println(count)
		}

	}
}

//Retrieves the trailing digits of a number in string format
func RetrieveTrailingDigits(number int, decimalplace int) string {

	numberstring := fmt.Sprintf("%v", number)

	trailingdigits := numberstring[len(numberstring)-decimalplace:]

	return trailingdigits

}

//Returns count for decimal places of a number
func DigitCount(number int) int {

	count := 0

	for number != 0 {
		number /= 10
		count += 1
	}

	return count

}

//Convert telephone number in string format to int in e164
//expected format "02080668800" or "200" for when number ranges are written like "02080668800 - 8850"
func ConverttoE164int(number string) int {

	number = strings.TrimSpace(number)

	if len(number) < 5 {

		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		return n
	}

	number = ConverttoE164(number)

	n, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}

	return n

}

//Takes a number in string format, converts it to e164 format.
//02080668866 to 442080668866
func ConverttoE164(number string) string {
	number = number[1:]
	number = "44" + number

	return number
}
