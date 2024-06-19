package pgtyconv

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/papaya147/randomize"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromInt64(t *testing.T) {
	i, err := randomize.Do[int64]()
	require.NoError(t, err)
	pgi := FromInt64(i)
	require.Equal(t, i, pgi.Int64)
	require.Equal(t, true, pgi.Valid)

	i = 0
	pgi = FromInt64(i)
	require.Equal(t, i, pgi.Int64)
	require.Equal(t, false, pgi.Valid)
}

func TestFromInt64Ptr(t *testing.T) {
	i, err := randomize.Do[int64]()
	require.NoError(t, err)
	pgi := FromInt64Ptr(&i)
	require.Equal(t, i, pgi.Int64)
	require.Equal(t, true, pgi.Valid)

	i = 0
	pgi = FromInt64Ptr(&i)
	require.Equal(t, i, pgi.Int64)
	require.Equal(t, false, pgi.Valid)

	pgi = FromInt64Ptr(nil)
	require.Equal(t, int64(0), pgi.Int64)
	require.Equal(t, false, pgi.Valid)
}

func TestToInt64(t *testing.T) {
	pgi, err := randomize.Do[pgtype.Int8]()
	require.NoError(t, err)
	pgi.Valid = true
	i := ToInt64(pgi)
	require.Equal(t, i, pgi.Int64)

	pgi.Valid = false
	i = ToInt64(pgi)
	require.Equal(t, int64(0), i)
}

func TestToInt64Ptr(t *testing.T) {
	pgi, err := randomize.Do[pgtype.Int8]()
	require.NoError(t, err)
	pgi.Valid = true
	i := ToInt64Ptr(pgi)
	require.Equal(t, *i, pgi.Int64)

	pgi.Valid = false
	i = ToInt64Ptr(pgi)
	require.Nil(t, i)
}
