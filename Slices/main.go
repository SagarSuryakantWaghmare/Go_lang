package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Here see how to use the slices in go lang")

	var fruitlist=[] string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitlist is %T\n",fruitlist)

	fruitlist=append(fruitlist, "Mango","Banana")
	fruitlist=append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highScores:=make([]int,4)

	highScores[0]=234
	highScores[1]=945
	highScores[2]=375
	highScores[3]=859

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}