/*
Task 1: Arrays
Implement function that returns an average value of array (sum / N)
input -> [1,2,3,4,5,6]
output -> 3.5
*/

package homework

import "fmt"

func average(input [15]float32) (result float32) {

	for index := range input {
		result += input[index]
		fmt.Println(result)
	}

	result = result / float32(len(input))
	return
}
