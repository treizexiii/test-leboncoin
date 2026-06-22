package services

import "fmt"

// Take 2 multiples num1 & num2
// Limite loop to limit
// Replace num1 with str1 and num2 with str2
// return an array of strings with the result
func FizzBuzz(num1 int, num2 int, limit int, str1 string, str2 string) []string {
	resut := make([]string, 0, limit)

	for i := 1; i <= limit; i++ {
		if i%num1 == 0 && i%num2 == 0 {
			resut = append(resut, str1+str2)
		} else if i%num1 == 0 {
			resut = append(resut, str1)
		} else if i%num2 == 0 {
			resut = append(resut, str2)
		} else {
			resut = append(resut, fmt.Sprintf("%d", i))
		}
	}

	return resut
}
