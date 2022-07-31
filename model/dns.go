package model

type DnsEntity struct {
	Subdomain string      `json:"sub_domain"`
	Records   []DnsRecord `json:"records`
}

type DnsRecord struct {
	TTL  int    `json:"ttl"`
	Type string `json:"type"`
	IP   string `json:"ip"`
}
