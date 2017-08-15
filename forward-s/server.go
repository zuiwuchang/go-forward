package main

import (
	"encoding/binary"
	"errors"
	"king-go/net/echo"
	"net"
	"time"
)

func getByteOrder() binary.ByteOrder {
	if getConfigure().BigEndian {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

type Server struct {
	cmds map[uint16]ICommand
}

func (s *Server) GetHeaderSize() int {
	return HeaderSize
}
func (s *Server) GetMessageSize(session echo.Session, b []byte) (int, error) {
	if len(b) != HeaderSize {
		e := errors.New("header size not match")
		logDebug.Println(e, session)
		return 0, e
	}

	order := getByteOrder()
	if order.Uint16(b) != HeaderFlag {
		e := errors.New("header flag not match")
		logDebug.Println(e, session)
		return 0, e
	}
	n := order.Uint16(b[2:])
	if n < HeaderSize {
		e := errors.New("message size not match")
		logDebug.Println(e, session)
		return 0, e
	}

	cmd := order.Uint16(b[4:])
	if _, ok := s.cmds[cmd]; !ok {
		e := errors.New("cmd not match")
		logDebug.Println(e, session)
		return 0, e
	}

	return int(n), nil
}
func (s *Server) NewSession(c net.Conn) (echo.Session, error) {
	session := &Session{c: c}
	logTrace.Println("one int :", session)
	session.timer = time.AfterFunc(time.Minute*2, func() {
		logTrace.Println("login timeout :", session)
		c.Close()
	})
	return session, nil
}
func (s *Server) DeleteSession(c net.Conn, session echo.Session) {
	logTrace.Println("one out :", session)
	s0 := session.(*Session)
	s0.Close()
}
func (s *Server) Message(c net.Conn, session echo.Session, b []byte) error {
	order := getByteOrder()
	cmd := order.Uint16(b[4:])

	if cmd, ok := s.cmds[cmd]; ok {
		var e error
		b = b[HeaderSize:]
		b, e = Decryption(b)
		if e != nil {
			return e
		}
		return cmd.Execute(c, session.(Session), b)
	}
	e := errors.New("command unknow")
	logDebug.Println(e, session)
	return e
}

func writeData(c net.Conn, cmd uint16, b []byte) error {
	//加密數據
	var e error
	if b != nil {
		b, e = Encryption(b)
		if e != nil {
			return e
		}
	}
	//創建 緩衝區
	n := len(b) + HeaderSize
	data := make([]byte, n)

	//格式化
	order := getByteOrder()
	order.PutUint16(data, HeaderFlag)
	order.PutUint16(data[2:], uint16(n))
	order.PutUint16(data[4:], cmd)

	copy(data[HeaderSize:], b)

	_, e = c.Write(data)
	return e
}
