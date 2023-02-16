package main

import (
	"fmt"
	"time"

	"github.com/kfelter/limiter"
)

func main() {
	limiter := limiter.New(10)
	for i := 0; i < 1_000; i++ {
		i := i
		limiter.Run(func() {
			// code to be executed with the limiter
			time.Sleep(100 * time.Millisecond)
			fmt.Println("hello", i)
		})
	}
	limiter.Wait()

	fmt.Println("done")
}
