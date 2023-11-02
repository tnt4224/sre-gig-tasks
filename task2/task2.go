package main

// Task 2 (Async Review)

import ( 
	"fmt"
	"math/rand"
)

func main() {

	//
	// 1. Generate a random number between 1 and 100
	//   i. If the number is higher than 50 print "It's closer to 100"
        //  ii. If the number is lower than 50 print "It's closer to 0"
        // iii. Print the generated random number
	//
	// 2. Modify the previous code to print "It's 50!" if the random number is 50

	// NOTE: 
	// rand.Seed has been deprecated since v 1.20 
	// https://tip.golang.org/doc/go1.20
	//
	//rand.Seed(time.Now().UnixNano())
	var number = genRandomInt(1, 100)

	if number > 50 {
		fmt.Println("It's closer to 100")
	} else if number < 50 {
		fmt.Println("It's closer to 0")
	} else {
		fmt.Println("It's 50!")
	}

}

func genRandomInt(min int, max int) int {
	return min + rand.Intn(max - min)
}
