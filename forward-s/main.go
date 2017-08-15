package main

import (
	"king-go/net/echo"
	"log"
	"time"
)

const (
	ConfigureFile = "forward-s.json"

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

	//創建服務器 模板
	cmds := make(map[uint16]ICommand)
	cmds[CmdLogin] = CommandLogin{}
	cmds[CmdConnect] = CommandConnect{}
	cmds[CmdWrite] = CommandWrite{}
	server := &Server{cmds: cmds}

	//創建服務器
	s, e := echo.NewServer(cnf.LAddr, cnf.Timeout*time.Second, server)
	if e != nil {
		logFault.Println(e)
		return
	}
	logInfo.Println("work at", cnf.LAddr)

	//運行服務器
	s.Run()

	//等待服務器 停止
	s.Wait()
}
