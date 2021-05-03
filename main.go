package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	on := true
	for on {
		fmt.Println("")
		switch task := GetInput("Write task number or 0 to exit : "); strings.ToLower(task) {
		case "1":
			fmt.Println("Reverse")
			fmt.Println(strings.Join(Reverse(strings.Split(GetInput("Enter : "), "")), ""))
		case "2":
			fmt.Println("Palindrome")
			fmt.Println(IsPalindrome(GetInput("Enter : ")))
		case "3":
			fmt.Println("Multiplication Table")
			fmt.Printf(GetMultiplicationTable(GetInput("Enter : ")))
		case "4":
			fmt.Println("Max Min")
			fmt.Printf(GetMaxMin(GetInput("String format -> 1,2,3... : ")))
		case "5":
			fmt.Println("Largest of 3")
			input := GetInput("String format -> 1,2,3... : ")
			if len(strings.Split(input, "")) <= 3 {
				fmt.Printf("The largest is %s", GetMax(input))
			} else {
				fmt.Println("Write 3 or less numbers")
			}
		case "6":
			fmt.Println("Calculate sum of number")
			fmt.Println(GetSumOfNumber(GetInput("Enter : ")))
		case "7":
			fmt.Println("Pyramid")
			fmt.Println(GetPyramid(GetInput("Enter : ")))
		case "8":
			fmt.Println("Odd or Even")
			fmt.Println(GetOddOrEven(GetInput("Enter : ")))
		case "0":
			on = false
		}
	}
}

func GetInput(sentence string) string {
	var input string
	fmt.Print(sentence)
	fmt.Scanln(&input)
	return input
}

func Reverse(input []string) []string {
	if len(input) <= 1 {
		return input
	}
	return append(Reverse(input[1:]), input[0])
}

func IsPalindrome(input string) bool {
	reversed := strings.Join(Reverse(strings.Split(input, "")), "")
	if strings.ToLower(input) == strings.ToLower(reversed) {
		return true
	}
	return false
}

func GetMultiplicationTable(input string) string {
	var result string
	value, _ := strconv.Atoi(input)
	for tick := 1; tick < 10; tick++ {
		result += fmt.Sprintf("%d X %d = %d\n", value, tick, value*tick)
	}
	return result
}

func GetMax(input string) string {
	array := strings.Split(input, ",")
	max, _ := strconv.Atoi(array[0])
	for _, str := range array {
		i, _ := strconv.Atoi(str)
		if i > max {
			max = i
		}
	}
	return fmt.Sprintf(strconv.Itoa(max))
}

func GetMin(input string) string {
	array := strings.Split(input, ",")
	min, _ := strconv.Atoi(array[0])
	for _, str := range array {
		i, _ := strconv.Atoi(str)
		if i < min {
			min = i
		}
	}
	return fmt.Sprintf(strconv.Itoa(min))
}

func GetMaxMin(input string) string {
	var result string
	result = fmt.Sprintf("Min value is %s\nMax value is %s\n", GetMin(input), GetMax(input))
	return result
}

func GetSumOfNumber(input string) string {
	var result int
	array := strings.Split(input, "")
	for _, str := range array {
		i, _ := strconv.Atoi(str)
		result += i
	}
	return fmt.Sprintf(strconv.Itoa(result))
}

func GetOddOrEven(input string) string {
	i, _ := strconv.Atoi(input)
	if i%2 == 0 {
		return "This is Even number"
	}
	return "This is Odd number"
}

func GetPyramid(input string) string {
	var result string
	i, _ := strconv.Atoi(input)
	asterisk := "*"
	empty := " "
	for tick := 1; tick < i+1; tick++ {
		if tick%2 != 0 {
			result += fmt.Sprintln(strings.Repeat(empty, i-tick/2), strings.Repeat(asterisk, tick), strings.Repeat(empty, i-tick/2))
		}
	}
	return result
}
