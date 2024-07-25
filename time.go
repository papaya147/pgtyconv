package pgtyconv

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// FromUnixTime converts an int64 to a pgtype.Timestamptz.
// An empty input int64 (0) will output a null pgtype.Timestamptz.
func FromUnixTime(i int64, u TimeUnit) pgtype.Timestamptz {
	p := u.precision()
	s := i / p
	n := (i % p) * (1e9 / p)
	return pgtype.Timestamptz{Time: time.Unix(s, n), Valid: !isZero(i)}
}

// FromUnixTimePtr converts a *int64 to a pgtype.Timestamptz.
// A nil input will output a null pgtype.Timestamptz, otherwise will call FromUnixTime.
func FromUnixTimePtr(i *int64, u TimeUnit) pgtype.Date {
	if i == nil {
		return pgtype.Date{}
	}
	return FromUnixDate(int64(*i), u)
}

// ToUnixTime converts a pgtype.Timestamptz to a int64.
func ToUnixTime(d pgtype.Timestamptz, u TimeUnit) int64 {
	p := u.precision()
	n := d.Time.UnixNano()
	return n / (1e9 / p)
}

// ToUnixTimePtr converts a pgtype.Timestamptz to a *int64.
func ToUnixTimePtr(d pgtype.Timestamptz, u TimeUnit) *int64 {
	if !d.Valid {
		return nil
	}
	p := ToUnixTime(d, u)
	return &p
}
