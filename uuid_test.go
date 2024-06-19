package pgtypes

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/papaya147/randomize"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromUUID(t *testing.T) {
	u, err := randomize.Do[uuid.UUID]()
	require.NoError(t, err)
	pgu := FromUUID(u)
	require.Equal(t, u, uuid.UUID(pgu.Bytes))
	require.Equal(t, true, pgu.Valid)

	u = uuid.Nil
	pgu = FromUUID(u)
	require.Equal(t, u, uuid.UUID(pgu.Bytes))
	require.Equal(t, false, pgu.Valid)
}

func TestFromUUIDPtr(t *testing.T) {
	u, err := randomize.Do[uuid.UUID]()
	require.NoError(t, err)
	pgu := FromUUIDPtr(&u)
	fmt.Println(u)
	fmt.Println(pgu)
	require.Equal(t, u, uuid.UUID(pgu.Bytes))
	require.Equal(t, true, pgu.Valid)

	u = uuid.Nil
	pgu = FromUUIDPtr(&u)
	require.Equal(t, u, uuid.UUID(pgu.Bytes))
	require.Equal(t, false, pgu.Valid)

	pgu = FromUUIDPtr(nil)
	require.Equal(t, u, uuid.UUID(pgu.Bytes))
	require.Equal(t, false, pgu.Valid)
}

func TestToUUID(t *testing.T) {
	pgu, err := randomize.Do[pgtype.UUID]()
	require.NoError(t, err)
	pgu.Valid = true
	u := ToUUID(pgu)
	require.Equal(t, uuid.UUID(pgu.Bytes), u)

	pgu.Valid = false
	u = ToUUID(pgu)
	require.Equal(t, uuid.Nil, u)
}

func TestToUUIDPtr(t *testing.T) {
	pgu, err := randomize.Do[pgtype.UUID]()
	require.NoError(t, err)
	pgu.Valid = true
	u := ToUUIDPtr(pgu)
	require.Equal(t, uuid.UUID(pgu.Bytes), *u)

	pgu.Valid = false
	u = ToUUIDPtr(pgu)
	require.Nil(t, u)
}
