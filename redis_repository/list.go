package redis_repository

import (
	"Redis/redis_repository/models"
	"encoding/json"
	"fmt"
)

func (s *RedisService) SetList(key string, goods []models.Good) error {

	fmt.Println("set goods:")
	for _, good := range goods {
		fmt.Println(good)

		g, err := json.Marshal(&good)
		if err != nil {
			return err
		}

		err = s.client.RPush(s.context, key, g).Err()
		if err != nil {
			return err
		}
	}

	fmt.Println()

	return nil

}

func (s *RedisService) GetList(key string, start int64, end int64) ([]models.Good, error) {

	response := s.client.LRange(s.context, key, start, end)
	if response.Err() != nil {
		return nil, response.Err()
	}

	var goods []models.Good

	for _, res := range response.Val() {
		fmt.Println(res)

		var good models.Good

		err := json.Unmarshal([]byte(res), &good)
		if err != nil {
			return nil, err
		}

		goods = append(goods, good)
	}

	return goods, nil
}

func (s *RedisService) DeleteFromList(key string, count int) ([]models.Good, error) {

	response := s.client.LPopCount(s.context, key, count)
	if response.Err() != nil {
		return nil, response.Err()
	}

	var goods []models.Good

	for _, res := range response.Val() {
		fmt.Println(res)

		var good models.Good

		err := json.Unmarshal([]byte(res), &good)
		if err != nil {
			return nil, err
		}

		goods = append(goods, good)
	}

	return goods, nil
}
