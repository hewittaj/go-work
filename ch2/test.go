package main

import "fmt"

func main() {
	// Exercise 1 and 2
	const value = 34
	i := 20
	f := float64(i)
	fmt.Println(i)
	fmt.Println(f)
	i = value
	f = value
	fmt.Println(i)
	fmt.Println(f)

	// Exercise 3
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
