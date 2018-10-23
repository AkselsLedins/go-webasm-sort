package main

import (
	"fmt"
	js "syscall/js"
)

func bubbleSort(array []js.Value) {
	a := make([]int, len(array))

	for index, element := range array {
		a[index] = element.Int()
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Printf("gof(bubbleSort) %v > %v\n", array, a)
}

func registerCallbacks() {
	js.Global().Set("bubbleSort", js.NewCallback(bubbleSort))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM: Initialized")

	registerCallbacks()

	<-c
	fmt.Println("WASM: Over")
}
