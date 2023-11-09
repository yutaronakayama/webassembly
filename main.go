package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
