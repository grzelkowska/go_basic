package main

import "fmt"

func canIDrink(age int) bool {
	// if age < 18 {
	// 	return false
	// }
	// return true
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	// switch age {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
	// return false

	// switch {
	// case age <= 18 :
	// 	return false
	// case age == 19:
	// 	return true
	// case age > 50:
	// 	return false
	// }
	// return false

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkSwitch(18))
}
