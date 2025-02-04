package classes

import (
	"karim/http_server/client"
	"karim/http_server/response"
	"karim/http_server/router"
	"log"
)

func AddClassesEndpoints(r *router.Router){
	r.RegisterEndpoint("GET /classes", GetClasses)
}


func GetClasses(clnt *client.Client){
	log.Printf("I'm in GetClasses")

	clnt.SendResponse(response.TmpResponse{
		Status: 200,
		Message: "Classes retreived successfully",
		Data: nil,
	})
}