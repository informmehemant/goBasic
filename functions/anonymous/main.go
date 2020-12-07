// go program to illustrate 
// Create an anonymous fucntion
package main 
import (
	"fmt"
	"sort"
	"strings"
	"time"
)
func main() {
	func() {
		// Anonymous function
		fmt.Println("Welcome! to GeeksforGeeks")
	}()

	// utilise on go function
	// can assign function to variable

	value := func() {
		fmt.Println("I can assigned to varible")
	}
	value()

	// passing function as closures
	// anon := func(s name) string {
	// 	return "Hemant, " + name  
	// }
	

	// anotherFunction(anon)
	partial := partialSum(3)
	fmt.Println(partial(7))
	
	// main function and the init functions in the Go

	s := []int{-2,7,1,90,100,5}
	sort.Ints(s)
	fmt.Println("sorted values ", s)
	fmt.Println("Index values: ", strings.Index("GeeksforGeeks","ks"))
	fmt.Println("Time: ", time.Now().Unix())
    
}
func sum(x,y int) int {
   return x + y
} 
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x,y)
	}
}
func anotherFunction(f func(string) string) {
	result := f("David")
	fmt.Println(result) // Prints "Hiya, David"
}