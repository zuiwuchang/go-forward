package main

import (
	"errors"
	"net"
)

type CommandWrite struct {
}

func (CommandWrite) Code() uint16 {
	return CmdWrite
}
func (CommandWrite) Execute(c net.Conn, session Session, b []byte) error {
	s := session.s
	if s == nil {
		e := errors.New(" not connect")
		logDebug.Println(e, session)
		return e
	}
	_, e := s.Write(b)
	if e != nil {
		logDebug.Println(e, session)
	}
	return e
}
