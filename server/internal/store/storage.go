package store

import (
	"context"
	"database/sql"
	"time"
)

type IPRepo interface {
	SaveIP(ctx context.Context, ip, reversedIP string) error
}
type Storage struct {
	IP IPRepo
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		IP: &IPStore{db}}
}

var (
	QueryBackgroundTimeout = 5 * time.Second
)
