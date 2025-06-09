package chat

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/beachrockhotel/chat-server/internal/client/db"
	"github.com/beachrockhotel/chat-server/internal/model"
	"github.com/beachrockhotel/chat-server/internal/repository"
	"github.com/beachrockhotel/chat-server/internal/repository/chat/converter"
	modelRepo "github.com/beachrockhotel/chat-server/internal/repository/chat/model"
)

const (
	tableName = "chat"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.ChatInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn).
		Values(info.Name, info.Email, info.Role, info.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "chat_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Chat, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "chat_repository.Get",
		QueryRaw: query,
	}

	var chat modelRepo.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&auth.ID, &auth.Info.Name, &auth.Info.Email, &auth.Info.Role, &auth.CreatedAt, &auth.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToAuthFromRepo(&chat), nil
}
