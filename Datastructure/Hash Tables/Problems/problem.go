package main

import "fmt"

// Given an array = [2, 1, 3, 2 ,3 ,4, 5, 6] find first reoccuring character.
/*
So for instance
array = [2, 1, 3, 2 ,3 ,4, 5, 6]
output = 2

array = [2,1,1,2,3,5,1,2,4]
output = 1

array = [2.3.4.5]
output = -1
*/

// var inputArray []int = []int{ 2, 1, 3, 2, 3, 4, 5, 6}
// var inputArray []int = []int{ 2, 1, 1, 2, 3, 5, 1, 2, 4}
var inputArray []int = []int{}
var frequencyCount map[int]int = make(map[int]int)

func main() {
	var undefined string
	for i := 0; i <len(inputArray); i++ {
		inputValue := inputArray[i]
		if _, ok := frequencyCount[inputValue]; !ok {
			frequencyCount[inputValue] = 1
		} else {
			fmt.Println(inputValue)
			undefined = "set"
			break;
			// value += 1
			// frequencyCount[inputValue] = value
		}
	}
	if undefined == "" {
		fmt.Println("undefined")
	}

}