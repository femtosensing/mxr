package mxnet

import (
	"fmt"
	"net"
	"time"
)

type Net struct {
	TX_URL string
	RX_URL string

	READ_TIMEOUT int

	NetCon      net.PacketConn
	IsListening bool

	inBuff []byte
}

type Bridge struct {
	Host_if         Net
	Interrogator_if Net
	Terminate       bool
	Show            bool
}

func NewBridge() *Bridge {
	b := Bridge{}
	b.Host_if = *NewNet("255.255.255.255:8001", ":4567")
	b.Interrogator_if = *NewNet("192.168.0.19:4567", ":8001")

	return &b
}

func NewNet(tx string, rx string) *Net {
	n := Net{}
	n.TX_URL = tx
	n.RX_URL = rx
	n.READ_TIMEOUT = 10

	n.inBuff = make([]byte, 1600)

	return &n
}

func (b *Bridge) Start() error {

	var err error

	b.Host_if.NetCon, err = net.ListenPacket("udp", b.Host_if.RX_URL)
	if err != nil {
		b.Host_if.IsListening = false
		return err
	}
	b.Host_if.IsListening = true

	fmt.Println("Listener Started, Port:", b.Host_if.RX_URL)

	b.Interrogator_if.NetCon, err = net.ListenPacket("udp", b.Interrogator_if.RX_URL)
	if err != nil {
		b.Interrogator_if.IsListening = false
		return err
	}
	b.Interrogator_if.IsListening = true

	fmt.Println("Listener Started, Port:", b.Interrogator_if.RX_URL)

	go b.HostLoop()
	go b.InterroLoop()

	for !b.Terminate {
		time.Sleep(time.Second)
	}

	return nil
}

func (b *Bridge) HostLoop() {

	if b.Host_if.NetCon == nil {
		fmt.Println("Listener not started, aborting Rx Loop")
		return
	}

	b.Terminate = false
	buf := make([]byte, 1500)

	for !b.Terminate {

		k, addr, err := b.Host_if.NetCon.ReadFrom(buf)

		if err == nil {
			bf := buf[:k]
			b.host_rx_handler(addr, bf)
		}
	}

}

func (b *Bridge) InterroLoop() {

	if b.Interrogator_if.NetCon == nil {
		fmt.Println("Listener not started, aborting Rx Loop")
		return
	}

	b.Terminate = false
	buf := make([]byte, 1500)

	for !b.Terminate {

		k, addr, err := b.Interrogator_if.NetCon.ReadFrom(buf)

		aa := addr.String()

		//Prvent Infinte boradlas loop back
		if aa != b.Interrogator_if.TX_URL {
			//fmt.Println("REJECT", aa)
			continue
		}

		if err == nil {
			bf := buf[:k]
			b.inter_rx_handler(addr, bf)
		}
	}

}

func (b *Bridge) inter_rx_handler(addr net.Addr, rx []byte) {

	nRes := len(rx)
	if nRes == 0 {
		return
	}
	b.logPacket(addr.String(), b.Host_if.TX_URL, len(rx))
	go tx(&b.Host_if, rx)

}

func (b *Bridge) host_rx_handler(addr net.Addr, rx []byte) {

	nRes := len(rx)
	if nRes == 0 {
		return
	}

	b.logPacket(addr.String(), b.Interrogator_if.TX_URL, len(rx))
	go tx(&b.Interrogator_if, rx)

}

func (b *Bridge) logPacket(src, dst string, len int) {
	if b.Show {
		t := time.Now()
		ts := t.Format("2006-01-02 15:04:05")
		fmt.Printf("%s       %s \t >> \t %s \t len %d\n", ts, src, dst, len)
	}
}

func tx(n *Net, buf []byte) {
	addr, _ := net.ResolveUDPAddr("udp", n.TX_URL)
	//fmt.Println(addr, len(buf))
	n.NetCon.WriteTo(buf, addr)
}
