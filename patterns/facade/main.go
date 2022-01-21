package main

import (
	"facade/pkg"
	"fmt"
	"log"
)

func main() {
	fmt.Println()
	walletFacade := pkg.NewWalletFacade("Anton", 123)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("Anton", 123, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
