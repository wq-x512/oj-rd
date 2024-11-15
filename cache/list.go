package cache

func LPush(key string, values ...interface{}) (int64, error) {
	return RedisClient.LPush(ctx, key, values...).Result()
}
func RPush(key string, values ...interface{}) (int64, error) {
	return RedisClient.RPush(ctx, key, values...).Result()
}
func LInsert(key string, op string, pivot interface{}, value interface{}) (int64, error) { //名为key的缓存项值第一个值为pivot的元素前后插入一个值，值为value
	// op [before,after]
	return RedisClient.LInsert(ctx, key, op, pivot, value).Result()
}
func LInsertBefore(key string, pivot interface{}, value interface{}) (int64, error) {
	return RedisClient.LInsertBefore(ctx, key, pivot, value).Result()
}
func LInsertAfter(key string, pivot interface{}, value interface{}) (int64, error) {
	return RedisClient.LInsertAfter(ctx, key, pivot, value).Result()
}
func LSet(key string, index int64, value interface{}) (string, error) { // 设置某个元素的值 下标是从0开始的
	return RedisClient.LSet(ctx, key, index, value).Result()
}
func LLen(key string) (int64, error) { //获取链表元素个数
	return RedisClient.LLen(ctx, key).Result()
}
func LIndex(key string, index int64) (string, error) { //获取链表下标对应的元素
	return RedisClient.LIndex(ctx, key, index).Result()
}
func LRange(key string, start int64, stop int64) ([]string, error) { // 获取某个选定范围的元素集
	return RedisClient.LRange(ctx, key, start, stop).Result()
}
func LPop(key string) (string, error) { // 从链表左侧弹出数据
	return RedisClient.LPop(ctx, key).Result()
}
func LRem(key string, count int64, value interface{}) (int64, error) { //根据值移除元素
	return RedisClient.LRem(ctx, key, count, value).Result()
}
