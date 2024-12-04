package main

import (
	"sync"

	"service-api/bootstrap"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go bootstrap.ServerApi(&wg)
	wg.Wait()
}
