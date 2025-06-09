package model

import (
	"database/sql"
	"github.com/beachrockhotel/auth/pkg/auth_v1"
	"time"
)

type Chat struct {
	ID        int64
	Info      AuthInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type ChatInfo struct {
	Name     string
	Email    string
	Role     auth_v1.Role
	Password string
}
