package main

import (
	"github.com/yaowenqiang/moderngo/protobuf"
    "fmt"
)

func main() {
    animal := new(protobuf.Animal)
    animal.Age = 10

    fmt.Println(animal)
    fmt.Println("modern go")
}
