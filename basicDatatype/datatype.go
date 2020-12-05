package main
import "fmt"

func main() {
// using 8-bit unsigned int

	var X uint8 = 225
	fmt.Println(X+1, X)
// Using 16-bit signed int
var Y int16 = 32767
fmt.Println(Y, Y+2)

// substraction of two floating number

a:= 20.45
b:= 34.89
c:= b-a
// Display the result

fmt.Printf("Result is: %f", c)

// Display the type of c variable 

fmt.Printf("\n The type of the c is: %T",c)

// the use of complex numbers
var complexA complex128 = complex(6,2)
var complexB complex64 = complex(9,2)
fmt.Println(complexA)
fmt.Println(complexB)

fmt.Printf("The Type of a is %T and "+
			"the type of b is %T", complexA,complexB)
			
// the use of the booleans

str1:= "GeeksforGeeks"
str2:= "geeksforgeeks"
str3:= "GeeksforGeeks"

fmt.Println(str1 == str2)
fmt.Println(str1 == str3)

fmt.Printf("Length of the string is: %d", len(str1))
fmt.Printf("\n String is: %s", str1)
}