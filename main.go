package main

import (
    "context"
	"./util"
	"./cache"
	"./models"
)

func main(){
	redisSetup  	:= &util.RedisConn{Ctx : context.Background()}
	redisConn,_     := redisSetup.Conn() 
	cache 	   		:= &cache.CacheProcess{Cache : redisConn, Ctx : context.Background()}

	employees := []models.Employee { 
		models.Employee {
			Id: 1, 
			Name: "Iron Man", 
			Nik: "111", 
			Division: "IT Developer",
		},
	}

	cache.Set("Employee", employees)
	cache.Get("Employee")

}