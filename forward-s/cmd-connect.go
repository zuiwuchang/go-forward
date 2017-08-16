package main

import (
	"errors"
	"net"
)

type CommandConnect struct {
}

func (CommandConnect) Code() uint16 {
	return CmdConnect
}
func (CommandConnect) Execute(c net.Conn, session *Session, b []byte) error {
	if !session.login {
		e := errors.New("not loin")
		logDebug.Println(e, session)
		return e
	}
	if session.s != nil {
		e := errors.New("already connect")
		logDebug.Println(e, session)
		return e
	}

	order := getByteOrder()
	id := order.Uint16(b)
	cnf := getConfigure()
	if addr, ok := cnf.Services[id]; ok {
		e := session.Init(addr)
		if e == nil {
			logTrace.Println("one connect", session)
		} else {
			logDebug.Println(e, session)
		}
		return e
	}

	e := errors.New("no service")
	logDebug.Println(e, session)
	return e
}
