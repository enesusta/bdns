package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/enesusta/bdns/dns"
	"github.com/enesusta/bdns/model"
	"github.com/enesusta/bdns/parser"
	"github.com/enesusta/bdns/udp"
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
	entities := make(map[string][]model.DnsRecord)

	for _, each := range config {
		entities[each.Subdomain] = each.Records
	}

	fmt.Println(entities)

	u := udp.InitializeUdpSocket()
	entity := dns.DnsConfiguration{
		Entities: entities,
	}

	for {
		dns.ServeDns(u, &entity)
	}
}
