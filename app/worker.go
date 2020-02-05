package main

import (
	"context"
	"fmt"
	"github.com/AdrianoKF/gocelery-demo/celery"
	"time"
)

func add(a int, b int) int {
	fmt.Printf("Calculating %d + %d\n", a, b)
	return a + b
}

func main() {
	fmt.Println("Starting Go Celery worker")

	client := celery.CreateClient()

	client.Register("add", add)

	ctx, cancel := context.WithCancel(context.Background())

	client.StartWorkerWithContext(ctx)
	time.Sleep(10 * time.Second)

	cancel()
	client.WaitForStopWorker()
}
