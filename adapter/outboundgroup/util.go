package outboundgroup

import (
	"fmt"
	"net"
	"time"

	C "github.com/opusb/clash-dev/constant"
)

func addrToMetadata(rawAddress string) (addr *C.Metadata, err error) {
	host, port, err := net.SplitHostPort(rawAddress)
	if err != nil {
		err = fmt.Errorf("addrToMetadata failed: %w", err)
		return
	}

	ip := net.ParseIP(host)
	if ip == nil {
		addr = &C.Metadata{
			Host:    host,
			DstIP:   nil,
			DstPort: port,
		}
		return
	} else if ip4 := ip.To4(); ip4 != nil {
		addr = &C.Metadata{
			Host:    "",
			DstIP:   ip4,
			DstPort: port,
		}
		return
	}

	addr = &C.Metadata{
		Host:    "",
		DstIP:   ip,
		DstPort: port,
	}
	return
}

func tcpKeepAlive(c net.Conn) {
	if tcp, ok := c.(*net.TCPConn); ok {
		tcp.SetKeepAlive(true)
		tcp.SetKeepAlivePeriod(30 * time.Second)
	}
}
