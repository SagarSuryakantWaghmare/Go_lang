package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)
func main(){
	fmt.Println("Here learn the file handing")
	content:="This needs to go in a file -text.file"

	file,err:=os.Create("./mytextfile.txt");

	if err!=nil{
		panic(err)
	}

	length,err:=io.WriteString(file,content)

	if(err!=nil){
		fmt.Println(err)
	}
	fmt.Println("Length is:",length)
	defer file.Close()
	readFile("./mytextfile.txt")

}

func readFile(filename string){
	datebyte,err:=ioutil.ReadFile(filename)
	if err!=nil{
		panic(err)
	}
	fmt.Println("Text data inside the file is \n",string(datebyte))
}