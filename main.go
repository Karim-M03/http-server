package main

import (
    "fmt"
    "log"
    "net"
    "os"
)

func main() {
    // Open log file
    logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening log file:", err)
        return
    }
    defer logFile.Close()

    // Configure log to write to the file
    log.SetOutput(logFile)

    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    log.Printf("Starting server on %s\n", listener.Addr().String())

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println("Error on accepting the connection:", err)
            break
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    log.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())
    defer conn.Close()
    var message = "Hello there.."
    n, _ := conn.Write([]byte(message))

    if n > 0 && n != len(message){
        log.Println("The message was sent partially")
    } else if n == 0 {
        log.Println("The message was not sent")
    }
}
