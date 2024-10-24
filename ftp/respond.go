package ftp

import (
	"fmt"
	"log"
)

const (
	status150 = "150 File status okay; about to open data connection."
	status200 = "200 Command okay."
	status220 = "220 Service ready for new user."
	status221 = "221 Service closing control connection."
	status226 = "226 Closing data connection. Requested file action successful."
	status230 = "230 User %s logged in, proceed."
	status425 = "425 Can't open data connection."
	status426 = "426 Connection closed; transfer aborted."
	status501 = "501 Syntax error in parameters or arguments."
	status502 = "502 Command not implemented."
	status504 = "504 Cammand not implemented for that parameter."
	status550 = "550 Requested action not taken. File unavailable."
)

func (c *FTPConn) respond(s string) {
	log.Println(">> " + fmt.Sprint(s, "\n"))
	_, err := fmt.Fprint(c.Conn, s, "\n")
	if err != nil {
		log.Printf("error sending data: %s\n", err)
	}
}
