package main

import (
	"errors"
	"king-go/net/echo"
	"net"
)

type Forward struct {
	//到客戶的連接
	cc net.Conn
	//到服務器的連接
	cs echo.IClient
}

func NewForward(c net.Conn) (*Forward, error) {
	cnf := getConfigure()
	//創建 服務器 模板
	cmds := make(map[uint16]int)
	cmds[CmdLogin] = 1
	cmds[CmdConnect] = 1
	cmds[CmdWrite] = 1

	//連接服務器
	cs, e := echo.NewClient(cnf.Addr, &ClientTemplate{cmds: cmds})
	if e != nil {
		return nil, e
	}
	logTrace.Println("connect ok")
	forward := &Forward{cc: c, cs: cs}

	//登入
	e = forward.login()
	if e != nil {
		return nil, e
	}
	logTrace.Println("login ok")
	//連接 服務
	e = forward.connect()
	if e != nil {
		return nil, e
	}
	logTrace.Println("connect ok")
	return forward, nil
}
func (f *Forward) connect() error {
	cnf := getConfigure()
	c := f.cs
	order := getByteOrder()

	b := make([]byte, 2)
	order.PutUint16(b, cnf.Service)

	e := writeData(c, CmdConnect, b)
	if e != nil {
		return e
	}
	b, e = c.GetMessage(0)
	if e != nil {
		return e
	}
	if order.Uint16(b[4:]) != CmdConnect {
		return errors.New("bad rs cmd CmdConnect")
	}
	if len(b) != HeaderSize {
		return errors.New("bad rs len CmdConnect")
	}
	return nil
}
func (f *Forward) login() error {
	cnf := getConfigure()
	c := f.cs
	order := getByteOrder()

	b := []byte(cnf.Pwd)
	e := writeData(c, CmdLogin, b)
	if e != nil {
		return e
	}
	b, e = c.GetMessage(0)
	if e != nil {
		return e
	}

	if order.Uint16(b[4:]) != CmdLogin {
		return errors.New("bad rs cmd CmdLogin")
	}
	if len(b) != HeaderSize {
		return errors.New("bad rs len CmdLogin")
	}
	return nil
}
func (f *Forward) Run() {
	ch := make(chan []byte, 5)
	chExit := make(chan int, 1)
	go func() {
		cs := f.cs
		defer cs.Close()
		for {
			select {
			case <-chExit:
				return
			case <-ch:

			}
		}
	}()
	go func() {
		order := getByteOrder()
		cs := f.cs
		cc := f.cc
		for {
			b, e := cs.GetMessage(0)
			if e != nil {
				logError.Println(e)
				break
			}
			if order.Uint16(b[4:]) != CmdWrite {
				logError.Println("bad rs CmdWrite")
				break
			}
			b = b[HeaderSize:]
			b, e = Decryption(b)
			if e != nil {
				logError.Println(e)
				break
			}
			_, e = cc.Write(b)
			if e != nil {
				logError.Println(e)
				break
			}
		}
		cs.Close()
	}()
	defer func() {
		chExit <- 1
	}()
	b := make([]byte, 1024)
	c := f.cc
	for {
		n, e := c.Read(b)
		if e != nil {
			break
		}
		e = writeData(c, CmdWrite, b[:n])
		if e != nil {
			break
		}
	}
	c.Close()
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
