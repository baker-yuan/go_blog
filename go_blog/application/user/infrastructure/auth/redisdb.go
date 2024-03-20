package auth

import "github.com/go-redis/redis/v7"

type RedisService struct {
	Auth   AuthInterface
	Client *redis.Client
}

func NewRedisDB() (*RedisService, error) {
	// todo
	redisClient := redis.NewClient(&redis.Options{
		//Addr:     host + ":" + port,
		//Password: password,
		//DB:       0,
	})
	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
