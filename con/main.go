package main

import (
	"sync"
	"fmt"
)


func infite(wg *sync.WaitGroup){
	defer wg.Done()
	i:= 0
	for j:= 0 ; j < 100 ; j++{
		fmt.Println("{%v} A", j)
		//i = j
		j=i
	}
}

func infi(wg *sync.WaitGroup){
	defer wg.Done()
	i:= 0
	for j:= 0 ; j < 100 ; j++{
		fmt.Println("{%v} B", j)
		//i = j
		j=i
	}
}

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	
	go infite(&wg)
	go infi(&wg)


	//go infite()
	//go infi()
	//infi()

//	i:= 0
//	for j:= range 10{
	//	fmt.Println("{%d} a", j)
//		j=i
//	}
	wg.Wait()
}
