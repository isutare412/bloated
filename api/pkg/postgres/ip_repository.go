package postgres

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/ent"
)

type IPRepository struct {
	conn *Connection
}

func NewIPRepository(conn *Connection) *IPRepository {
	return &IPRepository{
		conn: conn,
	}
}

func (r *IPRepository) CreateAll(ctx context.Context, ips []*ent.BannedIP) ([]*ent.BannedIP, error) {
	cli := r.conn.txClient(ctx)

	bulk := make([]*ent.BannedIPCreate, 0, len(ips))
	for _, ip := range ips {
		var country *string
		if ip.Country != "" {
			country = &ip.Country
		}

		bulk = append(bulk, cli.BannedIP.
			Create().
			SetIP(ip.IP).
			SetNillableCountry(country))

	}

	created, err := cli.BannedIP.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *IPRepository) FindAll(ctx context.Context) ([]*ent.BannedIP, error) {
	ips, err := r.conn.txClient(ctx).BannedIP.
		Query().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
