package cache

import (
	"../models"
	"context"
	"github.com/go-redis/redis"
	"encoding/json"
)

type CacheProcess struct {
	Cache		  *redis.Client
	Ctx           context.Context
}

func (redis *CacheProcess) Set(key string, value []models.Employee) (err error) {

	Json, _ := json.Marshal(value)
	set := redis.Cache.Set(redis.Ctx, key, string(Json), 0)
	if err := set.Err(); err != nil {
		return err
	}
	return nil
}

func (redis *CacheProcess) Get(key string) (res []models.Employee, err error) {

	val,err := redis.Cache.Get(redis.Ctx, key).Result()
	if err != nil {
		return []models.Employee{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		return []models.Employee{}, err
	}

	return res,err
}