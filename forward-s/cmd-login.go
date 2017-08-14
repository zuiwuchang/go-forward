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
func (CommandLogin) Execute(c net.Conn, session Session, b []byte) error {
	cnf := getConfigure()
	if cnf.Pwd != string(b) {
		return errors.New("pwd not match")
	}
	session.login = true
	return nil
}
