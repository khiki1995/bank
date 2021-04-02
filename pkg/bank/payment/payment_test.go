package payment

import (
	"fmt"

	"bank/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 100_00,
		},
		{
			ID:     2,
			Amount: 101_00,
		},
		{
			ID:     3,
			Amount: 99_00,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output:2
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN:     "5058 0001 0000 8888",
			Balance: 100_00,
			Active:  true,
		},
		{
			PAN:     "5058 0002 0000 8888",
			Balance: 200_00,
			Active:  false,
		},
		{
			PAN:     "5058 0003 0000 8888",
			Balance: 0,
			Active:  true,
		},
		{
			PAN:     "5058 0004 0000 8888",
			Balance: 1_00,
			Active:  true,
		},
	}

	paymentSources := PaymentSources(cards)

	for _, source := range paymentSources {
		fmt.Println(source.Number)
	}
	//Output:
	//5058 0001 0000 8888
	//5058 0004 0000 8888
}
