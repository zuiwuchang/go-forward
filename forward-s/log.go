package main

import (
	"bytes"
	"fmt"
	"log"
	//"os"
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

/*
var logTrace *log.Logger = log.New(os.Stdout, "[TRACE]", log.Lshortfile|log.LstdFlags)
var logDebug *log.Logger = log.New(os.Stdout, "[DEBUG]", log.Lshortfile|log.LstdFlags)
var logInfo *log.Logger = log.New(os.Stdout, "[INFO]", log.Lshortfile|log.LstdFlags)
var logError *log.Logger = log.New(os.Stdout, "[ERROR]", log.Lshortfile|log.LstdFlags)
var logFault *log.Logger = log.New(os.Stdout, "[FAULT]", log.Lshortfile|log.LstdFlags)
*/
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
