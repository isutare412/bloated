export interface BannedIp {
  id: number
  ip: string
  country?: string
  createTime: Date
  updateTime: Date
}

export interface ListBannedIpsResponse {
  bannedIps: BannedIp[]
}
