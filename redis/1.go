package main

import (
       "github.com/garyburd/redigo/redis"
       "fmt"
       "log"
       )
func main(){

      conn, err := redis.Dial("tcp", "172.17.0.2:6379")
      if err != nil {
          fmt.Println("Connection Failed",err)
          log.Fatal(err)
      }   
      defer conn.Close()
      var key string
        key = "True"
/*        psc := redis.PubSubConn{Conn: conn}
        if err := psc.PSubscribe(key); err != nil {
//      if err := psc.PSubscribe("__key*__:*"); err != nil {
            fmt.Errorf("error psubscribing key %s : %v", key, err)
            return 
        }
*/
        _,err2 := conn.Do("SET",key,"1")
        if err2 != nil {
            fmt.Errorf("error SETting key %s : %v", key, err2)
        }
        val,err3 := conn.Do("GET",key,)
            fmt.Printf("GETting key %s val %s: %v", key,val, err3)
    
       
//    DeliverMessages(psc)
    return 
}

/*
func DeliverMessages(psc redis.PubSubConn){

     for {
         switch v := psc.Receive().(type) {
         case redis.PMessage:
             fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
         case redis.Subscription:
             fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
         case error:
             fmt.Printf("Error")
         }
     }
}
*/
