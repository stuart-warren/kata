package fizzbuzz

import "strconv"

// Do takes a Number
// if number devides by 3, output fizz
// if number divides by 5, output buzz
// if number divides by 3 & 5, ouput fizzbuzz
// else output the number
func Do(input int) (string, error) {
	return strconv.Itoa(input), nil
}
