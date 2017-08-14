package main

import (
	//"errors"
	"net"
)

type CommandConnect struct {
}

func (CommandConnect) Code() uint16 {
	return CmdConnect
}
func (CommandConnect) Execute(c net.Conn, session Session, b []byte) error {

	return nil
}
