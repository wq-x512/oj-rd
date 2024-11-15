package cache

func SAdd(key string, members ...interface{}) (int64, error) { // 添加元素
	return RedisClient.SAdd(ctx, key, members).Result()
}
func SPop(key string) (string, error) { // 从集合中随机取出元素的
	return RedisClient.SPop(ctx, key).Result()
}
func SPopN(key string, count int64) ([]string, error) {
	return RedisClient.SPopN(ctx, key, count).Result()
}
func SRem(key string, members ...interface{}) (int64, error) { //删除集合里指定的值
	return RedisClient.SRem(ctx, key, members).Result()
}
func SMembers(key string) ([]string, error) { //所有元素
	return RedisClient.SMembers(ctx, key).Result()
}
func SIsMember(key string, member interface{}) (bool, error) { // 判断元素是否在集合中
	return RedisClient.SIsMember(ctx, key, member).Result()
}
func SCard(key string) (int64, error) { //获取集合元素个数
	return RedisClient.SCard(ctx, key).Result()
}
func SUnion(keys ...string) ([]string, error) { //并集
	return RedisClient.SUnion(ctx, keys...).Result()
}
func SDiff(keys ...string) ([]string, error) { // 差集
	return RedisClient.SDiff(ctx, keys...).Result()
}
func SInter(keys ...string) ([]string, error) { // 交集
	return RedisClient.SInter(ctx, keys...).Result()
}
