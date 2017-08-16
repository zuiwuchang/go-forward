package main

import (
	"bytes"
	"fmt"
	"log"
)

type ILog interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}
type Log struct {
	flag string
}

var logTrace *Log = &Log{flag: "TRACE"}
var logDebug *Log = &Log{flag: "DEBUG"}
var logInfo *Log = &Log{flag: "INFO"}
var logError *Log = &Log{flag: "ERROR"}
var logFault *Log = &Log{flag: "FAULT"}

func (l *Log) Printf(format string, v ...interface{}) {
	cnf := getConfigure()
	if _, ok := cnf.Log[l.flag]; ok {
		flag := fmt.Sprintf("[%v] ", l.flag)
		log.Printf(flag+format, v...)
	}
}
func (l *Log) Println(v ...interface{}) {
	cnf := getConfigure()
	if _, ok := cnf.Log[l.flag]; ok {
		var buf bytes.Buffer
		flag := fmt.Sprintf("[%v] ", l.flag)
		buf.WriteString(flag)
		for i, node := range v {
			if i != 0 {
				buf.WriteString(" ")

			}
			buf.WriteString(fmt.Sprint(node))
		}

		log.Println(buf.String())
	}
}
