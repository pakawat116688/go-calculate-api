package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

func NewCheckOperator() CheckOperator {
	return func(input Input) (result float64, err error) {
		switch input.Operator {
		case "+":
			return input.Number_1 + input.Number_2, nil
		case "-":
			return input.Number_1 - input.Number_2, nil
		case "*":
			return input.Number_1 * input.Number_2, nil
		case "/":
			if input.Number_2 == 0 {
				return 0, errors.New("number2 cannot be equal to 0")
			}
			return input.Number_1 / input.Number_2, nil
		default:
			return 0, errors.New("operator not found")
		}
	}
}

func NewGetRedis(redisClint *redis.Client) GetRedis {
	return func(ctx context.Context, key string) (result float64, err error) {
		if resultStr, err := redisClint.Get(ctx, key).Result(); err == nil {
			if json.Unmarshal([]byte(resultStr), &result) == nil {
				fmt.Println("valuse from redis.....")
				return result, nil
			}
		}
		return 0, errors.New("key not found")
	}
}

func NewSetRedis(redisClint *redis.Client) SetRedis {
	return func(ctx context.Context, key string, result float64) {
		if resultByte, err := json.Marshal(result); err == nil {
			redisClint.Set(ctx, key, string(resultByte), time.Second*10)
		}
		fmt.Println("set in redis.....")
	}
}
