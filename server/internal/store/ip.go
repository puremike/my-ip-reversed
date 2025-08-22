package store

import (
	"context"
	"database/sql"
)

// type IPModel struct {
// 	ID               string `json:"id"`
// 	OriginIP         string `json:"origin_ip"`
// 	OriginReversedIP string `json:"origin_reversed_ip"`
// 	CreatedAt        string `json:"created_at"`
// }

type IPStore struct {
	db *sql.DB
}

func (s *IPStore) SaveIP(ctx context.Context, ip, reversedIP string) error {
	ctx, cancel := context.WithTimeout(ctx, QueryBackgroundTimeout)
	defer cancel()

	query := `INSERT INTO ips (origin_ip, origin_reversed_ip) VALUES ($1, $2)`

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, query, ip, reversedIP); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
