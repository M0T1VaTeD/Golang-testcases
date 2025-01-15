package main

import (
	"fmt"
	"reflect"
	"math/rand"
)

type MyStruct struct {
	IntValue int
	StringValue string
}

func main() {
	flag:=rand.Intn(10)+1
	s := MyStruct{IntValue: 10, StringValue: "Hello"}
	fmt.Println("Before:", s)
	v:=reflect.ValueOf(&s).Elem()
	field:=v.FieldByName("StringValue")
	if flag==3{
		field.Set(reflect.ValueOf(1234))
	}
	fmt.Printf("%d %v", s.IntValue, s.StringValue)
}
