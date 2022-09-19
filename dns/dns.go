package dns

import (
	"net"

	"github.com/enesusta/bdns/model"
	"github.com/enesusta/bdns/udp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type DnsConfiguration struct {
	Entities map[string]model.DnsEntity
}

// for more info: https://datatracker.ietf.org/doc/html/rfc1035
func ServeDns(u *net.UDPConn, d *DnsConfiguration) {
	addr, tcp := udp.ReceiveFrom(u)
	d.serveDNS(u, addr, tcp)
}

func (d *DnsConfiguration) serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	domain := string(request.Questions[0].Name)

	r := layers.DNSResourceRecord{}
	r.Type = layers.DNSTypeA
	r.Class = layers.DNSClassIN

	found := d.Entities[domain].Ip
	ip, _, _ := net.ParseCIDR(found + "/24")

	r.IP = ip
	r.Name = []byte(domain)
	r.TTL = 3600

	setDnsHeaderWithSingleAnswer(request, r)

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}

	err := request.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}

func setDnsHeaderWithSingleAnswer(request *layers.DNS, answer layers.DNSResourceRecord) {
	request.QR = true
	request.ANCount = 1
	request.OpCode = layers.DNSOpCodeNotify
	request.AA = true

	request.Answers = append(request.Answers, answer)
	request.ResponseCode = layers.DNSResponseCodeNoErr
}
