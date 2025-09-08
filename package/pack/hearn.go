// Package hearn is my package
package pack

import (
	"errors"
	"fmt"
	"math/rand"

)

func Entry(name string) (string,error){
	//if no value is given, return with a message,instead

	if(name == ""){
		return "",errors.New("How rude, gimme yo name")
	}
	
	//message := fmt.Sprintf("Hi %v! I'm stupid. but I'm going through the best things",name)
	
	//	message := fmt.Sprintf(randFormat(),name)
	message := fmt.Sprintf(randFormat())
	return message,nil
}


func randFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi %v. Welcome!",
		"I hope you die a peaceful death, %v. :p",
		"Einigkeit und recht un freiheit",
		"Fur das deutsche vaterland!",
	}

	// Intn = unsigned integer in the range => 0 <= x < n
	return formats[rand.Intn(len(formats))]
}


//Howdies return a map that associates each of the names people with a greeting message
func Howdies(names[] string) (map[string]string,error){
	out := make(map[string]string)
	//looping through the recieved slice of names , calling 
	//the hello function to get a message for each name
	for _,name := range names {
		output, err  := Entry(name)
		if( err != nil ) {
			return nil,err
		}

		out[name] = output  
	}

	return out,nil
}
