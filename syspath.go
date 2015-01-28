package main

import (
	"fmt"
	"os"
	"strings"
)

func list() {
	elements := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
	for _, path := range elements {
		fmt.Println(path)
	}
}

func main() {
	list()
}
