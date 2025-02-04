package router

import (
	"errors"
	"fmt"
	"karim/http_server/client"
	"karim/http_server/request"
	"karim/http_server/response"
	"log"
	"net"
	"sync"
)

/*
In context there willl be
- Connection information of the user (that will be ahndled)
- return object
*/
type Router struct{
	Mutex sync.RWMutex
	Endpoints map[string]func(*client.Client)
}

func NewRouter() *Router {
	return &Router{
		Endpoints: make(map[string]func(*client.Client)),
	}
}


func (r *Router) RegisterEndpoint(endpoint string, handler func(*client.Client))(error){
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	if _, exist := r.Endpoints[endpoint]; exist{
		return errors.New("Endpoint already registered")
	}
	r.Endpoints[endpoint] = handler
	return nil
}


func (r *Router) HandleConnection(conn net.Conn) {
    defer conn.Close()
	var clientObj client.Client
	//log.Printf("Accepting client: %s", conn.LocalAddr().String())
	clientObj.Connection = conn

    buffer := make([]byte, 1024)

    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err.Error() == "EOF" {
                return
            }
            return
        }

        if n > 0 {
			log.Printf("Reading more than 0 bytes from the request")
            data := buffer[:n]
			request, code, err := request.DivideMessage(data)
			if code == -1{
				clientObj.SendResponse( response.TmpResponse{
					Status: 400,
					Message: err.Error(),
				})
				return
			}
			log.Printf("Reading Request \n%s", request.String())

			clientObj.Request = *request
			fmt.Println(request.StartLine.String(false))
			handler, exists := r.Endpoints[request.StartLine.String(false)]
			if !exists{
				
				clientObj.SendResponse(response.TmpResponse{
					Data: nil,
					Message: "Route not found",
					Status: 404,
				})
				return
			}

			handler(&clientObj)

			// for testing purpose, we may want to close it only if explicitly told
			clientObj.Connection.Close()

        }
    }
}