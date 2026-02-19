package main

import "fmt"

func main() {
	fmt.Println("Now here we learn the array:")

	var fruitlist[4] string
	fruitlist[0]="Apple"
	fruitlist[1]="Banana"
	fruitlist[3]="Mango"

	fmt.Println("Fruit list is:",fruitlist)
	fmt.Println("Fruit list lenght is:",len(fruitlist))

	var vegList=[3] string{"potato","beans","mushroom"}
	fmt.Println("Vegy list is :",vegList)

}