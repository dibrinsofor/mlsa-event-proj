package redis

import (
	"context"
	"log"

	"github.com/dibrinsofor/mlsa3/models"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "password",
		DB:       0,
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

	return val, err
}

func FindUserByID(ID string) (userObject models.User) {
	rdb := ConnectRedis()

	if err := rdb.HGetAll(ctx, ID).Scan(&userObject); err != nil {
		log.Println(err)
	}

	return userObject
}
