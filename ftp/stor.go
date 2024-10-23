package ftp

import (
	"bufio"
	"log"
	"net"
	"os"
)

func (c *FTPConn) stor(args []string) {

	filename := args[0]

	f, err := os.Create(filename)
	if err != nil {
		c.respond(status426)
	}
	defer f.Close()

	c.respond(status150)

	d, err := net.Dial("tcp", c.DataAddr)
	if err != nil {
		c.respond("Cannot open DTP")
	}
	defer d.Close()

	buf := make([]byte, 512)
	r := bufio.NewReader(d)
	i, err := r.Read(buf)
	if err != nil {
		c.respond(status426)
	}

	log.Printf("Reading %d bytes\n", i)

	w := bufio.NewWriter(f)
	i, err = w.Write(buf)
	w.Flush()
	if err != nil {
		c.respond(status426)
	}
	log.Printf("Writing %d bytes\n", i)

	c.respond(status226)
}
