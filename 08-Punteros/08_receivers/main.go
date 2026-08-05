package main

import "fmt"

type Counter struct {
	number int
}

func (counter *Counter) Increment() {
	counter.number++
}

func (counter Counter) Value() int {
	return counter.number
}

func main() {
	counter := Counter{10}
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()

	fmt.Println(counter.Value())
}
