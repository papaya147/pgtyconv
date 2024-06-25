package pgtyconv

import "github.com/jackc/pgx/v5/pgtype"

// FromFloat32 converts a float32 to a pgtype.Float4.
// An empty input float32 (0) will output a null pgtype.Float4.
func FromFloat32(f float32) pgtype.Float4 {
	return pgtype.Float4{Float32: f, Valid: !isZero(f)}
}

// FromFloat32Ptr converts a *float32 to a pgtype.Float4.
// A nil input will output a null pgtype.Float4, otherwise will call FromFloat32.
func FromFloat32Ptr(f *float32) pgtype.Float4 {
	if f == nil {
		return pgtype.Float4{}
	}
	return FromFloat32(*f)
}

// ToFloat32 converts a pgtype.Float4 to a float32.
func ToFloat32(t pgtype.Float4) float32 {
	return ternary(t.Valid, t.Float32, zero[float32]())
}

// ToFloat32Ptr converts a pgtype.Float4 to a *float32.
func ToFloat32Ptr(t pgtype.Float4) *float32 {
	if !t.Valid {
		return nil
	}
	p := ToFloat32(t)
	return &p
}

// FromFloat64 converts a float64 to a pgtype.Float8.
// An empty input float64 (0) will output a null pgtype.Float8.
func FromFloat64(f float64) pgtype.Float8 {
	return pgtype.Float8{Float64: f, Valid: !isZero(f)}
}

// FromFloat64Ptr converts a *float64 to a pgtype.Float8.
// A nil input will output a null pgtype.Float8, otherwise will call FromFloat64.
func FromFloat64Ptr(f *float64) pgtype.Float8 {
	if f == nil {
		return pgtype.Float8{}
	}
	return FromFloat64(*f)
}

// ToFloat64 converts a pgtype.Float8 to a float64.
func ToFloat64(t pgtype.Float8) float64 {
	return ternary(t.Valid, t.Float64, zero[float64]())
}

// ToFloat64Ptr converts a pgtype.Float8 to a *float64.
func ToFloat64Ptr(t pgtype.Float8) *float64 {
	if !t.Valid {
		return nil
	}
	p := ToFloat64(t)
	return &p
}
