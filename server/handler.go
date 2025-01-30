package server

import (
	"log"
	"net"
)


func HandleConnection(conn net.Conn) {
    log.Printf(
		"Accepted connection from %s\n", 
		conn.RemoteAddr().String(),
	)
    defer conn.Close()

    buffer := make([]byte, 1024)

    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err.Error() == "EOF" {
                log.Printf(
					"Client closed the connection from %s\n", 
					conn.RemoteAddr().String(),
				)
                return
            }
            log.Printf(
				"Error reading from %s: %s",
			 	conn.RemoteAddr().String(),
				err.Error(),
			)
            return
        }
        if n > 0 {
            data := buffer[:n]
            log.Printf("Received %d bytes: %q", n, data)
			log.Printf("Parsing: %s", ParseInput(data))
        }
    }
}