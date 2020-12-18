package main
import (
	"fmt"
	"unicode/utf8"
	"strings"
)
func main() {
	myValue1 := "Welcome to GeeksforGeeks"
	var myValue2 string
	myValue2 = "GeeksforGeeks"

	// Displaying strings

	fmt.Println("String 1: ",myValue1)
	fmt.Println("String 2: ", myValue2)

	// strings are immutables

	myValue2 = "h"
	fmt.Println("String 2: ",myValue1)

	// String as a range in the for loop
	for index, c := range "GeeksForGeeks" {
		fmt.Printf("The Index number of %c is %d\n", c, index)
	}

	str := "Welcome to GeeksforGeeks"

	for c := 0; c<len(str); c++ {
		fmt.Printf("\n Character = %c Bytes = %v", str[c], str[c])
	}
	// Creating and initializing a slice of byte
	mySlice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}

	// Displaying the string
	mystring1 := string(mySlice1)

	// Create and initializing a sliceof rune
	fmt.Println("String 1: ", mystring1)

	// Creating and initializing a slice of rune

	mySlice2 := []rune{0x0047, 0x0065,0x0065,0x006b,0x0073}

	// Create a string from the slice 
	myString2 := string(mySlice2)

	// Displaying the string 
	fmt.Println("\n String 2: ", myString2)

	length2 := utf8.RuneCountInString(str)
	fmt.Println("string: ", length2)

    // Creating and initializing string
    
	string12 := strings.Trim("!!welcome to GeeksforGeeks !!","!")

	string22  := strings.Trim("@@This is the tutorial of Golang$$","@")

	fmt.Println("\n String! after triming: " )
	fmt.Println("Result 1 :" , string12)
	fmt.Println(string22)
}