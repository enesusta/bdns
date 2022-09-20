package main

import (
	"github.com/enesusta/bdns/dns"
	"github.com/enesusta/bdns/model"
	"github.com/enesusta/bdns/parser"
	"github.com/enesusta/bdns/udp"
	"io/ioutil"
	"log"
)

var records = map[string]string{
	"amazon.com":         "176.32.103.205",
	"any.google.com":     "10.0.0.1",
	"babel.enesusta.net": "185.199.110.153",
}

func main() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config, err := parser.ParseEntities(content)
	entities := make(map[string]model.DnsEntity)

	for _, each := range config {
		entities[each.Domain] = each
	}

	udpConf := udp.UdpConfiguration{
		Host: "127.0.0.1",
		Port: 8080,
	}

	u := udp.InitializeUdpSocket(&udpConf)
	entity := dns.DnsConfiguration{
		Entities: entities,
	}

	for {
		dns.ServeDns(u, &entity)
	}
}
