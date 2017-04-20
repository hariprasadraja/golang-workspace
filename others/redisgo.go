package main

import (
	"gopkg.in/redis.v3"
	"log"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("Redis connection err:", err)
	}

	client.Set("name","Hariprasad",0)

}


