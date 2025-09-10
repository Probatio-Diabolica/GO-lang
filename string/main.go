package main

import (
	"fmt"
)

func main(){
	s1 := "string\nbroke"
	s2 := `raw \n string \n literal`
	//both are different
	fmt.Println(s1,'\n',s2) //returns the ascii val of \n
	fmt.Println(s1,"\n",s2) //does exactly what \n aims to do
	
	changingStr()
}


func changingStr(){
	s:= "fucker"
	//s[0] = 'D'
	//above is not allowed because string are immutable

	//we can change it like this:
	
	b := []byte(s)
	b[0] = 'D'
	
	s2 := string(b)
	fmt.Println(s2)
}
