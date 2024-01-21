package kv

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"lab.sda1.net/nexryai/altcore/internal/core/config"
	"lab.sda1.net/nexryai/altcore/internal/core/logger"
	"time"
)

var ctx = context.Background()

func genStoreKey(cacheKey string) string {
	return fmt.Sprintf("%s:cache/%s", config.Host, cacheKey)
}

func keyExists(ctx context.Context, client *redis.Client, key string) bool {
	// Existsメソッドでキーの存在を確認
	result, err := client.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return result == 1
}

func StoreKvCache(key string, data interface{}) {
	clinet := ConnectToRedis()

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = clinet.Set(ctx, genStoreKey(key), string(jsonData), 1*time.Hour).Err()
	if err != nil {
		panic(err)
	}

	logger.Debug(fmt.Sprintf("Cache saved!: %s", genStoreKey(key)))
}

func GetKvCache(key string, data interface{}) bool {
	clinet := ConnectToRedis()

	if !keyExists(ctx, clinet, genStoreKey(key)) {
		// 存在しないならfalseを返す
		return false
	}

	// Redisからデータを取得
	result, err := clinet.Get(ctx, genStoreKey(key)).Result()
	if err != nil {
		panic(err)
	}

	// JSONを構造体にデコード
	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		panic(err)
	}

	logger.Debug(fmt.Sprintf("Cache used!: %s", genStoreKey(key)))
	return true

}

// RamStoreは、起動時に初期化される頻繁に呼び出されるデータのDBからのコピーです。
func GetRamStore(key string) interface{} {
	var data interface{}

	exists := GetKvCache(key, data)

	if !exists {
		panic("GetRamStore: Data not found")
	}

	return data
}

func InitRamStore() {
}
