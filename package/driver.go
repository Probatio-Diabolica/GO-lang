package main

import (
	"fmt"
	"log"

	"hearn/pack"

)

func main() {

	//these are just the customizers for the standard log library
	log.SetPrefix("Howdy: ") //ok this just means our logs will begin with "Howdy :
	log.SetFlags(0) // 0 = false, which means it hides the default date/time infos

	// We are requesting a message
	out,err := pack.Entry("eigen")
	
	if(err != nil){
		log.Fatal(err)
	}

	fmt.Println(out)

	//out2,errr := pack.Entry("")

	//if(errr != nil){
	//	fmt.Println("wut") //execution happens
	//	log.Fatal(errr)
	//}
	
	//fmt.Println(out2)

	names := []string{"Eigen","Equation","Hearn"}

	message,prob := pack.Howdies(names)

	if(prob != nil){
		log.Fatal(prob)
	}

	fmt.Println(message)
}
