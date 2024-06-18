package redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IvaCheMih/Redis_pattern/redis_repository/models"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

type RedisService struct {
	client  *redis.Client
	context context.Context
}

func CreateRedisService(client *redis.Client) RedisService {
	ctx := context.Background()
	return RedisService{
		client:  client,
		context: ctx,
	}
}

func (s *RedisService) SetStruct(key string, value interface{}) error {
	valueString, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = s.client.Set(s.context, key, valueString, time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *RedisService) SetStructs(goods []models.Good) {
	for _, good := range goods {
		valueString, err := json.Marshal(good)
		if err != nil {
			continue
		}

		//err = s.client.Set(s.context, strconv.Itoa(good.Id), valueString, time.Minute).Err()
		err = s.client.Set(s.context, strconv.Itoa(good.Id), valueString, -1).Err()

		if err != nil {
			log.Println(err)
		}

	}
}

func (s *RedisService) GetStructs(ids []int) ([]models.Good, []int, []error) {

	var notFoundIds = []int{}
	var goods = []models.Good{}
	var errs = []error{}

	for i := 0; i < len(ids); i++ {

		fmt.Println(i)
		res, err := s.client.Get(s.context, strconv.Itoa(ids[i])).Result()
		if err != nil {
			notFoundIds = append(notFoundIds, ids[i])

			errs = append(errs, err)
		} else {
			var good models.Good

			err = json.Unmarshal([]byte(res), &good)
			if err != nil {
				errs = append(errs, err)

				notFoundIds = append(notFoundIds, ids[i])

				continue
			}
			goods = append(goods, good)
		}
	}

	return goods, notFoundIds, errs
}

func (s *RedisService) GetOneStructs(key string) (models.Good, error) {
	response, err := s.client.Get(s.context, key).Result()
	if err != nil {
		return models.Good{}, err
	}

	var good models.Good

	err = json.Unmarshal([]byte(response), &good)
	if err != nil {
		return models.Good{}, err
	}

	return good, err
}

func (s *RedisService) DeleteStruct(key string) error {
	return s.client.Del(s.context, key).Err()
}

func (s *RedisService) DeleteStructs(keys []string) {
	err := s.client.Del(s.context, keys...)
	if err != nil {
		fmt.Println("goodId: not deleted from redis")
	}

}
