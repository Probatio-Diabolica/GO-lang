package main

import (
	"fmt"
	//"os"
	"bytes"
	"sync"
)

type resource struct{
	lock sync.Mutex
	buffer bytes.Buffer
}

func main(){
	fmt.Println("howdy !!")
	
	red := new(resource) //red is a pointer | red = *resource

	var green resource //green is not a pointer | green = resource

	fmt.Println(*red,green)

	//new

//	fmt.Println(makeFile())

	//make 
	heapArr := make([] int,10,100)
	fmt.Println(heapArr[3])
	
	//array
	sol := []int {1,2,3,4,5}
	fmt.Println(sol)
	

	//2d slice
	tuD := [2][2]int16 {{1,2},{3,4}}
	fmt.Println(tuD)

	bites := [][]byte {
		[]byte("Howdy!! "),
		[]byte("Nada!! bitch"),
	}

	fmt.Println(bites)

	for _ , el := range bites{
		for ids,ele := range el{
			fmt.Println(ids," ",ele)
		}
		fmt.Println()
	}
}

func makeFile(fileDesc int, path string){
	if(fileDesc > 0 ) {
		fmt.Println()
	}
	
	 // this is ust a pointer to our file resource
	 //file := os.File{fileDesc, path , nil , 0}

	 //return &file

	 //return &os.File(fileDesc, path)
}
