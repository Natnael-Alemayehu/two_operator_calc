package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(x int, y int) (int, error)  { return x + y, nil }
func sub(x int, y int) (int, error)  { return x - y, nil }
func mult(x int, y int) (int, error) { return x * y, nil }
func div(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return x / y, nil
}

var myop = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mult,
	"/": div,
}

func main() {

	expressions := [][]string{
		{"4", "+", "2"},
		{"4", "-", "2"},
		{"4", "*", "2"},
		{"4", "/", "0"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", expression)
			continue
		}

		operator_1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		operator := expression[1]
		opFunc, ok := myop[operator]
		if !ok {
			fmt.Println("Unsupported operator: ", operator)
			continue
		}
		operator_2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, _ := opFunc(operator_1, operator_2)
		fmt.Println(result)
	}

}
