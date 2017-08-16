package main

import (
	"errors"
	"net"
)

type CommandLogin struct {
}

func (CommandLogin) Code() uint16 {
	return CmdLogin
}
func (CommandLogin) Execute(c net.Conn, session *Session, b []byte) error {
	cnf := getConfigure()
	if cnf.Pwd != string(b) {
		e := errors.New("pwd not match")
		if e != nil {
			logDebug.Println(e, session)
		}
		return e
	}
	session.timer.Stop()
	session.login = true
	logTrace.Println("one login", session)

	return writeData(c, CmdLogin, nil)
}
