package main

import "fmt"

func main() {
	var age int  = 34
	if(age < 30) {
		fmt.Println("Still have some lives is yet to live")
	} else if(age > 30 && age < 40) {
		fmt.Println("mid age Man, think of investments! or else old age will will be curse")

	} else if(age > 40 && age < 50  ) {
		fmt.Println(" Check ur Health ensurance  and plan better !")
	} else {
		fmt.Println(" Pray God visit Temples !")
	}
}