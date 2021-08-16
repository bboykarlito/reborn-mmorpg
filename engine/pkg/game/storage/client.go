package storage

import (
	"os"
	"encoding/json"
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/gavr-games/reborn-mmorpg/pkg/game/entity"
)

// declaration defined type 
type StorageClient struct {
	redisClient *redis.Client
	Updates chan *entity.GameObject
}

var instance *StorageClient = nil
var ctx = context.Background()
var once sync.Once

func GetClient() *StorageClient {
	once.Do(func() {
		opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
				panic(err)
		}
		rdb := redis.NewClient(opt)
		instance = &StorageClient{
			redisClient: rdb,
			Updates: make(chan *entity.GameObject),
		}
	})
	return instance
}

func (sc *StorageClient) SaveGameObject(obj *entity.GameObject) {
	message, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	setErr := sc.redisClient.Set(ctx, obj.Id, message, 0).Err()
  if setErr != nil {
    panic(setErr)
  }
}

func (sc *StorageClient) GetGameObject(id string) *entity.GameObject {
	val, redisErr := sc.redisClient.Get(ctx, id).Result()
	var obj entity.GameObject
	if redisErr != nil {
		panic(redisErr)
	}
	err := json.Unmarshal([]byte(val), &obj)
	if err != nil {
		panic(err)
	}
	return &obj
}

func (sc *StorageClient) ReadAllGameObjects(process func(*entity.GameObject)) int {
	i := 0
	iter := sc.redisClient.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		obj := sc.GetGameObject(iter.Val())
		process(obj)
		i++
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return i
}

func (sc *StorageClient) Run() {
	for {
		select {
		case obj := <-sc.Updates:
			sc.SaveGameObject(obj)
		}
	}
}
