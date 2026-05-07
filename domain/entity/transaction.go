package entity

import (
	"github.com/YogiTan00/Management-Transaction/domain/shared"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	id           uuid.UUID
	productID    uuid.UUID
	productTitle string
	totalUnit    int
	returnedUnit int
	date         time.Time
	notes        string

	shared.AuditFields
	shared.SoftDelete
}
