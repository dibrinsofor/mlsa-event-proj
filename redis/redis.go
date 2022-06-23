package redis

import (
	"context"
	"crypto/tls"
	"log"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/go-redis/redis/v9"
)

func ConnectRedis() *redis.Client {
	// host := os.Getenv("REDIS_HOST")
	// if host == "" {
	// 	host = "127.0.0.1:6379"
	// }
	// pass := os.Getenv("REDIS_PASSWORD")
	// if pass == "" {
	// 	pass = "password"
	// }
	host := "mlsa3.redis.cache.windows.net:6380"
	pass := "elf7mgl7Tp9NNVaZADQOfIvDOCdQHtvUOAzCaBFGll8="
	tlsConfig := &tls.Config{MinVersion: tls.VersionTLS12}
	rdb := redis.NewClient(&redis.Options{
		Addr:      host,
		Password:  pass,
		DB:        0,
		TLSConfig: tlsConfig,
	})

	return rdb
}

func AddUserInstance(user *models.User) ([]redis.Cmder, error) {
	rdb := ConnectRedis()

	val, err := rdb.Pipelined(context.TODO(), func(rdb redis.Pipeliner) error {
		rdb.HSet(context.TODO(), user.ID, "FirstName", user.FirstName)
		rdb.HSet(context.TODO(), user.ID, "LastName", user.LastName)
		rdb.HSet(context.TODO(), user.ID, "Email", user.Email)
		rdb.HSet(context.TODO(), user.ID, "CreatedAt", user.CreatedAt)
		return nil
	})
	if err != nil {
		panic(err)

	}

	defer rdb.Close()
	return val, err
}

func FindUserByID(ID string) (userObject models.User) {
	rdb := ConnectRedis()

	if err := rdb.HGetAll(context.TODO(), ID).Scan(&userObject); err != nil {
		log.Println(err)
	}

	defer rdb.Close()
	return userObject
}
