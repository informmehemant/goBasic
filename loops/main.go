package main

 
import "fmt"

func main() {

	rSting := []string{"Banana","Apple","Orange"}
	// Basic Loops 
	for i :=0; i< 4; i++ {
		fmt.Println("GeeksforGeeks")
	}
	// infinity loops

	// for{
	// 	fmt.Printf("GeeksofGeeks")
	// }

	// while loop

	i := 0
	for i<3 {
		i += 2
		fmt.Println(i)
	}

	// simple Range for for loop
	for idx, val:=range rSting {
		fmt.Println(idx, val)
	}
	// let loop in character from strings
	for index, chr := range "Go-lang-is-cools" {
		fmt.Printf(" %c is %d \n", chr,index )
	}

	// using map 
    mVal := map[int]string {
		22: "Jack fruits",
		23: "Lemon",
		34: "Musk Melon",
	}

	for idx, val:= range mVal {
		fmt.Println(idx, val)
	}

	//  use for loop using channel
	chn1 := make(chan int)
	go func() {
		chn1 <- 100
		chn1 <- 1000
		chn1 <- 100000
		chn1 <- 1000000
		close(chn1)
	}()
	//  This is the above function to create channel
	for i:= range chn1 {
		fmt.Println(i)
	}

	
}