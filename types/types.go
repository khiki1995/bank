package types

type PAN string
type Money int64
type Currency string
type CardColor string
type Category string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

const (
	GOLD  CardColor = "GOLD"
	BLACK CardColor = "BLACK"
)

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      CardColor
	Name       string
	Active     bool
}

type Payment struct {
	ID       int
	Amount   Money
	Category Category
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}
