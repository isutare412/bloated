package postgres

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/fx"

	"github.com/isutare412/bloated/api/pkg/core/ent"
)

type Client struct {
	cli *ent.Client
}

func NewClient(lc fx.Lifecycle, cfg ClientConfig) (*Client, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	entCli, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening PostgreSQL connection: %w", err)
	}

	client := &Client{
		cli: entCli,
	}

	lc.Append(fx.Hook{
		OnStart: client.Start,
		OnStop:  client.Stop,
	})
	return client, nil
}

func (c *Client) Start(ctx context.Context) error {
	return c.RunDevCode(ctx)
}

func (c *Client) Stop(ctx context.Context) error {
	return c.cli.Close()
}

func (c *Client) RunDevCode(ctx context.Context) error {
	todo, err := c.cli.Todo.Create().
		SetUserID("redshore").
		SetTask("doing fx test").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating todo: %w", err)
	}

	fmt.Printf("todo: %+v\n", todo)

	return nil
}
