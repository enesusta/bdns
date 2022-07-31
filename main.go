package main

import (
	"github.com/enesusta/bdns/dns"
	"github.com/enesusta/bdns/udp"
)

var records = map[string]string{
	"amazon.com":         "176.32.103.205",
	"any.google.com":     "10.0.0.1",
	"babel.enesusta.net": "185.199.110.153",
}

func main() {
	u := udp.InitializeUdpSocket()
	entity := dns.Entity{
		Records: records,
	}

	for {
		dns.ServeDns(u, &entity)
	}
}
