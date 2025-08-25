package main

import "fmt"

// Function without return
func greet() {
	fmt.Println("Hello from Go Bootcamp!")
}

// Function with parameters (no return)
func add(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// Function with return type
func multiply(a int, b int) int {
	return a * b
}

// Function with multiple return values
func divide(a int, b int) (int, int) {
	quo := a / b
	rem := a % b
	return quo, rem
}
// Function with multiple return values with name(important to remember this)
func divide1(a int, b int) (quo1 int, rem1 int) {
	quo1= a / b
	rem1 = a % b
	return;
}

func maxi(a int,b int)bool{
	if a>b{
		return true
	}
	return false
}
func judge(a int ) bool{
	if a%2==0{
		return true
	}
	return false
}

func main() {
	greet()

	add(3, 7)

	res := multiply(3, 7)
	fmt.Println("Product of two numbers is:", res)

	q, r := divide(20, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	q1, r1 := divide1(20, 3)
	fmt.Println("Quotient:", q1, "Remainder:", r1)

	result:=maxi(5,6)
	fmt.Println(result)

	ans:=judge(6)
	fmt.Println(ans)
}
