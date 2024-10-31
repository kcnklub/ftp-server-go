package ftp

import (
	"log"
	"path"
)

func (c *FTPConn) cwd(arg []string) {

	relativePath := arg[0]

	c.Root = path.Join(c.Root, relativePath)
	log.Printf("New Path Dir is %s\n", c.Root)
	c.respond(status200)
}
