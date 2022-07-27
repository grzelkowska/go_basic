package main

import "fmt"

func main() {
	names_five := [5]string{"nico", "phil", "grze"}
	names_five[3] = "alala"
	names_five[4] = "lalala"
	// names[5] = "nonono"
	fmt.Println(names_five)

	names := []string{"phil", "grze", "nico"}
	fmt.Println(names)
	names = append(names, "flynn")
	fmt.Println(names)
}
