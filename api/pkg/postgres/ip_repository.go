package postgres

import (
	"context"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/ent/bannedip"
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

func (r *IPRepository) FindAllOrderByCountryAscIPAsc(ctx context.Context) ([]*ent.BannedIP, error) {
	ips, err := r.conn.txClient(ctx).BannedIP.
		Query().
		Order(bannedip.ByCountry(), bannedip.ByIP()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
