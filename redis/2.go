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
//      var key string
//        key = "True"
/*        psc := redis.PubSubConn{Conn: conn}
        if err := psc.PSubscribe(key); err != nil {
//      if err := psc.PSubscribe("__key*__:*"); err != nil {
            fmt.Errorf("error psubscribing key %s : %v", key, err)
            return 
        }
*/
    _,err=conn.Do("CONFIG","SET", "notify-keyspace-events", "KEA")
    if err!=nil {
        fmt.Printf("CONFIG KEA error : %s\n", err)
    }
    psc := redis.PubSubConn{Conn: conn}
    psc.PSubscribe("__key*__:*")
    DeliverMessages(psc)
    return 
}


func DeliverMessages(psc redis.PubSubConn){

     for {
         switch msg := psc.Receive().(type) {
         case redis.Message:
             fmt.Printf("%s: message: %s\n", msg.Channel, msg.Data)
         case redis.PMessage:
             fmt.Printf("%s , %s: Pmessage: %s\n", msg.Pattern, msg.Channel, msg.Data)
         case redis.Subscription:
             fmt.Printf("%s: %s, %d\n", msg.Channel, msg.Kind, msg.Count)
         case error:
             fmt.Printf("Error")
         }
     }
}

func handler(){
    fmt.Printf("====HANDLER=== \n")
}
