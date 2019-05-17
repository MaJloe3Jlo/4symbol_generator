package util

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

// func RedisConnect is return connector for Redis.
func RedisConnect() redis.Conn {
	// Get the connection object
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal("Error connect to Redis")
	}

	return conn
}
