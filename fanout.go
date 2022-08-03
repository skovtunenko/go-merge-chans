package main

import "fmt"

func mergeTwo(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for in1 != nil || in2 != nil {
			select {
			case val, ok := <-in1:
				if !ok {
					in1 = nil
					continue
				}
				out <- val
			case val, ok := <-in2:
				if !ok {
					in2 = nil
					continue
				}
				out <- val
			}
		}
	}()
	return out
}

func main() {
	in1 := asChan(1, 1, 1)
	in2 := asChan(2, 2, 2)
	for val := range mergeTwo(in1, in2) {
		fmt.Println(val)
	}
}
