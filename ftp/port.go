package ftp

import (
	"fmt"
	"strconv"
	"strings"
)

func (conn *FTPConn) setDataAddr(args []string) {

	// assume localhost at the moment.
	data := strings.Split(args[0], ",")

	p1, p2 := data[4], data[5]

	p1int, err := strconv.Atoi(p1)
	if err != nil {
		conn.respond(status426)
	}
	p2int, err := strconv.Atoi(p2)
	if err != nil {
		conn.respond(status426)
	}

	p := (p1int * 256) + p2int

	conn.DataAddr = fmt.Sprintf("127.0.0.1:%d", p)

	conn.respond(status200)
}
