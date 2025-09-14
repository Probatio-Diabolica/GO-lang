package main

import (
	"fmt"
	"sync"
)

func infite(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 100; j++ {
		fmt.Println("{%v} A", j)
	}
}

func infi(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 100; j++ {
		fmt.Println("{%v} B", j)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go infite(&wg)
	go infi(&wg)

	wg.Wait()
}
