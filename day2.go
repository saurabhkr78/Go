package main

import(
	"fmt"
)
//defining constants
const pi=1.314
const Greeting="heloow everyone"
const year=2025
func main(){
	//explicitly define a var
	var name string="saurabh"
	//type inference: auto guess the type by go
	var age=25

	//short declaration (only inside fucntion)

	city:="Bengaluru"


	var loved bool=false
	var money int=123
	var height float64=5.9
	var study string="CU"




	fmt.Println(loved,money,height,study,year)


	fmt.Println(name,age,city,year)
	fmt.Println("pi:",pi)
	fmt.Println(Greeting)
}




