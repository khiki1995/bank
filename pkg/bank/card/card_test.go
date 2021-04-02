package card

import (
	"bank/types"
	"fmt"
)

func ExampleDeposit() {
	card := types.Card{
		Active: true,
	}

	Deposit(&card, 100_00) //positive
	fmt.Println(card.Balance)

	Deposit(&card, 50_000_01) //limit
	fmt.Println(card.Balance)

	card.Active = false
	Deposit(&card, 100_00) //limit
	fmt.Println(card.Balance)
	// Output:
	// 10000
	// 10000
	// 10000
}

func ExampleAddBonus() {
	card := types.Card{
		Active:     true,
		Balance:    100_00,
		MinBalance: 10_000_00,
	}

	AddBonus(&card, 3, 30, 365) //positive
	fmt.Println(card.Balance)

	card.Active = false
	AddBonus(&card, 3, 30, 365) //inActive
	fmt.Println(card.Balance)

	card.Active = true
	card.Balance = 0
	AddBonus(&card, 3, 30, 365) //Balance > 0
	fmt.Println(card.Balance)

	card.Active = true
	card.Balance = 100_00
	card.MinBalance = 1_000_000
	AddBonus(&card, 3, 30, 365) //limit
	fmt.Println(card.Balance)
	// Output:
	// 12465
	// 12465
	// 0
	// 12465
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 100_00,
			Active:  true,
		},
		{
			Balance: 100_00,
			Active:  true,
		},
		{
			Balance: 100_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	//Output:30000
}
