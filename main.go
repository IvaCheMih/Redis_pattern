package main

import (
	"Redis/redis_repository"
	"Redis/redis_repository/models"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisRepository redis_repository.RedisService

func Init() {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	redisRepository = redis_repository.CreateRedisService(client)

}

func main() {
	Init()

	var good2 = models.Good{
		Id:         2,
		SomeString: "hello again",
		SomeInt:    2,
		SomeBool:   true,
	}

	var good1 = models.Good{
		Id:         1,
		SomeString: "hello",
		SomeInt:    1,
		SomeBool:   true,
	}

	//redisRepository.SetStructs([]models.Good{good1, good2})

	//
	//response, err := redisRepository.GetOneStructs("2")
	//if err != nil {
	//	log.Println(err)
	//}

	//goods, notFoundIds, errs := redisRepository.GetStructs([]int{1, 2})
	//if len(errs) > 0 {
	//	for i, er := range errs {
	//		fmt.Printf("[ERROR] id: %v error: %s \n", notFoundIds[i], er)
	//
	//	}
	//}
	//
	//for _, g := range goods {
	//	fmt.Printf("[*] id: %+v \n", g)
	//}

	err := redisRepository.SetList("list1", []models.Good{good1, good2})
	if err != nil {
		log.Println(err)
	}

	//err = redisRepository.SetList("list1", good2)
	//if err != nil {
	//	log.Println(err)
	//}

	goods, err := redisRepository.GetList("list1", 0, 2)
	if err != nil {
		log.Println(err)
	}

	fmt.Println()

	fmt.Println("Goods: ")
	for _, good := range goods {
		fmt.Println(good)
	}

	goods, err = redisRepository.DeleteFromList("list1", 3)
	if err != nil {
		log.Println(err)
	}

	fmt.Println()

	fmt.Println("Delete Goods: ")
	for _, good := range goods {
		fmt.Println(good)
	}

}
