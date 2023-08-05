package http

import (
	"time"

	"github.com/isutare412/bloated/api/pkg/core/ent"
)

type bannedIP struct {
	ID         int       `json:"id"`
	IP         string    `json:"ip"`
	Country    string    `json:"country,omitempty"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (b *bannedIP) fromEntity(ip *ent.BannedIP) {
	b.ID = ip.ID
	b.IP = ip.IP
	b.Country = ip.Country
	b.CreateTime = ip.CreateTime
	b.UpdateTime = ip.UpdateTime
}

type listBannedIPsResponse struct {
	BannedIPs []bannedIP `json:"bannedIps"`
}

func (resp *listBannedIPsResponse) fromEntities(ips []*ent.BannedIP) {
	resp.BannedIPs = make([]bannedIP, len(ips))
	for i, ip := range ips {
		resp.BannedIPs[i].fromEntity(ip)
	}
}
