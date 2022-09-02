package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	var n int
	n, _ = strconv.Atoi(args[0]) //https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go

	arr := make([]int, n)
	for i := 0; i < n+1; i++ {
		arr[i] = rand.Intn(100)
	}

	start := time.Now()
	sort.Sort([]int(arr))
	end := time.Now()
	arrRuntime := start.Sub(end)
	fmt.Println("Sort for Array: ", arrRuntime)

	arr = make([]int, n)
	for i := 0; i < n+1; i++ {
		arr[i] = rand.Intn(100)
	}

	start = time.Now()
	sort.Stable(arr)
	end = time.Now()
	arrRuntime = start.Sub(end)
	fmt.Println("Stablex for Array: ", arrRuntime)

	s := make([]int, n)
	for i := 0; i < n+1; i++ {
		s[i] = rand.Intn(100)
	}

	start = time.Now()
	sort.Sort(s)
	end = time.Now()
	sRuntime := start.Sub(end)
	fmt.Println("Sort for Slice: ", sRuntime)

	s = make([]int, n)
	for i := 0; i < n+1; i++ {
		s[i] = rand.Intn(100)
	}

	start = time.Now()
	sort.Stable(s)
	end = time.Now()
	sRuntime = start.Sub(end)
	fmt.Println("Stable for Slice: ", sRuntime)

}
