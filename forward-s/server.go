package main

import (
	"encoding/binary"
	"errors"
	"king-go/net/echo"
	"net"
)

func getByteOrder() binary.ByteOrder {
	if getConfigure().BigEndian {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

type Session struct {
	//是否已經登入
	login bool
	//c net.Conn
}
type Server struct {
	cmds map[uint16]ICommand
}

func (s *Server) GetHeaderSize() int {
	return HeaderSize
}
func (s *Server) GetMessageSize(session echo.Session, b []byte) (int, error) {
	if len(b) != HeaderSize {
		return 0, errors.New("header size not match")
	}

	order := getByteOrder()
	if order.Uint16(b) != HeaderFlag {
		return 0, errors.New("header flag not match")
	}
	n := order.Uint16(b[2:])
	if n < HeaderSize {
		return 0, errors.New("message size not match")
	}

	cmd := order.Uint16(b[4:])
	if _, ok := s.cmds[cmd]; !ok {
		return 0, errors.New("cmd not match")
	}

	return int(n), nil
}
func (s *Server) NewSession(c net.Conn) (session echo.Session, e error) {
	return &Session{}, e
}
func (s *Server) DeleteSession(c net.Conn, session echo.Session) {
}
func (s *Server) Message(c net.Conn, session echo.Session, b []byte) error {
	order := getByteOrder()
	cmd := order.Uint16(b[4:])

	if cmd, ok := s.cmds[cmd]; ok {
		return cmd.Execute(c, session.(Session), b)
	}
	return errors.New("command unknow")
}
