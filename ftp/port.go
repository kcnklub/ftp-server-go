package ftp

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (conn *FTPConn) connect(args []string) error {

	// assume localhost at the moment.
	data := strings.Split(args[0], ",")

	p1, p2 := data[4], data[5]

	p1int, err := strconv.Atoi(p1)
	if err != nil {
		return err
	}
	p2int, err := strconv.Atoi(p2)
	if err != nil {
		return err
	}

	p := (p1int * 256) + p2int

	_, err = net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", p))
	if err != nil {
		return err
	}

	conn.respond(status200)

	return nil
}
