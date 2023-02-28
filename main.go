package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	startTime := time.Now()
	quit := make(chan int)
	length := 15
	elem := rand.Intn(100)
	var arr = []int {elem}
	
	for i := 0; i < length - 1; i++ {
		elem = rand.Intn(100)
		arr = append(arr, elem)
	}
	
	fmt.Printf("%v\n\n", arr)
	
	var timeParts = []time.Time {}

	for i := 0; i < 3; i++ {
		timeParts = append(timeParts, time.Now())
		var part = arr[length/3 * i : length/3 * (i + 1)]
		go calculate(part, i + 1, quit)
	}
	getTime(quit, timeParts, startTime)
}

func calculate(arr []int, index int, quit chan int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}

	fmt.Println(arr)

	if index == 3 {
		quit <- index
	}
}

func getTime(quit chan int, timeParts []time.Time, startTime time.Time) {
	for {
		select {
			case <-quit:
				fmt.Println("------------------------------")

				for i := 0; i < len(timeParts); i++ {
					finishTimePart := time.Now()
					diffPart := finishTimePart.Sub(timeParts[i]).String()
					fmt.Printf("(%d) Time elapsed: %s\n", i + 1, diffPart)	
				}

				finishTime := time.Now()
				diff := finishTime.Sub(startTime).String()
				fmt.Printf("(*) Time elapsed: %s\n", diff)
				return
			default:
				break
		}
	}
}
