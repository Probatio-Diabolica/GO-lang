package pack

import (
	"testing"
	"regexp"
)

func TestHelloName(t * testing.T){
	name := "Hearn"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg,err := Entry("Hearn")

	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Entry("Hearn")) = %q, %v , want match  for %#q, nil`, msg,err,want)
	}
}

func TestHelloEmpty(t* testing.T){
	msg, err := Entry("")

	if(msg != "" || err == nil){
		t.Errorf(`Entry("") = %q, %v , want match for nil`, msg, err)
	}
}
