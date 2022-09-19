package udp

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

const (
	PORT        = 8080
	PROTOCOL    = "udp"
	BUFFER_SIZE = uint16(2048)
)

var IP = net.ParseIP("127.0.0.1")

func InitializeUdpSocket() *net.UDPConn {
	addr := net.UDPAddr{
		Port: PORT,
		IP:   IP,
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
