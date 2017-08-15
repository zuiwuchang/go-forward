package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type Configure struct {
	LAddr     string
	Timeout   time.Duration
	Key       string
	Pwd       string
	BigEndian bool
	Service   []Service
	Services  map[uint16]string
	Logs      []string
	Log       map[string]int
	LogLine   bool
}
type Service struct {
	Id   uint16
	Addr string
}

var g_cnf Configure

func getConfigure() *Configure {
	return &g_cnf
}
func initConfigure() error {
	b, e := ioutil.ReadFile(ConfigureFile)
	if e != nil {
		return e
	}
	cnf := getConfigure()
	e = json.Unmarshal(b, cnf)
	if e != nil {
		return e
	}

	cnf.Services = make(map[uint16]string)
	for _, node := range cnf.Service {
		cnf.Services[node.Id] = node.Addr
	}
	cnf.Service = nil

	cnf.Log = make(map[string]int)
	for _, node := range cnf.Logs {
		cnf.Log[node] = 1
	}
	cnf.Logs = nil
	return nil
}
