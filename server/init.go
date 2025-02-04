package server

import (
	"karim/http_server/endpoints"
	"karim/http_server/router"
	"log"
	"net"
)

func Init(){
	/* logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening log file:", err)
        return
    }
    defer logFile.Close()

    // Configure log to write to the file
    log.SetOutput(logFile) */

    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    log.Printf("Starting server on %s\n", listener.Addr().String())

    router := router.NewRouter()
	endpoints.AddEndpoints(router)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Error on accepting the connection:", err)
            break
        }
        go router.HandleConnection(conn)
    }
}