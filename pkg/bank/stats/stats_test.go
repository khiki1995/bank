package stats

import (
	"bank/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1212,
			Amount:   3_000,
			Category: "Аква",
		},
		{
			ID:       2212,
			Amount:   2_000,
			Category: "Быстро",
		},
		{
			ID:       3212,
			Amount:   10_000,
			Category: "Бар",
		},
	}

	fmt.Print(Avg(payments))
	//Output: 5000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1212,
			Amount:   3_000,
			Category: "Аква",
		},
		{
			ID:       2212,
			Amount:   2_000,
			Category: "Аква",
		},
		{
			ID:       3212,
			Amount:   10_000,
			Category: "Бар",
		},
	}

	fmt.Print(TotalInCategory(payments, "Аква"))
	//Output: 5000
}
