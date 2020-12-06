package main

import "fmt"

func main() {
	 // Swtich statement with both
	//  optional statement 
	// and expression

	// switch <optional statements>; <expression> {}

	switch day:=4 ; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wedneday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Firday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	//  without expression statements and expression

	var value string = "India"
	switch {
	case value == "Uk":
		fmt.Println("Hello")
	case value == "India":
		fmt.Println("Namskara")
	case value == "France":
		fmt.Println("Bonjour")
	}	

	// without default statements
	// Mutiple values in case statement
    var val string = "three"
	switch val {
	case "one":
		fmt.Println("c#")
	case "two", "three":
		fmt.Println("go")
	case "four", "five", "six":
		fmt.Println("Java")
	}
}