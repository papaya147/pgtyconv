package pgtycoon

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// FromUUID converts a uuid.UUID to a pgtype.UUID.
// An empty input uuid (uuid.Nil) will output a null pgtype.UUID.
func FromUUID(u uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: u, Valid: !isZero(u)}
}

// FromUUIDPtr converts a *uuid.UUID to a pgtype.UUID.
// A nil input will output a null pgtype.UUID, otherwise will call FromUUID.
func FromUUIDPtr(u *uuid.UUID) pgtype.UUID {
	if u == nil {
		return pgtype.UUID{}
	}
	return FromUUID(*u)
}

// ToUUID converts a pgtype.UUID to a uuid.UUID.
func ToUUID(u pgtype.UUID) uuid.UUID {
	return ternary(u.Valid, u.Bytes, zero[uuid.UUID]())
}

// ToUUIDPtr converts a pgtype.UUID to a *uuid.UUID.
func ToUUIDPtr(u pgtype.UUID) *uuid.UUID {
	if !u.Valid {
		return nil
	}
	p := ToUUID(u)
	return &p
}
