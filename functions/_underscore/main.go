package main

import "fmt"
func main() {
	mul, _ := mulDiv(105, 7)

	fmt.Println("105 * 7 = ", mul)
}

func mulDiv(n1 int, n2 int) (int,int){
    return n1 * n2, n1/n2
}