package network

import (
	"context"
	"fmt"
	"net"
	"strconv"
)

// Conn wrap net.Conn
type UDPConn struct {
	udpConn   *net.UDPConn
	port      string
	addr      *net.UDPAddr
	tAddr     *net.UDPAddr
	messageCh chan []byte
	sendCh    chan []byte
	done      chan error
}

// NewConn create new conn
func NewUDPConn(c *net.UDPConn, u string) *UDPConn {
	conn := &UDPConn{
		udpConn:   c,
		port:      u,
		sendCh:    make(chan []byte, 1024),
		messageCh: make(chan []byte, 1024),
		done:      make(chan error),
	}
	return conn
}

// Get UDPADDR
func (c *UDPConn) GetAddr() *net.UDPAddr {
	return c.addr
}

// Set TUDPADDR
func (c *UDPConn) SetTAddr(i *net.UDPAddr) {
	n, err := net.ResolveUDPAddr("udp", i.IP.String()+":"+strconv.Itoa(i.Port))
	if err != nil {
	}
	c.tAddr = n
}

// Get TUDPADDR
func (c *UDPConn) GetTAddr() *net.UDPAddr {
	return c.tAddr
}

// Close close connection
func (c *UDPConn) Close() {
	c.udpConn.Close()
}

// SendMessage send message
func (c *UDPConn) SendMessage(msg []byte) error {
	c.sendCh <- msg
	return nil
}

func (c *UDPConn) udpReadCorroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			buf := make([]byte, 1024)
			n, conn, err := c.udpConn.ReadFromUDP(buf)
			if err != nil {
				c.done <- err
				continue
			}
			if conn == nil {
				c.done <- err
				continue
			}
			c.addr = conn
			c.messageCh <- buf[:n]
		}
	}
}

func (c *UDPConn) udpWriteCorroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case pkt := <-c.sendCh:

			if pkt == nil {
				continue
			}
			/*
				parts := strings.Split(c.session.GetUserID(), ":")
				_a := parts[0] + c.port
				addr, err := net.ResolveUDPAddr("udp", _a)
				if err != nil {
					continue
				}
			*/
			_, err := c.udpConn.WriteToUDP(pkt, c.tAddr)
			if err != nil {
				fmt.Printf("Couldn't send response %v", err)
				c.done <- err
			}
		}
	}
}
