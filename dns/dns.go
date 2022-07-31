package dns

import (
	"net"

	"github.com/enesusta/bdns/model"
	"github.com/enesusta/bdns/udp"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type DnsConfiguration struct {
	Entities map[string][]model.DnsRecord
}

// for more info: https://datatracker.ietf.org/doc/html/rfc1035
func ServeDns(u *net.UDPConn, d *DnsConfiguration) {
	addr, tcp := udp.ReceiveFrom(u)
	d.serveDNS(u, addr, tcp)
}

func (d *DnsConfiguration) serveDNS(u *net.UDPConn, clientAddr net.Addr, request *layers.DNS) {
	subDomain := request.Questions[0].Name
	// res := make([]layers.DNSResourceRecord, 10)

	if dnsRecords, ok := d.Entities[string(subDomain)]; ok {
		for _, dnsRecord := range dnsRecords {
			// A, _, _ := net.ParseCIDR(dnsRecord.IP + "/24")
			name := request.Questions[0].Name

			r := layers.DNSResourceRecord{}
			r.Type = layers.DNSTypeCNAME
			// r.IP = A
			r.CNAME = []byte(dnsRecord.IP)
			r.Name = []byte(name)
			r.Class = layers.DNSClassIN
			// res = append(res, r)
			setDnsHeaderWithSingleAnswer(request, r)
		}
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{}

	// setDnsHeaderWithSingleAnswer(request, r)
	// setDnsHeaders(request, res)

	err := request.SerializeTo(buf, opts)
	if err != nil {
		panic(err)
	}
	u.WriteTo(buf.Bytes(), clientAddr)
}

func setDnsHeaders(request *layers.DNS, answers []layers.DNSResourceRecord) {
	request.QR = true
	request.ANCount = 1
	request.OpCode = layers.DNSOpCodeNotify
	request.AA = true
	request.Answers = answers

	request.ResponseCode = layers.DNSResponseCodeNoErr
}

func setDnsHeaderWithSingleAnswer(request *layers.DNS, answer layers.DNSResourceRecord) {
	request.QR = true
	request.ANCount = 1
	request.OpCode = layers.DNSOpCodeNotify
	request.AA = true

	request.Answers = append(request.Answers, answer)
	request.ResponseCode = layers.DNSResponseCodeNoErr
}
