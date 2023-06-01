package dbUtil

import (
	"context"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"main/storageUtil"
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
	//gariuavbansal hv gaiurav bansal gaurav bansal hv data
	fmt.Println(val)
}

func GetRedis() *redis.Client {
	return redC
}

func addMsg(msg storageUtil.Message, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	redC.RPush(ctx, key, msg) //pus hstruct as json?
}

func GetMsg(msg string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//closure and func

	return redC.Get(ctx, msg).Result()
	//redis db key value database
	//I love redis because it is key value and allows you to store any data structure in its secure and amazingf datavbase with ovre 10s of sdks for different lkangague form go to rbyuy. ity also fhosts on a dedicated port on linux devices, 5147, which make it ideal for develoeprs to uytilize in their own applications with a custom port to host their atabse on.
}
