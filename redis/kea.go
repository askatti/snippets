package main

import (
       "github.com/garyburd/redigo/redis"
       "fmt"
       "sync"
       )

type RedisClient struct {
    mRedisServer     string
    mRedisConn       redis.Conn
    mWg              sync.WaitGroup
}

func (rc *RedisClient) Run() {
    conn, err := redis.Dial("tcp", ":6379")
    if err != nil {
        fmt.Println(err)
        return
    }
    rc.mRedisConn = conn
    fmt.Println(conn)
    rc.mRedisConn.Do("CONFIG", "SET", "notify-keyspace-events", "KEA")

    fmt.Println("Set the notify-keyspace-events to KEA")
    defer rc.mRedisConn.Close()
    rc.mWg.Add(2)
    psc := redis.PubSubConn{Conn: rc.mRedisConn}
    go func() {
        defer rc.mWg.Done()
        for {
            switch msg := psc.Receive().(type) {
            case redis.Message:
                fmt.Printf("Message: %s %s\n", msg.Channel, msg.Data)
            case redis.PMessage:
                fmt.Printf("PMessage: %s %s %s\n", msg.Pattern, msg.Channel, msg.Data)
            case redis.Subscription:
                fmt.Printf("Subscription: %s %s %d\n", msg.Kind, msg.Channel, msg.Count)
                if msg.Count == 0 {
                    return
                }
            case error:
                fmt.Printf("error: %v\n", msg)
                return
            }
        }
    }()
    go func() {
        defer rc.mWg.Done()
        psc.PSubscribe("__key*__:*")
        select {}
    }()
    rc.mWg.Wait()
}

func main(){
  // var rc *RedisClient
    Run()
}
