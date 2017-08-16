package main

import (
	"encoding/binary"
	"errors"
)

func getByteOrder() binary.ByteOrder {
	if getConfigure().BigEndian {
		return binary.BigEndian
	}
	return binary.LittleEndian
}

//默認的 模板實現
type ClientTemplate struct {
	cmds map[uint16]int
}

func (c *ClientTemplate) GetHeaderSize() int {
	return HeaderSize
}
func (c *ClientTemplate) GetMessageSize(b []byte) (int, error) {
	if len(b) != HeaderSize {
		e := errors.New("header size not match")
		logDebug.Println(e)
		return 0, e
	}

	order := getByteOrder()
	if order.Uint16(b) != HeaderFlag {
		e := errors.New("header flag not match")
		logDebug.Println(e)
		return 0, e
	}
	n := order.Uint16(b[2:])
	if n < HeaderSize {
		e := errors.New("message size not match")
		logDebug.Println(e)
		return 0, e
	}

	cmd := order.Uint16(b[4:])
	if _, ok := c.cmds[cmd]; !ok {
		e := errors.New("cmd not match")
		logDebug.Println(e)
		return 0, e
	}

	return int(n), nil
}
