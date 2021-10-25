package model

import (
	"context"
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://unicorn_user:magical_password@localhost:5432/test_events?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Warn().Err(err).Msg("cannot connect to db")
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

func createRandomEvent(t *testing.T) Events {
	arg := Events{
		Name: "Microsoft Azure",
		Time: Date{Day: "11", Month: "03", Year: "2022"},
		Desc: "Description Microsoft Azure event",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestDelereAccount(t *testing.T) {
	newEvent := createRandomEvent(t)

}
