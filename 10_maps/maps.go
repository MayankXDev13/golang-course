package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// get an element 
	// fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it return zero value

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))


	// delete(m, "price")

	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m)


	// fmt.Println(m)
	
	
	// v, ok := m["price"]
	// fmt.Println(v, ok) // 0 true
	
	
	// if  ok {
		// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	
	m1 := map[string]int{"price": 0, "age": 0}
	m2 := map[string]int{"price": 0, "age": 0}


	fmt.Println(maps.Equal(m1, m2))



	
}

