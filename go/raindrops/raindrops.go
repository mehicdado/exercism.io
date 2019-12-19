package raindrops

import "strconv"

//Convert function returns string based on given integer devision reminder when deviding with 3, 5, or 7
//eg. if factor is 3 then ouput is "Pling", if 3 and 5 then "PlingPlang", for 7 Plong
func Convert(number int) string {
	var result string

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
