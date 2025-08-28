// array and slices

package main

import (
	"fmt"
)

func learnmap() {
	store := map[string]int{
		"milk":     90,
		"sewaiyan": 100,
	}
	// add new entry
	store["cake"] = 20

	// access
	fmt.Println("Quantity of Cake:", store["cake"])

	// check existence of an element
	quantity, exists := store["oil"]
	if exists {
		fmt.Println("Exists in map and quantity is", quantity)
	} else {
		fmt.Println("Doesn't exist in map")
	}

	// loop over elements of map
	for item, quantity := range store {
		fmt.Println(item, "->", quantity)
	}
}

/*
_ means “I don’t care about this variable.”
*/
func database() {
	users := make(map[string]string) // map[username]password

	// Register function
	register := func(username, password string) {
		if _, exists := users[username]; exists {
			fmt.Println("User already exists:", username)
		} else {
			users[username] = password
			fmt.Println("User registered:", username)
		}
	}

	// Login function
	login := func(username, password string) {
		if storedPass, exists := users[username]; exists {
			if storedPass == password {
				fmt.Println("Login successful for:", username)
			} else {
				fmt.Println("Wrong password for:", username)
			}
		} else {
			fmt.Println("User not found:", username)
		}
	}

	// Test cases
	register("saurabh", "1234")
	register("alice", "abcd")
	register("saurabh", "xyz")
	login("alice", "abcd")
	login("alice", "wrong")
	login("bob", "test")
}

func main() {
	// Array
	var arr [5]int
	arr[0] = 10
	arr[1] = 11

	fmt.Println("Array:", arr)
	fmt.Println("Length:", len(arr))

	// Slice
	nums := []int{1, 2, 3}
	fmt.Println("Slice nums:", nums)

	// Append values to slice
	nums = append(nums, 4, 5)
	fmt.Println("Slice after append:", nums)

	// Slicing technique
	fmt.Println("First 3 elements:", nums[:3])
	fmt.Println("Last 2 elements:", nums[3:])

	learnmap()
	database()
}
