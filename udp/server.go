package udp

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type UdpConfiguration struct {
	Host string
	Port int
}

const (
	PROTOCOL    = "udp"
	BUFFER_SIZE = uint16(2048)
)

func InitializeUdpSocket(conf *UdpConfiguration) *net.UDPConn {
	addr := net.UDPAddr{
		Port: conf.Port,
		IP:   net.ParseIP(conf.Host),
	}

	sock, _ := net.ListenUDP(PROTOCOL, &addr)
	return sock
}

func ReceiveFrom(u *net.UDPConn) (net.Addr, *layers.DNS) {
	buffer := make([]byte, BUFFER_SIZE)
	_, addr, _ := u.ReadFrom(buffer)
	tcp := prepareLayer(&addr, buffer)
	return addr, tcp
}

func prepareLayer(addr *net.Addr, buffer []byte) *layers.DNS {
	packet := gopacket.NewPacket(buffer, layers.LayerTypeDNS, gopacket.Default)
	dnsPacket := packet.Layer(layers.LayerTypeDNS)
	tcp, _ := dnsPacket.(*layers.DNS)
	return tcp
}
