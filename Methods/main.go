package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods learning")
	sagar:=User{"Sagar","sagar@gmail.com",true,23}
	fmt.Printf("Sagar details are:%v\n",sagar)
    sagar.GetStatus()
	sagar.NewMail()
}

// the name should be First letter Capital
type User struct{
	Name string
	Email string
	Status bool 
	Age int
}
func (u User) GetStatus(){
	fmt.Println("Is User active:",u.Status)
}

func (u User) NewMail(){
	u.Email="test@go.dev"
	fmt.Println("Email of this use is:",u.Email)
}
