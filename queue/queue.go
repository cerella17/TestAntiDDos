package queue

import (
    "context"
    "time"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
    return client
}

func EnqueueRequest(client *redis.Client, request string) error {
    return client.LPush(ctx, "request_queue", request).Err()
}

func DequeueRequest(client *redis.Client) (string, error) {
    return client.RPop(ctx, "request_queue").Result()
}
func ProcessQueue(){
   client := ConnectRedis()
   for{
	request, err := DequeueRequest(client)
	if err == nil && request != " "{
		fmt.Printf("Processing request: %s\n",request)
	} else if err != nil {
		fmt.Printf("error processing: %v\n",err)          
	 }
        time.Sleep(1*time.Second)
	}
     }
