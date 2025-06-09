package model

import (
	"database/sql"
	"github.com/beachrockhotel/chat-server/pkg/chat_v1"
	"time"
)

type Chat struct {
	ID        int64
	Title     string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
