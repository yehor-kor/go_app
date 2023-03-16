package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	start(100, 5)
}

func start(length, times int) {
	startTime := time.Now()
	quit := make(chan int)
	elem := rand.Intn(100)
	var arr = []int {elem}
	
	for i := 0; i < length - 1; i++ {
		elem = rand.Intn(100)
		arr = append(arr, elem)
	}
	
	fmt.Printf("%v\n\n", arr)
	
	var timeParts = []time.Time {}

	for i := 0; i < times; i++ {
		timeParts = append(timeParts, time.Now())
		var startPos = length/times * i
		var endPos int

		if i < times - 1 {
			endPos = length/times * (i + 1)
		} else {
			endPos = length/times * (i + 1) + length % times
		}

		var part = arr[startPos : endPos]
		go calculate(part, i + 1, times, quit)
	}
	getTime(quit, timeParts, startTime)
}

func calculate(arr []int, index int, times int, quit chan int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}

	// fmt.Println(arr)

	if index == times {
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
