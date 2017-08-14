package main

import (
	"net"
)

type ICommand interface {
	Code() uint16
	Execute(c net.Conn, session Session, b []byte) error
}
