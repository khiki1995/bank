package payment

import "github.com/khiki1995/bank/types"

func Max(payments []types.Payment) (maxPayment types.Payment) {
	for _, payment := range payments {
		if maxPayment.Amount < payment.Amount {
			maxPayment = payment
		}
	}
	return
}

func PaymentSources(cards []types.Card) (pSources []types.PaymentSource) {
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			pSources = append(pSources, types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			})
		}
	}
	return
}
