package main

import (
	"net"
	"time"
)

type Session struct {
	//必須在 2分鐘 內完成 登入
	timer *time.Timer
	//是否已經登入
	login bool

	//與客戶的連接
	c net.Conn
	//與服務器的連接
	s net.Conn
}

func (s *Session) Init(addr string) error {
	c, e := net.Dial("tcp", addr)
	if e != nil {
		return e
	}

	s.s = c
	if e = writeData(s.c, CmdConnect, nil); e != nil {
		return e
	}

	go s.read()

	return nil
}
func (s *Session) Close() {
	if s.s != nil {
		s.s.Close()
		s.s = nil
	}
}
func (s *Session) read() {
	b := make([]byte, 1024)
	c := s.s
	for {
		n, e := c.Read(b)
		if e != nil {
			s.c.Close()
			return
		}

		if e := writeData(s.c, CmdWrite, b[:n]); e != nil {
			s.c.Close()
			return
		}
	}
}
