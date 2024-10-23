package ftp

import (
	"fmt"
	"strconv"
	"strings"
)

func (conn *FTPConn) setDataAddr(args []string) {

	data := strings.Split(args[0], ",")

	i1, i2, i3, i4 := data[0], data[1], data[2], data[3]

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

	conn.DataAddr = fmt.Sprintf("%s.%s.%s.%s:%d", i1, i2, i3, i4, p)
	conn.respond(status200)
}
