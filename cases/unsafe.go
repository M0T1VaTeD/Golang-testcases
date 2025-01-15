package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//задана переменная x типа int
	var x int = 52

	//переменной присвоен указатель при помощи unsafe.Pointer
	ptr := unsafe.Pointer(&x)

	//указатель преобразуется из int в float
	floatPtr := (*float32)(ptr)

	fmt.Println(*floatPtr)
}
