package pgtyconv

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/papaya147/randomize"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromFloat32(t *testing.T) {
	i, err := randomize.Do[float32]()
	require.NoError(t, err)
	pgi := FromFloat32(i)
	require.Equal(t, i, pgi.Float32)
	require.Equal(t, true, pgi.Valid)

	i = 0
	pgi = FromFloat32(i)
	require.Equal(t, i, pgi.Float32)
	require.Equal(t, false, pgi.Valid)
}

func TestFromFloat32Ptr(t *testing.T) {
	i, err := randomize.Do[float32]()
	require.NoError(t, err)
	pgi := FromFloat32Ptr(&i)
	require.Equal(t, i, pgi.Float32)
	require.Equal(t, true, pgi.Valid)

	i = 0
	pgi = FromFloat32Ptr(&i)
	require.Equal(t, i, pgi.Float32)
	require.Equal(t, false, pgi.Valid)

	pgi = FromFloat32Ptr(nil)
	require.Equal(t, float32(0), pgi.Float32)
	require.Equal(t, false, pgi.Valid)
}

func TestToFloat32(t *testing.T) {
	pgi, err := randomize.Do[pgtype.Float4]()
	require.NoError(t, err)
	pgi.Valid = true
	i := ToFloat32(pgi)
	require.Equal(t, i, pgi.Float32)

	pgi.Valid = false
	i = ToFloat32(pgi)
	require.Equal(t, float32(0), i)
}

func TestToFloat32Ptr(t *testing.T) {
	pgi, err := randomize.Do[pgtype.Float4]()
	require.NoError(t, err)
	pgi.Valid = true
	i := ToFloat32Ptr(pgi)
	require.Equal(t, *i, pgi.Float32)

	pgi.Valid = false
	i = ToFloat32Ptr(pgi)
	require.Nil(t, i)
}
