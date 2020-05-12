package network

import (
	"net"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

// Session struct
type Session struct {
	sID      string
	uID      string
	conn     *Conn
	uAddr    *net.UDPAddr
	settings map[string]interface{}
}

// NewSession create a new session
func NewSession(conn *Conn) *Session {
	id := uuid.NewV4()
	session := &Session{
		sID:      id.String(),
		uID:      "",
		conn:     conn,
		settings: make(map[string]interface{}),
	}

	return session
}

// GetSessionID get session ID
func (s *Session) GetSessionID() string {
	return s.sID
}

// BindUserID bind a user ID to session
func (s *Session) BindUserID(uid string) {
	s.uID = uid
}

// GetUserID get user ID
func (s *Session) GetUserID() string {
	return s.uID
}

// GetUDPAddr get user ID
func (s *Session) GetUDPAddr() *net.UDPAddr {
	return s.uAddr
}

// SetUDPAddr get user ID
func (s *Session) SetUDPAddr(u *net.UDPAddr) {
	n, err := net.ResolveUDPAddr("udp", u.IP.String()+":"+strconv.Itoa(u.Port))
	if err != nil {
	}
	s.uAddr = n
}

// GetConn get serverplugin.Conn pointer
func (s *Session) GetConn() *Conn {
	return s.conn
}

// SetConn set a serverplugin.Conn to session
func (s *Session) SetConn(conn *Conn) {
	s.conn = conn
}

// GetSetting get setting
func (s *Session) GetSetting(key string) interface{} {

	if v, ok := s.settings[key]; ok {
		return v
	}

	return nil
}

// SetSetting set setting
func (s *Session) SetSetting(key string, value interface{}) {
	s.settings[key] = value
}
