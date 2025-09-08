package main

import (
	"fmt"
)

func main(){
	cFOR()
	//cWHILE()
	//cInf();
	/*s := []string{
		"equation",
		"eigen",
		"gin",
		"hearn",
	}*/
	
	//iterating(s)
	//fmt.Printf("\n")
	//runTimeIT();
}


func cFOR(){
	//syntax: for init; condition ; post {}
	for i :=0; i < 10; i++{
		fmt.Printf("%v \n", i)
	}

	for i , j := 0 , 1 ; i < 10 && j < 5 ; i , j = i+1, j +6{
		fmt.Println("I hate myself\n")
	}
}

func cWHILE() {
	//syntax : for condition { ... }
	val := 10
	for val < 20{
		fmt.Printf("%v\n",val)
		val++
	}
}

func cInf(){
	//syntax : for(;;)
	for {}
}

// iterating is just iterating on a string slice
func iterating(a []string){
	for idx,val := range a{
		fmt.Printf("idx : %v val: %v\n",idx , val)
	}
}

func runTimeIT(){
	for idx,el := range "grenzen_hearn12345678987654dfghj"{
		fmt.Printf("%v | %c \n",idx,el)
	}
}
