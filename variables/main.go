package main

import "fmt"

const LoginToken string="sagar" //public
func main() {
	fmt.Println("Sagar will be shown everyone , why he known as the unknown ultima")
	var username string = "Sagar"
	fmt.Println(username)
	fmt.Printf("Variables is of type:%T \n", username)

	var isLoggedIn bool=true;
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type:%T \n",isLoggedIn);

	var smallVal uint8=255;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type :%T \n",smallVal);

	var smallFloat float32=238723.2;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n",smallFloat);

	//Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable);
	fmt.Printf("Variable is of type :%T \n",anotherVariable);

	// implicit type
	var website ="sagarwaghmare.online"
	fmt.Println(website);
	

	//no var style
	numberOfUser:=300000
    fmt.Println(numberOfUser);

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of Type :%T \n",LoginToken)
}
