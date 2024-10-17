package db

import (
	"context"
	"log"

	"github.com/avila-r/env"

	app "github.com/avila-r/wiredcraft-challenge"
	"github.com/avila-r/wiredcraft-challenge/sql"

	"github.com/jackc/pgx/v5"
)

var (
	Conn = func() *sql.Conn {
		if err := env.Load(app.RootPath); err != nil {
			log.Fatal(err.Error())
		}

		ctx := context.Background()

		conn, err := pgx.Connect(ctx, env.Get("DB_DSN"))

		if err != nil {
			log.Fatal(err.Error())
		}

		return sql.New(conn)
	}()
)
