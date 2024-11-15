package cache

func HSet(key string, values ...interface{}) (int64, error) {
	return RedisClient.HSet(ctx, key, values...).Result()
}
func HMSet(key string, values ...interface{}) (bool, error) { //批量设置
	return RedisClient.HMSet(ctx, key, values...).Result()
}
func HGet(key string, field string) (string, error) { //获取某个元素
	return RedisClient.HGet(ctx, key, field).Result()
}
func HDel(key string, fields ...string) (int64, error) {
	return RedisClient.HDel(ctx, key, fields...).Result()
}
func HGetAll(key string) (map[string]string, error) {
	return RedisClient.HGetAll(ctx, key).Result()
}
func HExists(key string, field string) (bool, error) { //判断元素是否存在
	return RedisClient.HExists(ctx, key, field).Result()
}
func HLen(key string) (int64, error) { //获取长度
	return RedisClient.HLen(ctx, key).Result()
}
