package main

import (
	"fmt"
	"sync"
	"time"
)

func getRemoteData(ms time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	duration := ms * time.Millisecond
	time.Sleep(duration)
	fmt.Println("retrieving data in : ", duration)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	go getRemoteData(1000, &wg)
	go getRemoteData(800, &wg)
	go getRemoteData(650, &wg)
	go getRemoteData(100, &wg)

	wg.Wait()
	fmt.Println("finished getting all the data")
}
