package main

import (
	"strings"
	"fmt"
)
func joinstr(element...string) string {
	return strings.Join(element, "-")
}

func main() {
	// Variadic function to joun strings
	fmt.Println(joinstr())

	// mutiple arguments
	fmt.Println(joinstr("Geek","Geeks"))
	fmt.Println(joinstr("Geek","for","Geeks"))
	fmt.Println(joinstr("G","E","E","k","s"))

	// pass a slice in variadic function
	
}