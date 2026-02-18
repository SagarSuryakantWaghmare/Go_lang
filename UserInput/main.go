package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	welcome:="Welcome to user input"
	fmt.Println(welcome)
	// There is package for that input: bufio
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of your skills out of 5:")

	//comma ok || error err 
	// In go lang there is no try catch
	input,_:=reader.ReadString('\n')
	fmt.Println("Thanks for rating,",input)
	fmt.Printf("Type of this rating is %T",input)

	// Time how to use it
	fmt.Println("Here we see how to use time")
	presentTime:=time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
    createdDate:=time.Date(2020,time.August,19,3,4,2,5,time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday")) 
	//for building we directly write go build
	// There is like different go os
	//GOOS='darwin' =mac and then go build 
	//GOOS='windows' go build
	//GOOS='linux' go build
	// in this way we get the .exe file to execute
	
	//memory mangement
	//new() and make()
	// new() allocate you will get a memory address and zero storage
	// make() you will get a memory address and non-zero storage
	// Garbage collections happens automatically




}