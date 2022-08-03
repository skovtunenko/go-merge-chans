package main

func asChan(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, val := range nums {
			ch <- val
		}
	}()
	return ch
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	fn := func(val int) int { return val * val }
	go func() {
		defer close(out)
		for val := range in {
			out <- fn(val)
		}
	}()
	return out
}
