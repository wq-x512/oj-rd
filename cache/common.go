package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)
var ctx = context.Background()

func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("ini load error")
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPwd").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		DB:   int(db),
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis connect error", err)
		panic(err)
	}
	RedisClient = client
}
func Keys() ([]string, error) { // 根据正则获取keys
	return RedisClient.Keys(ctx, "*").Result()
}
func Type(key string) (string, error) { // 获取一个key对应值的类型
	return RedisClient.Type(ctx, key).Result()
}
func Del(keys []string) (int64, error) {
	return RedisClient.Del(ctx, keys...).Result()
}
func Exists(keys []string) (int64, error) {
	return RedisClient.Exists(ctx, keys...).Result()
}

func Expire(key string, duration time.Duration) (bool, error) { //某个时间段(time.Duration)后过期
	return RedisClient.Expire(ctx, key, duration).Result()
}
func ExpiresAt(key string, time time.Time) (bool, error) { //某个时间点(time.Time)过期失效
	return RedisClient.ExpireAt(ctx, key, time).Result()
}
func GetTTL(key string) (time.Duration, error) { // 获取某个键的剩余有效期s
	return RedisClient.TTL(ctx, key).Result()
}
func GetPTTL(key string) (time.Duration, error) { // 获取某个键的剩余有效期ms
	return RedisClient.PTTL(ctx, key).Result()
}
func DBSize() (int64, error) { //查看当前数据库key的数量
	return RedisClient.DBSize(ctx).Result()
}
func FlushDB() (string, error) { // 清空当前数据
	return RedisClient.FlushDB(ctx).Result()
}
func FlushAll() (string, error) { //清空所有数据库
	return RedisClient.FlushAll(ctx).Result()
}
