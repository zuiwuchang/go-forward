package main

import (
	"net"
)

type CommandWrite struct {
}

func (CommandWrite) Code() uint16 {
	return CmdWrite
}
func (CommandWrite) Execute(c net.Conn, session Session, b []byte) error {
	return nil
}
