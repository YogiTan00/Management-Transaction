package valueobject

type Discount int

func NewDiscount(discount int) Discount {
	if discount < 0 {
		return Discount(0)
	}
	if discount > 100 {
		return Discount(100)
	}
	return Discount(discount)
}
