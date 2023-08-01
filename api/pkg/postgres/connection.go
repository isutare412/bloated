package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/fx"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/transaction"
)

type Connection struct {
	client *ent.Client
}

func NewConnection(lc fx.Lifecycle, cfg ClientConfig) (*Connection, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening PostgreSQL connection: %w", err)
	}

	conn := &Connection{
		client: client,
	}

	lc.Append(fx.Hook{
		OnStop: conn.Stop,
	})
	return conn, nil
}

func (c *Connection) Stop(ctx context.Context) error {
	return c.client.Close()
}

func (c *Connection) WithTx(ctx context.Context) (transaction.Context, error) {
	tx, err := c.client.Tx(ctx)
	if err != nil {
		return transaction.Context{}, fmt.Errorf("starting tx: %w", err)
	}

	return transaction.Context{
		Ctx:      context.WithValue(ctx, ctxKeyTx{}, tx),
		Commit:   tx.Commit,
		Rollback: tx.Rollback,
	}, nil
}

func (c *Connection) WithTxOption(ctx context.Context, opts *sql.TxOptions) (transaction.Context, error) {
	tx, err := c.client.BeginTx(ctx, opts)
	if err != nil {
		return transaction.Context{}, fmt.Errorf("starting tx: %w", err)
	}

	return transaction.Context{
		Ctx:      context.WithValue(ctx, ctxKeyTx{}, tx),
		Commit:   tx.Commit,
		Rollback: tx.Rollback,
	}, nil
}

func (c *Connection) txClient(ctx context.Context) *ent.Client {
	tx, ok := ctx.Value(ctxKeyTx{}).(*ent.Tx)
	if !ok {
		return c.client
	}
	return tx.Client()
}

type ctxKeyTx struct{}
