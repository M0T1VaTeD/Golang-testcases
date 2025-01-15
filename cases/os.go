package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	name := fmt.Sprintf("tmp.txt")
	file, err := os.Create(name)
	if err!=nil {
		log.Fatalf("%v", err)
	}
	n, err := file.WriteString("Hello world")
	if err!=nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(n)
}
