package card

import (
	"bank/pkg/bank/types"
)

// PaymentSources function for choosing payment source, return payment source
func PaymentSources (cards []types.Card) []types.PaymentSource {
	
	payment := []types.PaymentSource{}

	for _, card := range cards {
		if card.Active == true && card.Balance > 0 {
			payment = append(payment, types.PaymentSource{Number: card.PAN})
		}
	}
	return payment
}