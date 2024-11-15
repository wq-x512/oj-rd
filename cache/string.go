package cache

import (
	"time"
)

func Set(key string, value interface{}, expiration time.Duration) error { //最后一个参数表示过期时间，0表示永不过期
	return RedisClient.Set(ctx, key, value, expiration).Err()
}
func SetEX(key string, value interface{}, expiration time.Duration) error { // 设置并指定过期时间
	return RedisClient.SetEX(ctx, key, value, expiration).Err()
}
func SetNX(key string, value interface{}, expiration time.Duration) error { // 设置并指定过期时间
	return RedisClient.SetNX(ctx, key, value, expiration).Err()
}

// SexNX()仅当key不存在的时候才设置，如果key已经存在则不做任何操作，而SetEX()方法不管该key是否已经存在缓存中直接覆盖
func Get(key string) error { // 如果要获取的key在缓存中并不存在，Get()方法将会返回redis.Nil
	return RedisClient.Get(ctx, key).Err()
}
func GetRange(key string, start int64, end int64) (string, error) { // 字符串截取 ,即使key不存在，调用GetRange()也不会报错，只是返回的截取结果是空""
	return RedisClient.GetRange(ctx, key, start, end).Result()
}
func Incr(key string) (int64, error) {
	return RedisClient.Incr(ctx, key).Result()
}
func IncrBy(key string, value int64) (int64, error) {
	return RedisClient.IncrBy(ctx, key, value).Result()
}
func Decr(key string) (int64, error) {
	return RedisClient.Decr(ctx, key).Result()
}
func DecrBy(key string, value int64) (int64, error) {
	return RedisClient.DecrBy(ctx, key, value).Result()
}
func Append(key string, value string) (int64, error) { //表示往字符串后面追加元素，返回值是字符串的总长度
	return RedisClient.Append(ctx, key, value).Result()
}
func Strlen(key string) (int64, error) { //获取字符串的长度
	return RedisClient.StrLen(ctx, key).Result()
}
