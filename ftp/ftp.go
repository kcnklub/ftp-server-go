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
			return
		case "PORT":
			conn.connect(args)
			return
		default:
			conn.respond(status502)
		}
	}
}
