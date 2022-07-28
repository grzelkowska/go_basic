package main

import (
	"fmt"

	"github.com/grzelkowska/go_basic/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{}
	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)

	dictionary := mydict.Dictionary{"first": "First word"}
	// fmt.Println(dictionary["first"])

	// Search
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	definition1, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition1)
	}
	fmt.Println("\n")

	// Add
	word := "hello"
	meaning := "Greeting"

	err1 := dictionary.Add(word, meaning)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		definition2, err := dictionary.Search("hello")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition2)
		}
	}

	err2 := dictionary.Add(word, meaning)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		definition3, err := dictionary.Search("hello")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition3)
		}
	}
	fmt.Println("\n")

	// Update
	word2 := "update"
	err3 := dictionary.Add(word2, "before update")
	if err3 != nil {
		fmt.Println(err3)
	}
	definition4, _ := dictionary.Search(word2)
	fmt.Println(definition4)
	err4 := dictionary.Update("word2", "after update")
	if err4 != nil {
		fmt.Println(err4)
	}
	definition5, _ := dictionary.Search(word2)
	fmt.Println(definition5)
	fmt.Println("\n")

	// Delete
	err5 := dictionary.Delete(word2)
	if err5 != nil {
		fmt.Println(err5)
	}
	definition6, err := dictionary.Search(word2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(definition6)

}
