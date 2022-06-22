package redis

import (
	"context"
	"crypto/tls"
	"log"
	"os"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/go-redis/redis/v9"
)

var ctx = context.TODO()

func ConnectRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "127.0.0.1:6379"
	}
	pass := os.Getenv("REDIS_PASSWORD")
	if pass == "" {
		pass = "password"
	}
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

	val, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, user.ID, "FirstName", user.FirstName)
		rdb.HSet(ctx, user.ID, "LastName", user.LastName)
		rdb.HSet(ctx, user.ID, "Email", user.Email)
		rdb.HSet(ctx, user.ID, "CreatedAt", user.CreatedAt)
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

	if err := rdb.HGetAll(ctx, ID).Scan(&userObject); err != nil {
		log.Println(err)
	}

	defer rdb.Close()
	return userObject
}
