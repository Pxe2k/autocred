package helpers

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var Ctx = context.Background()

var Redis = redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDIS_ADDRESS") + ":" + os.Getenv("REDIS_PORT"),
	//Addr:     "127.0.0.1:6379",
	Password: "",
})

func GenerateCode(phone string) (string, error) {
	authCode := RandEmailCode()

	err := Redis.Set(Ctx, phone, authCode, 6000000000000).Err()
	if err != nil {
		return "error", err
	}

	return authCode, nil
}

func SetToken(bank, token string) error {
	err := Redis.Set(Ctx, bank, token, 6000000000000).Err()
	if err != nil {
		return err
	}

	return nil
}
