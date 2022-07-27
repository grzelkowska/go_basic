package main

import "fmt"

func main() {
	phil := map[string]string{"name": "phil", "age": "12"}

	for key, value := range phil {
		fmt.Println(key, value)
	}

	fmt.Println(phil)
}