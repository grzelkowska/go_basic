package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	// maps
	phil := map[string]string {"name": "phil", "age": "12"}
	fmt.Println(phil)

	// structs
	favFood := []string{"pizza", "hamburger", "chicken"}
	grze := person{"grze", 19, favFood}
	fmt.Println(grze)
	fmt.Println(grze.name)
	fmt.Println(grze.age)
	fmt.Println(grze.favFood)

	grze2 := person{name: "grze", age: 18, favFood: favFood}
	fmt.Println(grze2)

	
}