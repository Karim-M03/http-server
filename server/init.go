package server

import (
	"karim/http_server/endpoints"
	"karim/http_server/logger"
	"karim/http_server/router"
	"net"
    "os"
)

func Init(){
    port := os.Getenv("SERVER_PORT")
	if port == "" {
		logger.ErrorLogger.Println("Failed to get environment variable: SERVER_PORT")
	}
    listener, err := net.Listen("tcp", port)
    if err != nil {
        logger.ErrorLogger.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    logger.InfoLogger.Printf("Starting server on %s\n", listener.Addr().String())

    router := router.NewRouter()
	endpoints.AddEndpoints(router)

    for {
        conn, err := listener.Accept()
        if err != nil {
            logger.ErrorLogger.Printf("Error on accepting the connection:", err)
            break
        }
        go router.HandleConnection(conn)
    }
}