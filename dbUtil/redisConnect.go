package dbUtil

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"time"
)

var redC *redis.Client //value? or poinbter

func ConnectToRedis() {
	redC = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//kv database

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() //closure lambda isngle iline anomyous func

	_, err := redC.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error Connecting to Redis")
		panic(err)
	}

	redC.Set(ctx, "key", "value", 0)
	redC.RPush(ctx, "sql dbs", "postgres", "sqlite", "acas", "mysql")
	val, _ := redC.LRange(ctx, "sql dbs", 0, -1).Result()
	fmt.Println(val)
}

func GetRedis() *redis.Client {
	return redC
}
