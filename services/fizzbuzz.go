package services

import "fmt"

type FizzBuzzRequest struct {
	Num1  int    `json:"num1" binding:"required"`
	Num2  int    `json:"num2" binding:"required"`
	Limit int    `json:"limit" binding:"required"`
	Str1  string `json:"str1" binding:"required"`
	Str2  string `json:"str2" binding:"required"`
}

func (r *FizzBuzzRequest) Validate() bool {
	if r.Num1 <= 0 {
		return false
	}
	if r.Num2 <= 0 {
		return false
	}
	if r.Limit <= 0 {
		return false
	}
	if r.Str1 == "" {
		return false
	}
	if r.Str2 == "" {
		return false
	}
	return true
}

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
