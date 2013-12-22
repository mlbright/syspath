package main

import (
	"fmt"
	"os"
	"strings"
)

func list() {
	fmt.Println("elements of %PATH%:\n")
	elements := strings.Split(os.Getenv("PATH"), ";")
	for _, path := range elements {
		fmt.Println(path)
	}
}

func main() {
	list()
}
