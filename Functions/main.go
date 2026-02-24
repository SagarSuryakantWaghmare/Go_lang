package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to functions in golang")
	greeterTwo()

	result:=adder(2,5)
	fmt.Println("Addition of this numbers:",result)
	// multiple values
	mulAdder:=proAdder(1,3,4,2,2,2,42,14,2,2)
	fmt.Println("Multiples values adding :",mulAdder)
}
func adder(valOne int,valTwo int) int{
 return valOne+valTwo 
}
// Adding numbers of values
func proAdder(values...int) int{
	total:=0

	for _,val:=range values{
		total+=val
	}
	return total
}
func greeterTwo(){
		fmt.Println("Another method")
	}

func greeter(){
	fmt.Println("Hello Sam")
}