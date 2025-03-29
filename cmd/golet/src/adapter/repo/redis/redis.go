package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Init(ctx context.Context) error {
	rds := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:237",
		Password: "",
		DB:       0,
	})
	status := rds.Ping(ctx)
	if status != nil {
		err := fmt.Errorf("connet to redis failed, err=%s", status)
		return err
	}
	return nil
}
