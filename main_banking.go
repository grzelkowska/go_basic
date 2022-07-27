package main

import (
	"fmt"
	// "log"

	"github.com/grzelkowska/go_basic/accounts"
)

func main() {
	// account := accounts.Account{Owner: "grze"}
	// account.Owner = "phil"
	// fmt.Println(account)

	account := accounts.NewAccount("phil")
	account.Deposit(100)
	fmt.Println(account.Balance())

	err := account.Withdraw(10)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance())

	err1 := account.Withdraw(1000)
	if err1 != nil {
		// log.Fatalln(err)
		fmt.Println(err1)
	}
	fmt.Println(account.Balance())


	fmt.Println(account.Owner())
	account.ChangeOwner("grze")
	fmt.Println(account.Owner())

	fmt.Println(account)


}
