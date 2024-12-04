package main

import (
	"sync"

	"service-api/bootstrap"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go bootstrap.ServerApi(&wg)
	go bootstrap.HeapProfile(&wg)
	wg.Wait()
}
