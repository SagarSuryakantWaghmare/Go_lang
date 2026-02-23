package main

import "fmt"

func main(){
	fmt.Println("Structs in golang")
	//There is no inheritance in golang; No super or parent
	Sagar:=User{"Sagar","sagar@go.dev",true,16}
	fmt.Println(Sagar)
	fmt.Printf("Sagar details are:%+v\n",Sagar)
	fmt.Printf("Name is %v and email is %v\n",Sagar.Name,Sagar.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

