package main

import (
	"fmt"

	"github.com/buraksekili/aigent-application/task1"
	"github.com/buraksekili/aigent-application/task2"
)

func main() {
	input1 := "cancel transaction"
	input2 := [][]string{{"terminate", "close"}, {"deal", "prospect"}}

	result := task1.Task1(input1, input2)

	fmt.Println("TASK1:")
	for _, r := range result {
		fmt.Println(r)
	}
	
	fmt.Println("\nTASK2:")
	task2.Task2()
}
