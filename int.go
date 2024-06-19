package pgtypes

import "github.com/jackc/pgx/v5/pgtype"

// FromInt64 converts an int64 to a pgtype.Int8.
// An empty input int64 (0) will output a null pgtype.Int8.
func FromInt64(i int64) pgtype.Int8 {
	return pgtype.Int8{Int64: i, Valid: !isZero(i)}
}

// FromInt64Ptr converts a *int64 to a pgtype.Int8.
// A nil input will output a null pgtype.Int8, otherwise will call FromInt64.
func FromInt64Ptr(i *int64) pgtype.Int8 {
	if i == nil {
		return pgtype.Int8{}
	}
	return FromInt64(*i)
}

// ToInt64 converts a pgtype.Int8 to a int64.
func ToInt64(t pgtype.Int8) int64 {
	return ternary(t.Valid, t.Int64, zero[int64]())
}

// ToInt64Ptr converts a pgtype.Int8 to a *int64.
func ToInt64Ptr(t pgtype.Int8) *int64 {
	if !t.Valid {
		return nil
	}
	p := ToInt64(t)
	return &p
}
