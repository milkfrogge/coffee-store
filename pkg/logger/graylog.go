package logger

import (
	"net"
)

/*
 * @project coffee-store
 * @author nick
 */

type GrayLogLogger struct {
	conn net.Conn
}

func NewGrayLogLogger(host string) (*GrayLogLogger, error) {

	c, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return &GrayLogLogger{conn: c}, nil
}

func (g *GrayLogLogger) Write(b []byte) (int, error) {
	n, err := g.conn.Write(b)
	return n, err
}
