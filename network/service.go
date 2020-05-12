package network

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"
)

// SocketService struct
type SocketService struct {
	onMessage    func(*Session, *Message)
	onUDPMessage func(*UDPConn, *Session, []byte)
	onConnect    func(*Session)
	onDisconnect func(*Session, error)
	sessions     *sync.Map
	hbInterval   time.Duration
	hbTimeout    time.Duration
	laddr        string
	portUdp      string
	status       int
	tListener    net.Listener
	stopCh       chan error
	uListener    *net.UDPConn
	uDone        chan error
	uRcvMessage  chan []byte
	uSendMessage chan []byte
}

// NewSocketService create a new socket service
func NewSocketService(laddr string, _u string) (*SocketService, error) {
	l, err := net.Listen("tcp", laddr)
	if err != nil {
		return nil, err
	}

	uaddr, err := net.ResolveUDPAddr("udp", _u)
	if err != nil {
		return nil, err
	}

	u, err := net.ListenUDP("udp", uaddr)
	if err != nil {
		return nil, err
	}

	s := &SocketService{
		sessions:   &sync.Map{},
		stopCh:     make(chan error),
		hbInterval: 0 * time.Second,
		hbTimeout:  0 * time.Second,
		laddr:      laddr,
		status:     STInited,
		portUdp:    _u,
		tListener:  l,
		uListener:  u,
	}

	return s, nil
}

// RegUDPMessageHandler register UDP message handler
func (s *SocketService) RegUDPMessageHandler(handler func(*UDPConn, *Session, []byte)) {
	s.onUDPMessage = handler
}

// RegMessageHandler register message handler
func (s *SocketService) RegMessageHandler(handler func(*Session, *Message)) {
	s.onMessage = handler
}

// RegConnectHandler register connect handler
func (s *SocketService) RegConnectHandler(handler func(*Session)) {
	s.onConnect = handler
}

// RegDisconnectHandler register disconnect handler
func (s *SocketService) RegDisconnectHandler(handler func(*Session, error)) {
	s.onDisconnect = handler
}

// Serv Start socket service
func (s *SocketService) Serv() {

	s.status = STRunning
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		s.status = STStop
		cancel()
		s.tListener.Close()
		s.uListener.Close()
	}()
	uconn := NewUDPConn(s.uListener, s.portUdp)
	go uconn.udpReadCorroutine(ctx)
	go uconn.udpWriteCorroutine(ctx)
	go s.acceptHandler(ctx, uconn)
	for {
		select {
		case <-s.stopCh:
			return
		}
	}
}

func (s *SocketService) acceptHandler(ctx context.Context, u *UDPConn) {
	for {
		c, err := s.tListener.Accept()
		if err != nil {
			s.stopCh <- err
			return
		}

		go s.connectHandler(ctx, c, u)
	}
}

func (s *SocketService) connectHandler(ctx context.Context, c net.Conn, u *UDPConn) {
	conn := NewConn(c, s.hbInterval, s.hbTimeout)
	session := NewSession(conn)
	s.sessions.Store(session.GetSessionID(), session)

	connctx, cancel := context.WithCancel(ctx)

	defer func() {
		cancel()
		conn.Close()
		s.sessions.Delete(session.GetSessionID())
	}()

	go conn.readCoroutine(connctx)
	go conn.writeCoroutine(connctx)

	if s.onConnect != nil {
		s.onConnect(session)
	}

	for {
		select {
		case err := <-conn.done:

			if s.onDisconnect != nil {
				s.onDisconnect(session, err)
			}
			return

		case msg := <-conn.messageCh:
			if s.onMessage != nil {
				s.onMessage(session, msg)
			}

		case umsg := <-u.messageCh:
			if s.onUDPMessage != nil {
				s.onUDPMessage(u, session, umsg)
			}
		}
	}
}

// GetStatus get socket service status
func (s *SocketService) GetStatus() int {
	return s.status
}

// Stop stop socket service with reason
func (s *SocketService) Stop(reason string) {
	s.stopCh <- errors.New(reason)
}

// SetHeartBeat set heart beat
func (s *SocketService) SetHeartBeat(hbInterval time.Duration, hbTimeout time.Duration) error {
	if s.status == STRunning {
		return errors.New("Can't set heart beat on service running")
	}

	s.hbInterval = hbInterval
	s.hbTimeout = hbTimeout

	return nil
}

// GetConnsCount get connect count
func (s *SocketService) GetConnsCount() int {
	var count int
	s.sessions.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}

// Unicast Unicast with session ID
func (s *SocketService) Unicast(sid string, msg *Message) {
	v, ok := s.sessions.Load(sid)
	if ok {
		session := v.(*Session)
		err := session.GetConn().SendMessage(msg)
		if err != nil {
			return
		}
	}
}

// Broadcast Broadcast to all connections
func (s *SocketService) Broadcast(msg *Message) {
	s.sessions.Range(func(k, v interface{}) bool {
		s := v.(*Session)
		if err := s.GetConn().SendMessage(msg); err != nil {
			// log.Println(err)
		}
		return true
	})
}
