package main

import (
	"fmt"

	go_say_hello "github.com/CoderJava/go-hello-world/v2"
)

func main() {
	fmt.Println(go_say_hello.SayHello("Yudi"))	
}

// go get -u . -> update semua dependency-nya
// go get -u github.com/CoderJava/go-hello-world -> untuk menambahkan/update dependency dari plugin ini saja
