package pgtyconv

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type TimeUnit int

const (
	TimeUnitNanoseconds TimeUnit = iota
	TimeUnitMicroseconds
	TimeUnitMilliseconds
	TimeUnitSeconds
)

func (u *TimeUnit) precision() int64 {
	switch *u {
	case TimeUnitNanoseconds:
		return 1e9
	case TimeUnitMicroseconds:
		return 1e6
	case TimeUnitMilliseconds:
		return 1e3
	default:
		return 1
	}
}

// FromUnixDate converts an int64 to a pgtype.Date.
// An empty input int64 (0) will output a null pgtype.Date.
func FromUnixDate(i int64, u TimeUnit) pgtype.Date {
	p := u.precision()
	s := i / p
	n := (i % p) * (1e9 / p)
	return pgtype.Date{Time: time.Unix(s, n), Valid: !isZero(i)}
}

// FromUnixDatePtr converts a *int64 to a pgtype.Date.
// A nil input will output a null pgtype.Date, otherwise will call FromUnixDate.
func FromUnixDatePtr(i *int64, u TimeUnit) pgtype.Date {
	if i == nil {
		return pgtype.Date{}
	}
	return FromUnixDate(int64(*i), u)
}

// ToUnixDate converts a pgtype.Date to a int64.
func ToUnixDate(d pgtype.Date, u TimeUnit) int64 {
	p := u.precision()
	n := d.Time.UnixNano()
	return n / (1e9 / p)
}

// ToUnixDatePtr converts a pgtype.Date to a *int64.
func ToUnixDatePtr(d pgtype.Date, u TimeUnit) *int64 {
	if !d.Valid {
		return nil
	}
	p := ToUnixDate(d, u)
	return &p
}
