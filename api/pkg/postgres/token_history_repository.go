package postgres

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/ent"
)

type TokenHistoryRepository struct {
	conn *Connection
}

func NewTokenHistoryRepository(conn *Connection) *TokenHistoryRepository {
	return &TokenHistoryRepository{
		conn: conn,
	}
}

func (r *TokenHistoryRepository) CreateTokenHistory(
	ctx context.Context,
	token *ent.TokenHistory,
) (*ent.TokenHistory, error) {
	created, err := r.conn.txClient(ctx).TokenHistory.
		Create().
		SetEmail(token.Email).
		SetUserName(token.UserName).
		SetIssuedFrom(token.IssuedFrom).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return created, nil
}
