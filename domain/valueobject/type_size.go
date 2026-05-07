package valueobject

type TypeSize int

func NewTypeSize(typeSize int) TypeSize {
	switch typeSize {
	case 0:
		return TypeSize(0)
	case 1:
		return TypeSize(1)
	default:
		return TypeSize(0)
	}
}
