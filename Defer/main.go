package main

import "fmt"

func main() {
	// Defer exectues in the manner of last in first out
	defer fmt.Println("Sagar Bhai")
	defer fmt.Println("Atharva bhai")
	fmt.Println("Defer in go lang")
	myDefer()
}

func myDefer(){
	for i:=0; i<5;i++{
		defer fmt.Print(i)
	}
}