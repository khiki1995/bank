package card

import (
	"github.com/khiki1995/bank/types"
)

func IssueCard(currency types.Currency, color types.CardColor, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	}

	if card.Balance < amount {
		return card
	}

	if amount < 0 {
		return card
	}

	if amount > 20_000_00 {
		return card
	}

	card.Balance = card.Balance - amount
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}

	if amount > 50_000_00 {
		return
	}

	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance <= 0 {
		return
	}

	bonus := int(card.MinBalance) * percent / 100.0 * daysInMonth / daysInYear
	if bonus > 5_000_00 {
		return
	}
	card.Balance += types.Money(bonus)
}

func Total(cards []types.Card) (total types.Money) {
	for _, card := range cards {
		total += card.Balance
	}
	return
}
