/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 14:34:35
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 15:58:17
 */

package tools

// Abs returns absolute value of number:int
func Abs(number int) int {
	if number > 0 {
		return number
	}
	return -number
}

// Max returns maximum of given numbers
func Max(numbers ...int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Min returns minimum of given numbers
func Min(numbers ...int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

// CheckDuplicates checks for duplicates in sorted array
func CheckDuplicates(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i-1] == array[i] {
			return false
		}
	}
	return true
}

// CheckNumber returns true if given string is a valid number.
// Otherwise false
func CheckNumber(str string) bool {
	minus := false
	if len(str) == 0 {
		return false
	}
	if str[0] == '-' && len(str) == 1 {
		return false
	}
	i := 0
	for i < len(str) && str[i] == '-' {
		if minus == true && str[i] == '-' {
			return false
		}
		if str[i] == '-' {
			minus = true
		}
	}
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		i++
	}
	if i != len(str) {
		return false
	}
	return true
}
