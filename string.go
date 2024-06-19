package pgtycoon

import "github.com/jackc/pgx/v5/pgtype"

// FromString converts a string to a pgtype.UUID.
// An empty input string ("") will output a null pgtype.Text.
func FromString(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: !isZero(s)}
}

// FromStringPtr converts a *string to a pgtype.Text.
// A nil input will output a null pgtype.Text, otherwise will call FromString.
func FromStringPtr(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}
	return FromString(*s)
}

// ToString converts a pgtype.Text to a string.
func ToString(t pgtype.Text) string {
	return ternary(t.Valid, t.String, zero[string]())
}

// ToStringPtr converts a pgtype.Text to a *string.
func ToStringPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	p := ToString(t)
	return &p
}
