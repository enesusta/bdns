package dns

import (
	"net"

	"github.com/enesusta/bdns/udp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Entity struct {
	Records map[string]string
}

// for more info: https://datatracker.ietf.org/doc/html/rfc1035
func ServeDns(u *net.UDPConn, d *Entity) {
	addr, tcp := udp.ReceiveFrom(u)
	d.serveDNS(u, addr, tcp)
}

func (d *Entity) serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	var record layers.DNSResourceRecord
	record.Type = layers.DNSTypeA
	var ip string
	var err error
	var ok bool
	ip, ok = d.Records[string(request.Questions[0].Name)]

	if !ok {
		ip = "0.0.0.0"
	}

	a, _, _ := net.ParseCIDR(ip + "/24")
	record.Type = layers.DNSTypeA
	record.IP = a
	record.Name = []byte(request.Questions[0].Name)
	record.Class = layers.DNSClassIN

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}

	setDnsHeaders(request, record)

	err = request.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}

func setDnsHeaders(request *layers.DNS, answer layers.DNSResourceRecord) {
	request.QR = true
	request.ANCount = 1
	request.OpCode = layers.DNSOpCodeNotify
	request.AA = true
	request.Answers = append(request.Answers, answer)
	request.ResponseCode = layers.DNSResponseCodeNoErr
}
