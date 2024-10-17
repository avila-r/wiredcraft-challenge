package db_test

import (
	"testing"
	"time"

	"github.com/avila-r/wiredcraft-challenge/db"
	"github.com/avila-r/wiredcraft-challenge/sql"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	DateOfBirth = func() pgtype.Date {
		var (
			day   = 15
			month = 10
			year  = 2024
		)

		target := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

		return pgtype.Date{
			Time:  target,
			Valid: true,
		}
	}()
)

func Test_Conn(t *testing.T) {
	db := db.Conn

	u, err := db.CreateUser(sql.CreateUserParams{
		Name:        "test-user",
		Dob:         DateOfBirth,
		Description: "test-description",
	})

	if err != nil {
		t.Error(err.Error())
	}

	db.DeleteUser(u.ID)
}
