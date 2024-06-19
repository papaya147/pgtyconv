package pgtycoon

import (
	"github.com/papaya147/randomize"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromUnixDate(t *testing.T) {
	d, err := randomize.Do[uint32]()
	require.NoError(t, err)
	tu, err := randomize.Do[TimeUnit]()
	require.NoError(t, err)
	pgd := FromUnixDate(int64(d), tu)
	e := ToUnixDate(pgd, tu)
	require.Equal(t, int64(d), e)

	d = 0
	pgd = FromUnixDate(int64(d), tu)
	require.Equal(t, false, pgd.Valid)
	require.Equal(t, int64(d), pgd.Time.Unix())
}
