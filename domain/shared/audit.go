package shared

import "time"

type AuditFields struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SoftDelete struct {
	DeletedAt *time.Time
}
