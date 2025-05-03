package main

import (
	"context"
	"database/sql"
	"github.com/brianvoe/gofakeit"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	dbDSN = "host=localhost port=54321 dbname=chat user=chat-user password=chat-password sslmode=disable"
)

func main() {
	ctx := context.Background()

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer con.Close(ctx)

	// Вставка записи
	res, err := con.Exec(ctx, "INSERT INTO chat (title) VALUES ($1)", gofakeit.Name())
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}
	log.Printf("inserted %d rows", res.RowsAffected())

	// Выборка записей
	rows, err := con.Query(ctx, "SELECT id, title, created_at, updated_at FROM chat")
	if err != nil {
		log.Fatalf("failed to select chat: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var createdAt time.Time
		var updatedAt sql.NullTime

		err = rows.Scan(&id, &title, &createdAt, &updatedAt)
		if err != nil {
			log.Fatalf("failed to scan chat: %v", err)
		}

		log.Printf("id: %d, title: %s, created_at: %v, updated_at: %v\n", id, title, createdAt, updatedAt)
	}
}
