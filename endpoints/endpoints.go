package endpoints

import (
	"karim/http_server/endpoints/classes"
	"karim/http_server/endpoints/exams"
	"karim/http_server/endpoints/students"
	"karim/http_server/router"
)

func AddEndpoints(r *router.Router){
	students.AddStudentsEndpoints(r)
	classes.AddClassesEndpoints(r)
	exams.AddExamsEndpoints(r)
}