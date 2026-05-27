package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	//integers
	var num1 int8 = 127
	var num2 uint8 = 255

	//decimals
	var num3 float32 = 32.5678999
	var num4 float64 = 32.5678999

	//type casting
	var num5 int32 = 10
	var num6 float32 = 10.1
	var result = float32(num5) + num6

	//strings
	var str1 string = "Hello \nWorld"
	var str2 string = `Hello 
	World`

	var str3 string = "Hello"
	var str4 string = "World"
	var strResult string = str3 + " " + str4

	//lengths of strings
	var str string = "hello world"
	var byteLength int = len(str)
	var runeLength int = utf8.RuneCountInString(str)

	//runes
	var rune1 rune = 'a' //prints the unicode code point
	var rune2 rune = 'b'

	//booleans
	var myBoolean bool = true

	//default values
	var intDefault int
	var int8Default int8
	var int16Default int16
	var int32Default int32
	var int64Default int64

	var uintDefault uint
	var uint8Default uint8
	var uint16Default uint16
	var uint32Default uint32
	var uint64Default uint64

	var float32Default float32
	var float64Default float64

	var runeDefault rune

	var strDefault string

	var boolDefault bool

	//shorthand

	//without type
	var myVar1 = "text" //type of variable is inferred here

	//without `var` keyword
	myVar2 := 1

	//constants
	const pi float32 = 3.14

	//will throw error
	// pi = 3.1

	//calling a function

	var divResult, divRemainder, err = intDivision(20, 5)

	if err != nil {
		fmt.Printf("%s", err.Error())
	} else if divRemainder == 0 {
		fmt.Printf("The result of the integer division is %v \n", divResult)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder is %v \n", divResult, divRemainder)
	}

	//print
	fmt.Println("Hello World!")

	fmt.Println("Integers:", num1, num2)
	fmt.Println("Decimals:", num3, num4, num5, num6, result)
	fmt.Println("Strings:", str1, str2, str3, strResult, str)
	fmt.Println("Byte Length:", byteLength)
	fmt.Println("Unicode Code points:", runeLength)
	fmt.Println("Characters Unicode:", rune1, rune2)
	fmt.Printf("Characters:%c %c \n", rune1, rune2)
	fmt.Println("Boolean:", myBoolean)

	fmt.Println("Integer defaults:", intDefault, int8Default, int16Default, int32Default, int64Default)
	fmt.Println("Unsigned Integer defaults:", uintDefault, uint8Default, uint16Default, uint32Default, uint64Default)
	fmt.Println("Float defaults:", float32Default, float64Default)
	fmt.Println("Rune default:", runeDefault)
	fmt.Println("String default:", strDefault)
	fmt.Println("Boolean default:", boolDefault)

	fmt.Println("Shorthand:", myVar1, myVar2)

	fmt.Println("Constant:", pi)
}

func intDivision(numerator int, denominator int) (int, int, error) {

	var err error //default value is nil

	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}

	result := numerator / denominator
	remainder := numerator % denominator

	return result, remainder, err
}
