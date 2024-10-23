package ftp

import (
	"bufio"
	"log"
	"strings"
)

func Serve(conn FTPConn) {
	conn.respond(status220)
	scanner := bufio.NewScanner(conn.Conn)
	for scanner.Scan() {
		input := scanner.Text()
		log.Println("<< " + input)

		split := strings.Split(input, " ")
		command, args := split[0], split[1:]

		switch command {
		case "USER":
			conn.user(args)
		case "QUIT":
			conn.respond(status221)
		case "PORT":
			conn.setDataAddr(args)
		case "LIST":
			conn.list()
		case "RETR":
			conn.retr(args)
		case "STOR":
			conn.stor(args)
		default:
			conn.respond(status502)
		}
	}
}
