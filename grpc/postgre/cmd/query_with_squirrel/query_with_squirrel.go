package main

import (
	"context"
	"database/sql"
	"github.com/brianvoe/gofakeit"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	dbDSN = "host=localhost port=54321 dbname=chat user=chat-user password=chat-password sslmode=disable"
)

func main() {
	ctx := context.Background()

	// Подключение к пулу
	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// === INSERT ===
	insertQuery, insertArgs, err := sq.Insert("chat").
		PlaceholderFormat(sq.Dollar).
		Columns("title").
		Values(gofakeit.Company()).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		log.Fatalf("failed to build insert query: %v", err)
	}

	var chatID int
	err = pool.QueryRow(ctx, insertQuery, insertArgs...).Scan(&chatID)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}
	log.Printf("inserted chat with id: %d", chatID)

	// === SELECT (list) ===
	selectQuery, selectArgs, err := sq.Select("id", "title", "created_at", "updated_at").
		From("chat").
		PlaceholderFormat(sq.Dollar).
		OrderBy("id ASC").
		Limit(10).
		ToSql()
	if err != nil {
		log.Fatalf("failed to build select query: %v", err)
	}

	rows, err := pool.Query(ctx, selectQuery, selectArgs...)
	if err != nil {
		log.Fatalf("failed to select chats: %v", err)
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

		log.Printf("id: %d, title: %s, created_at: %v, updated_at: %v", id, title, createdAt, updatedAt)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("error while iterating over chats: %v", err)
	}

	// === UPDATE ===
	updateQuery, updateArgs, err := sq.Update("chat").
		PlaceholderFormat(sq.Dollar).
		Set("title", gofakeit.Company()).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": chatID}).
		ToSql()
	if err != nil {
		log.Fatalf("failed to build update query: %v", err)
	}

	updateRes, err := pool.Exec(ctx, updateQuery, updateArgs...)
	if err != nil {
		log.Fatalf("failed to update chat: %v", err)
	}
	log.Printf("updated %d rows", updateRes.RowsAffected())

	// === SELECT (by id) ===
	selectOneQuery, selectOneArgs, err := sq.Select("id", "title", "created_at", "updated_at").
		From("chat").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": chatID}).
		Limit(1).
		ToSql()
	if err != nil {
		log.Fatalf("failed to build select-one query: %v", err)
	}

	var title string
	var createdAt time.Time
	var updatedAt sql.NullTime

	err = pool.QueryRow(ctx, selectOneQuery, selectOneArgs...).Scan(&chatID, &title, &createdAt, &updatedAt)
	if err != nil {
		log.Fatalf("failed to fetch updated chat: %v", err)
	}

	log.Printf("id: %d, title: %s, created_at: %v, updated_at: %v", chatID, title, createdAt, updatedAt)
}
