package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	add := flag.Bool("add", false, "Add two numbers")
	sub := flag.Bool("subtract", false, "Subtract two numbers")
	mult := flag.Bool("multiply", false, "Multiply two numbers")
	div := flag.Bool("divide", false, "Divide two numbers")

	flag.Parse()
	fmt.Println("Welcome to the simple calculator!")
	fmt.Println(time.Now().Format(time.ANSIC))

	firstValue := getInput()
	secondValue := getInput()

	switch {
	case *add:
		fmt.Printf("Additon: %.2f \n", addition(firstValue, secondValue))
	case *sub:
		fmt.Printf("Difference: %.2f \n", subtraction(firstValue, secondValue))
	case *mult:
		fmt.Printf("Product: %.2f \n", multiplication(firstValue, secondValue))
	case *div:
		fmt.Printf("Division: %.2f \n", division(firstValue, secondValue))
	default:
		fmt.Fprintln(os.Stderr, "Please use one of the following flags when calling the application -add, -subtract, -multiply or -divide")
		os.Exit(1)
	}
}

func getInput() float64 {
	fmt.Print("Please insert your value: ")
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		println(err.Error())
		panic("Invalid input, must be a number.")
	}
	return value
}

func addition(firstValue float64, secondValue float64) float64 {
	sum := firstValue + secondValue
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n\n", firstValue, secondValue, sum)
	return sum
}

func subtraction(firstValue float64, secondValue float64) float64 {
	sum := firstValue - secondValue
	sum = math.Round(sum*100) / 100
	fmt.Printf("The subtraction of %.2f and %.2f is %.2f\n\n", firstValue, secondValue, sum)
	return sum
}

func multiplication(firstValue float64, secondValue float64) float64 {
	sum := firstValue * secondValue
	sum = math.Round(sum*100) / 100
	fmt.Printf("The multiplication of %.2f and %.2f is %.2f\n\n", firstValue, secondValue, sum)
	return sum
}

func division(firstValue float64, secondValue float64) float64 {
	sum := firstValue / secondValue
	sum = math.Round(sum*100) / 100
	fmt.Printf("The division of %.2f and %.2f is %.2f\n\n", firstValue, secondValue, sum)
	return sum
}
