package valueobject

type Money int64

func NewMoney(money int64) Money {
	if money < 0 {
		return Money(0)
	}
	return Money(money)
}
