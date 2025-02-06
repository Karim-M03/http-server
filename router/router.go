package router

import (
	"errors"
	"karim/http_server/client"
	"karim/http_server/constants"
	"karim/http_server/logger"
	"karim/http_server/request"
	"karim/http_server/response"
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


func (r *Router) RegisterEndpoint(endpoint string, handler func(*client.Client)) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if _, exist := r.Endpoints[endpoint]; exist {
		return errors.New("Endpoint already registered")
	}
	r.Endpoints[endpoint] = handler
	logger.InfoLogger.Printf("Registered endpoint: %s", endpoint)
	return nil
}


func (r *Router) HandleConnection(conn net.Conn) {
    defer conn.Close()
	var clientObj client.Client
	clientObj.Connection = conn

    buffer := make([]byte, constants.BUFFER_SIZE)

    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err.Error() == "EOF" {
                return
            }
            return
        }

        if n > 0 {
			logger.InfoLogger.Println("Reading more than 0 bytes from the request")
            data := buffer[:n]
			request, code, err := request.DivideMessage(data)
			if code == -1{
				clientObj.SendResponse( response.TmpResponse{
					Status: 400,
					Message: err.Error(),
				})
				return
			}
			logger.InfoLogger.Printf("Reading Request \n%s", request.String())

			clientObj.Request = *request
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