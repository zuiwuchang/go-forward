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
	Services  map[int64]string
}
type Service struct {
	Id   int64
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

	cnf.Services = make(map[int64]string)
	for _, node := range cnf.Service {
		cnf.Services[node.Id] = node.Addr
	}
	cnf.Service = nil
	return nil
}
