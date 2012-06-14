package tools

import (
	"errors"
	"net"
	"time"
)

var ErrNoDeadlines = errors.New("TimoutConn does not support deadlines")

var _ net.Conn = new(TimeoutConn)

type TimeoutConn struct {
	net.Conn
	Timeout time.Duration
}

func (t *TimeoutConn) update() {
	if t.Timeout == 0 {
		t.Conn.SetDeadline(time.Time{})
		return
	}
	t.Conn.SetDeadline(time.Now().Add(t.Timeout))
}

func (t *TimeoutConn) Read(b []byte) (int, error) {
	t.update()
	return t.Read(b)
}

func (t *TimeoutConn) Write(b []byte) (int, error) {
	t.update()
	return t.Write(b)
}

func (t *TimeoutConn) SetDeadline(_ time.Time) error { return ErrNoDeadlines }
func (t *TimeoutConn) SetWriteDeadline(_ time.Time) error { return ErrNoDeadlines }
func (t *TimeoutConn) SetReadDeadline(_ time.Time) error { return ErrNoDeadlines }
