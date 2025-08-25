package main

import (
	"fmt"
)

func main() {
	age := 18

	// if-else example
	if age >= 18 {
		fmt.Println("You are eligible to purchase liquor")
	} else {
		fmt.Println("You are not eligible to purchase liquor")
	}

	// declaring var inside if statement
	if score := 85; score >= 40 {
		fmt.Println("You passed the examination")
	} else {
		fmt.Println("Ahh! you missed the criteria to pass the examination")
	}

	// switch example
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("Start of the week")
	case "friday":
		fmt.Println("Mid week")
	case "saturday":
		fmt.Println("Weekend")
	default: 
		fmt.Println("Sunday")
	}

	// traditional forloop
	for i:=0;i<=10;i++{
		fmt.Println(i);
	}
	// While-style loop
	x:=1
	for x<=3{
		fmt.Println(x)
		x++
	}
}
