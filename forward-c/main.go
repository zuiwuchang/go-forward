package main

import (
	"log"
	"net"
)

const (
	ConfigureFile = "forward-c.json"

	HeaderFlag = 1102
	HeaderSize = 6

	CmdLogin   = 1
	CmdConnect = 2
	CmdWrite   = 3
)

func main() {
	//加載配置
	e := initConfigure()
	if e != nil {
		log.Fatalln(e)
	}
	cnf := getConfigure()
	//初始化 日誌
	if cnf.LogLine {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	//初始化 加密模塊
	initCrypto(cnf.Key)

	//創建服務器
	lister, e := net.Listen("tcp", cnf.LAddr)
	if e != nil {
		logFault.Println(e)
		return
	}
	logInfo.Println("work at", cnf.LAddr)
	defer lister.Close()
	for {
		c, e := lister.Accept()
		if e != nil {
			logError.Println(e)
			continue
		}
		go doWork(c)
	}
}
func doWork(c net.Conn) {
	forward, e := NewForward(c)
	if e != nil {
		logError.Println(e)
		return
	}
	forward.Run()
}
