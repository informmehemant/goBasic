package main

import "fmt"

// Employee => defining a structure
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// Author stucture
type Author struct {
	name string 
	branch string 
	year int
}

// HR nesting structure
type HR struct {

	details Author
}

func main() {
	emp8 := Employee{"Sam", "Anderson", 55, 6000}

	fmt.Println("First Name: ", emp8.firstName)

	emp7 := &Employee{"Hemant", "Singh", 78, 5000}

	fmt.Println("First Name: ", (*emp7).firstName)

	fmt.Println("Age: ", (*emp7).age)
	fmt.Println(emp8.age)
	
	// of the structure 

	result := HR {
		details: Author{"Sona","ECE", 2013},
	}

	// Display the values

	fmt.Println("\n Details of Author")
	fmt.Println(result)

	//  Creating Anonmous functions

	Element := struct {
		name  string
		branch string 
		language string 
		lanaguage string 
		particles int
	} {
		name: "Pikachu",
		branch: "ECE",
		language: "go language",
		particles: 456,
	}

	fmt.Println(Element)
}
