package main

import (
	"fmt"
)

func main() {
	// age:= 16

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }

	// age:= 10

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is a teenager")
	// } else {
	// 	fmt.Println("person is a kid")
	// }

	var role = "admin"
	var hasPermissions = false

	// or operator
	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}

	// and operator
	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	}

	//  we can decalre var inside if construct
	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is an teanager", age)
	} else {
		fmt.Println("person is a kid", age)
	}

	//  go does not have terranry operator

}
