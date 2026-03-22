package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://localhost:5173/learn?lang=go&course=golang"

func main() {
	fmt.Println("Here we learn about URL in golang")
	fmt.Println("URL is:", myurl)

	// parsing the url
	result, _ := url.Parse(myurl)
	fmt.Println("Scheme is:", result.Scheme)
	fmt.Println("Host is:", result.Host)
	fmt.Println("Path is:", result.Path)
	fmt.Println("Port is:",result.Port())
	fmt.Println("Raw query is:", result.RawQuery)

	qparams:=result.Query()
	fmt.Printf("The type of query params is %T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _,val:=range qparams{
		fmt.Println("Param is:",val)
	}

	partsOfUrl:=&url.URL{
		Scheme:"https",
		Host:"localhost:5173",
		Path:"/tutorial",
		RawPath: "user=sagar",
	}
	anotherUrl:=partsOfUrl.String()
	fmt.Println("Another URL is:",anotherUrl)



}
