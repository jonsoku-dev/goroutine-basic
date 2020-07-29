package main

import (
	"fmt"
)

func accounts() {
	account := accounts.NewAccount("jongseok")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	// fmt.Println(account.Balance(), account.Owner())
	account.String()
}
