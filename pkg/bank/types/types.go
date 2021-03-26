package types

// Payment to present of information of history payments
type Payment struct {
	ID int
	Amount Money
}

// Money to present yourself payment minimaled summe in currency.
type Money int64

// Currency to present yourself code of currency
type Currency string

// codes of currency
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN to present number of card
type PAN string 

// Card to present yourself informational of a payment card 
type Card struct {
	ID int
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}

// PaymentSource presents a resource for choosing a payment method (card)
type PaymentSource struct {
	Type string
	Number PAN
	Balance Money
}
