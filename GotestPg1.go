package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	args := os.Args[1:]
	var n int
	n, _ = strconv.Atoi(args[0]) //https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go

	// Array
	arr := make([]int, n+1)
	start := time.Now()
	for i := 0; i < n+1; i++ {
		arr[i] = i
	}
	end := time.Now()
	arrRuntime := start.Sub(end)
	fmt.Println("Array initialization runtime: ", arrRuntime)

	//Slice
	s := make([]int, n+1)
	start = time.Now()
	for i := 0; i < n+1; i++ {
		s[i] = i
	}
	end = time.Now()
	sRuntime := start.Sub(end)
	fmt.Println("Slice initialization runtime: ", sRuntime)

	//Map
	m := make(map[int]int)
	start = time.Now()
	for i := 0; i < n+1; i++ {
		m[i] = i
	}
	end = time.Now()
	mRuntime := start.Sub(end)
	fmt.Println("Map initialization runtime: ", mRuntime)

	// Array Increment
	start = time.Now()
	for i := 0; i < n+1; i++ {
		arr[i] = arr[i] + 1
	}
	end = time.Now()
	arrRuntime = start.Sub(end)
	fmt.Println("Array Increment runtime: ", arrRuntime)

	//Slice Increment
	start = time.Now()
	for i := 0; i < n+1; i++ {
		s[i] = s[i] + 1
	}
	end = time.Now()
	sRuntime = start.Sub(end)
	fmt.Println("Slice Increment runtime: ", sRuntime)

	//Map Increment
	m = make(map[int]int)
	start = time.Now()
	for i := 0; i < n+1; i++ {
		m[i] = m[i] + 1
	}
	end = time.Now()
	mRuntime = start.Sub(end)
	fmt.Println("Map increment runtime: ", mRuntime)

}
