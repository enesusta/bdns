package model

type DnsEntity struct {
	Domain string `json:"domain"`
	Ip     string `json:"ip"`
	TTL    uint16 `json:"ttl"`
}
