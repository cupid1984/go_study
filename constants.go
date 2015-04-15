package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Printf("Hello %s\n", World)
	fmt.Printf("Happy %v Day\n", Pi)

	const Truth = true
	fmt.Printf("Go rules? %v\n", Truth)
}