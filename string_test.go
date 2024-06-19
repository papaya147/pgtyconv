package pgtyconv

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/papaya147/randomize"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromString(t *testing.T) {
	s, err := randomize.Do[string]()
	require.NoError(t, err)
	pgs := FromString(s)
	require.Equal(t, s, pgs.String)

	s = ""
	pgs = FromString(s)
	require.Equal(t, s, pgs.String)
	require.Equal(t, false, pgs.Valid)
}

func TestFromStringPtr(t *testing.T) {
	s, err := randomize.Do[string]()
	require.NoError(t, err)
	pgs := FromStringPtr(&s)
	require.Equal(t, s, pgs.String)

	s = ""
	pgs = FromStringPtr(&s)
	require.Equal(t, s, pgs.String)
	require.Equal(t, false, pgs.Valid)

	pgs = FromStringPtr(nil)
	require.Equal(t, "", pgs.String)
	require.Equal(t, false, pgs.Valid)
}

func TestToString(t *testing.T) {
	pgs, err := randomize.Do[pgtype.Text]()
	require.NoError(t, err)
	pgs.Valid = true
	s := ToString(pgs)
	require.Equal(t, pgs.String, s)

	pgs.Valid = false
	s = ToString(pgs)
	require.Equal(t, "", s)
}

func TestToStringPtr(t *testing.T) {
	pgs, err := randomize.Do[pgtype.Text]()
	require.NoError(t, err)
	pgs.Valid = true
	s := ToStringPtr(pgs)
	require.Equal(t, pgs.String, *s)

	pgs.Valid = false
	s = ToStringPtr(pgs)
	require.Nil(t, s)
}
