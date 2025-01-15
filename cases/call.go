package main

import (
	"fmt"
	"syscall"
	"log"
	"os"
)

func main() {
	file, err:=syscall.Open("tmp.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err!=nil {
		log.Fatalf("%v", err)
	}
	data := []byte("Hello world")
	n, err:=syscall.Write(file, data)
	if err!=nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(n)
}
