package main

import (
	//"karim/http_server/logger"
	"karim/http_server/server"

	//"github.com/joho/godotenv"
)

func main() {
   /*  err := godotenv.Load()
	if err != nil {
		logger.ErrorLogger.Fatal("Error loading .env file")
	} */

	
    // init the db
    
   /*  if err != nil{
        logger.ErrorLogger.Fatal(err.Error())
    } */


    // Open log file
   server.Init()
}



