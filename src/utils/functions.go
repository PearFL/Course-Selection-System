package utils

import (
	"fmt"
	"unicode"
)

// 测试函数

func helloFunc() {
	fmt.Println("Hello World!")
}

/*
@title	StrIsLetter
@description	判断string字符串是否均由字符组成
@auth	马信宏	时间（2022/2/9   16:25 ）
@param	str	string
@return	无	bool
*/

func StrIsLetter(str string) bool {
	for _, val := range str {
		if unicode.IsLetter(val) == false {
			return false
		}
	}
	return true
}

/*
@title	StrIsLowerLetterAndUpperLetterAndNumber
@description	判断string字符串是否同时由字符或者数字组成
@auth	马信宏	时间（2022/2/9   16:25 ）
@param	str	string
@return	无	bool
*/

func StrIsLowerLetterAndUpperLetterAndNumber(str string) bool {
	isUpper, isLower, isNumber := false, false, false
	for _, val := range str {
		if unicode.IsUpper(val) {
			isUpper = true
		} else if unicode.IsLower(val) {
			isLower = true
		} else if unicode.IsNumber(val) {
			isNumber = true
		} else {
			// 出现不是字母和数字的字符
			return false
		}
	}
	if isUpper == true && isLower == true && isNumber == true {
		return true
	} else {
		return false
	}
}
