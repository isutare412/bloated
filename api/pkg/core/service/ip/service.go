package ip

import (
	"context"
	"fmt"

	"github.com/isutare412/bloated/api/pkg/core/ent"
	"github.com/isutare412/bloated/api/pkg/core/port"
)

type Service struct {
	txMgr  port.TransactionManager
	ipRepo port.IPRepository
}

func NewService(txMgr port.TransactionManager, ipRepo port.IPRepository) *Service {
	return &Service{
		txMgr:  txMgr,
		ipRepo: ipRepo,
	}
}

func (s *Service) AllBannedIPs(ctx context.Context) ([]*ent.BannedIP, error) {
	ips, err := s.ipRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("finding all IPs: %w", err)
	}

	return ips, nil
}
