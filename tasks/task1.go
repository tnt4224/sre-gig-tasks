package main

import ( 
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//
	// 1. Define the following array "Menu"}
	//   i. Append to it the following item: "hamburger"
	//  ii. Append to it the following item: "salad"
	// iii. Iterate over the list and print for each item 
	//      Food: <Food name>. Make sure to replace <Food name> with item from the array
	//
	Menu := []string{}
	Menu = append(Menu, "hamburger")
	Menu = append(Menu, "salad")
	printMenu(Menu)

	//
        // 2. Define an array of 5 items
        //   i. Iterate over it and print for each item the 
	//      following: This is <ITEM> and its index in 
	//      the array is <INDEX>
        //
	//var arraySize = 5
	var s [5]float64
	rand.Seed(time.Now().UnixNano())

	// create array of 5 random floats
	for i :=0; i < 5; i++ {
		s[i] = rand.Float64()
	}

	// iterrate over the array and print out in 
	// format specificed in 2.i
	for index, item := range s {
		fmt.Println("This is ",item," and its index in the array is ",index)
	}

}

func printMenu(menu []string){
	for i := 0; i < len(menu); i++ {
		fmt.Println("Food:",menu[i])
	}
}
