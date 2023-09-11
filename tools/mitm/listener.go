package main

import (
	"net"
)

type Listener struct {
	connChan chan net.Conn
}

func NewListener() (l *Listener, err error) {
	l = &Listener{
		connChan: make(chan net.Conn),
	}
	return
}

func (l *Listener) Accept() (net.Conn, error) {
	select {
	case c := <-l.connChan:
		return c, nil
	}
}
func (l *Listener) Close() error {
	close(l.connChan)
	return nil
}
func (l *Listener) Addr() net.Addr {
	return nil
}
func (l *Listener) AddConn(conn net.Conn) {
	l.connChan <- conn
}
