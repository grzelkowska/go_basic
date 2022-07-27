package main

import (
	"fmt"
	"strings"
	// // "math"
	// "github.com/grzelkowska/go_basic/something"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	defer fmt.Println("Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// for number := range numbers {
	// 	fmt.Println(number + 1)
	// }

	// for index, number := range numbers{
	// 	fmt.Println(index, number)
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// fmt.Println("hello, world")
	// something.SayHello()

	// math.Phi(0)

	// Const: can't be changed
	const name_const string = "Grze"
	const yesOrNo bool = true
	fmt.Println(name_const)
	fmt.Println(yesOrNo)

	// Var: can be changed
	var name_var string = "Grze"
	name_var = "Phil"
	fmt.Println(name_var)

	// instead of    => var name_var string = "Grze" =>
	name_var2 := "phil" // only inside of function // auto-type
	name_var2 = "Phil"
	fmt.Println(name_var2)

	//
	fmt.Println(multiply(2, 2))

	//
	fmt.Println(lenAndUpper("phil"))
	totalLength, upperName := lenAndUpper("phil")
	fmt.Println(totalLength, upperName)
	totalLength2, _ := lenAndUpper("grze")
	fmt.Println(totalLength2)
	//
	totalLength3, upperName3 := lenAndUpper2("grze")
	fmt.Println(totalLength3, upperName3)

	//
	repeatMe("phil", "grze", "dal", "marl", "flynn")

	//
	// superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))
}
