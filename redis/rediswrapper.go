package main

import (
    "github.com/garyburd/redigo/redis"
    "os"
    "time"
    "fmt"
)

var (
    Pool *redis.Pool
)

// dbinit Get Tedis server Ip:port and connect to it
func dbinit(server string ) {
    redisHost := os.Getenv("REDIS_HOST")
    if redisHost == "" {
        redisHost = server
    }
    fmt.Printf("RedisServer: %s",redisHost)
    Pool = connectServer(redisHost)
}

//connectServer connect to Redis server
func connectServer(server string) *redis.Pool {

    return &redis.Pool{

        MaxIdle:     5,
        IdleTimeout: 1300 * time.Second,

        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", server)
            if err != nil {
                fmt.Errorf("error connecting to Redis server  %s: %v", server, err)
                return nil, err
            }
            return c, err
        },

    }
}

// Delete key value from DB
func Delete(key string) (error) {

    conn := Pool.Get()
    defer conn.Close()

    _, err := redis.Int(conn.Do("DEL", key))
    return err
}

// Get key value from DB
func Get(key string) (string, error) {

    conn := Pool.Get()
    defer conn.Close()

    var data string
    data, err := redis.String(conn.Do("GET", key))
    if err != nil {
        return data, fmt.Errorf("error getting key %s: %v", key, err)
    }
    return data, err
}

// Create and/or set the key and value in DB
func Set(key string, value string) error {

    conn := Pool.Get()
    defer conn.Close()

    _, err := conn.Do("SET", key, value)
    if err != nil {
        return fmt.Errorf("error setting key %s to %s: %v", key, value, err)
    }
    return err
}

// publish key and value which will be received by all subscribers
func Publish(key string, value string) error {
	conn := Pool.Get()
	conn.Do("PUBLISH", key, value)
	return nil
}

// Subscribe to pattern and waite for subscripton, message
func  Subscribe(key string) (string, error) {
	rc := Pool.Get()
	psc := redis.PubSubConn{Conn: rc}

	// subscribe to any keyspace events
    if err := psc.Subscribe("__key*__:*"); err != nil {
    //if err := psc.Subscribe(key); err != nil {
        fmt.Errorf("error psubscribing key %s : %v", key, err)
		return "", err
	}

    kvar := DeliverMessages(psc)
	return kvar, nil

}

func DeliverMessages(psc redis.PubSubConn) (string) {
	fmt.Println("Entering Sleep")
    //timeout := time.Duration(10 * time.Second)
	for {
    //switch v := psc.ReceiveWithTimeout(timeout).(type) {
    switch v := psc.Receive().(type) {
         case redis.Message:
             fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			 return string(v.Data)
         }
	}
	fmt.Println("Exiting Sleep")
	fmt.Printf("None of the ase match:returning")
    return ""
}
