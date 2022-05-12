package main

import (
	"fmt"

	go_hello_world "github.com/CoderJava/go-hello-world"
)

func main() {
	fmt.Println(go_hello_world.SayHello())
}

// go get -u . -> update semua dependency-nya
// go get -u github.com/CoderJava/go-hello-world -> untuk menambahkan/update dependency dari plugin ini saja
