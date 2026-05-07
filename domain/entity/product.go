package entity

import (
	"github.com/YogiTan00/Management-Transaction/domain/shared"
	"github.com/YogiTan00/Management-Transaction/domain/valueobject"
	"github.com/google/uuid"
)

type Product struct {
	id            uuid.UUID
	typeSize      valueobject.TypeSize
	title         string
	price         valueobject.Money
	discount      valueobject.Discount
	discountPrice valueobject.Money

	shared.AuditFields
	shared.SoftDelete
}
