package redisclient

import (
	"github.com/go-redis/redis"
	"log"
	"time"
)

func Init(addr string) *redis.Client {
	client := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	if _, err := client.Ping().Result(); err != nil {
		handleErr(err)
	}
	return client
}

func Setex(client *redis.Client, key string, value string, minutes int) {
	if _, err := client.Set(key, value, 0).Result(); err != nil {
		handleErr(err)
	}
	client.Expire(key, time.Duration(minutes) * time.Minute)
}

func GetKeyValue(client *redis.Client, key string) string {
	value, err := client.Get(key).Result()
	if err != nil {
		handleErr(err)
	}
	return value
}

func IsExists(client *redis.Client, key string) bool {
	value,err := client.Exists(key).Result()
	if err != nil {
		handleErr(err)
	}
	return value == 1;
}

func Delete(client *redis.Client, key string) {
	_, err := client.Del(key).Result()
	if err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	log.Print(err.Error())
	panic(err)
}


