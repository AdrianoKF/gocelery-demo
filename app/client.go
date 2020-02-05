package main

import (
	"fmt"
	"github.com/AdrianoKF/gocelery-demo/celery"
	"math/rand"
	"time"
)

func main() {
	// prepare arguments
	taskName := "add"
	argA := rand.Intn(10)
	argB := rand.Intn(10)

	// run task
	client := celery.CreateClient()

	asyncResult, err := client.Delay(taskName, argA, argB)
	if err != nil {
		panic(err)
	}

	result, _ := asyncResult.Get(1 * time.Second)
	fmt.Println("Result: ", result)
}
