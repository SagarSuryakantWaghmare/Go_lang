package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages :=make(map[string]string)
	languages["JS"]="JavaScript"
	languages["PY"]="Python"
	languages["RB"]="Ruby"
	fmt.Println(languages)
	fmt.Println("JS stands for",languages["JS"])

	delete(languages,"RB")
	fmt.Println("Updated languages map:", languages)

    // Loops are used to iterate over the map
	for key,value:=range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
}