package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func RLERevert(str string) (string, error) {
	data := []rune(str)
	result := make([]rune, 0)
	if unicode.IsDigit(data[0]) {
		return "", fmt.Errorf("invalid string")
	}
	i := 0
	for i < len(data) {
		if unicode.IsDigit(data[i]) {
			number := make([]rune, 0)
			for i < len(data) && unicode.IsDigit(data[i]) {
				number = append(number, data[i])
				i++
			}
			i--
			n, _ := strconv.Atoi(string(number))
			for j := 0; j < n-1; j++ {
				result = append(result, result[len(result)-1])
			}
		} else {
			result = append(result, data[i])
		}
		i++
	}
	return string(result), nil
}

func main() {
	str, err := RLERevert("a4bc2d5e3")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
