package main

import (
	"bufio"
	"fmt"
	"os"
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
}