package ftp

import (
	"fmt"
	"strings"
)

func (c *FTPConn) user(n []string) {
	c.respond(fmt.Sprintf(status230, strings.Join(n, " ")))
}
