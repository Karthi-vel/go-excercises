package main

import (
	"fmt"
	"strings"
)

func main() {
	name := []string{"Karthi", "Priya"}
	byteName := []byte(strings.Join(name, ","))
	fmt.Println(byteName)
}