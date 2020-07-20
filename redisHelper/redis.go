package sessionHelper

import (
	"time"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client
var prefix string

func InitRedisHelper(host string, port string, password string, db int, pre string) {

	if host == "" {
		host = "127.0.0.1"
	}

	if port == "" {
		port = "6379"
	}

	// 根据redis配置初始化一个客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port, // redis地址
		Password: password,          // redis密码，没有则留空
		DB:       db,                // 默认数据库，默认是0
	})

	prefix = pre
}

func getRealKey(key string) string {
	return prefix + key
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetCacheData(key string) string {
	return redisClient.Get(getRealKey(key)).Val()
}

func SetCacheData(key string, value []byte, expire time.Duration) {
	redisClient.Set(getRealKey(key), value, expire)
}

func Incr(key string, expireTime time.Duration) {
	redisClient.Incr(getRealKey(key))
	redisClient.Expire(getRealKey(key), expireTime)
}

func Decr(key string, expireTime time.Duration) {
	redisClient.Decr(getRealKey(key))
	redisClient.Expire(getRealKey(key), expireTime)
}

func RemoveCacheData(key string) {
	redisClient.Del(getRealKey(key))
}

func IsNameExits(key string) int64 {
	return redisClient.Exists(getRealKey(key)).Val()
}

func SetExpireTime(key string, expireTime time.Duration) {
	redisClient.Expire(getRealKey(key), expireTime)
}

func RenameKey(key string, newKey string) {
	redisClient.Rename(getRealKey(key), newKey)
}

func PushListFromLeft(key string, values ...interface{}) {
	redisClient.LPush(getRealKey(key), values)
}

func PushListFromRight(key string, values ...interface{}) {
	redisClient.RPush(getRealKey(key), values)
}

func PopListFromLeft(key string) string {
	return redisClient.LPop(getRealKey(key)).Val()
}

func GetListValueByIndex(key string, index int64) string {
	return redisClient.LIndex(getRealKey(key), index).Val()
}

func InsertIntoListByValue(key string, index int64, value interface{}) string {
	return redisClient.LSet(getRealKey(key), index, value).Val()
}

//正数 从左往右 负数从右往左 第n个
func RemoveListValue(key string, count int64, value interface{}) int64 {
	return redisClient.LRem(getRealKey(key), count, value).Val()
}

/************** hash 操作 ********************/

//设置单独的hashkey
func SetSingleHashData(key string, field string, value interface{}, expireTime time.Duration) {
	redisClient.HSet(getRealKey(key), field, value)
	redisClient.Expire(getRealKey(key), expireTime)
}

func GetSingleHashData(key string, field string) string {
	return redisClient.HGet(getRealKey(key), field).Val()
}

func ResetAllHashData(key string, fields map[string]interface{}, expireTime time.Duration) {
	redisClient.HMSet(getRealKey(key), fields)
	redisClient.Expire(getRealKey(key), expireTime)
}

func GetAllHashKeys(key string) []string {
	return redisClient.HKeys(getRealKey(key)).Val()
}

func GetAllHashValues(key string) []string {
	return redisClient.HVals(getRealKey(key)).Val()
}

func GetAllHashMaps(key string) map[string]string {
	return redisClient.HGetAll(getRealKey(key)).Val()
}

func IsHashMapExists(key string, field string) bool {
	return redisClient.HExists(getRealKey(key), field).Val()
}

func RemoveSingleHashMap(key string, field string) bool {
	return redisClient.HExists(getRealKey(key), field).Val()
}

func HashIncrInt(key string, field string, offset int64, expireTime time.Duration) {
	redisClient.HIncrBy(getRealKey(key), field, offset)
	redisClient.Expire(getRealKey(key), expireTime)
}

func HashDecrInt(key string, field string, offset int64, expireTime time.Duration) {
	redisClient.HIncrBy(getRealKey(key), field, -offset)
	redisClient.Expire(getRealKey(key), expireTime)
}

func HashIncrFloat(key string, field string, offset float64, expireTime time.Duration) {
	redisClient.HIncrByFloat(getRealKey(key), field, offset)
	redisClient.Expire(getRealKey(key), expireTime)
}

func HashDecrFloat(key string, field string, offset float64, expireTime time.Duration) {
	redisClient.HIncrByFloat(getRealKey(key), field, -offset)
	redisClient.Expire(getRealKey(key), expireTime)
}

/*************** sorted list ************************/

func AddSortSet(key string, members ...redis.Z) {
	redisClient.ZAdd(getRealKey(key), members...)
}

func GetSortedSetLength(key string) int64 {
	return redisClient.ZCard(getRealKey(key)).Val()
}

// 根据index 从start到end获取有序数组的值
// start，有序集合索引起始位置（非分数）
//  end，有序集合索引结束位置（非分数）
func GetSortedSetByIndexByRange(key string, start int64, stop int64) []string {
	return redisClient.ZRange(getRealKey(key), start, stop).Val()
}

// 根据Score 从start到end获取有序数组的值
func GetSortedSetByScoresByRange(key string, opt redis.ZRangeBy) []string {
	return redisClient.ZRangeByScore(getRealKey(key), opt).Val()
}

//按照分数范围获取name对应的有序集合的元素, 其他同上 start.num 在结果集中取分页
func GetSortedSetByScoresByRangeDesc(key string, opt redis.ZRangeBy) []string {
	return redisClient.ZRevRangeByScore(getRealKey(key), opt).Val()
}

// 查看分数范围内的值的个数
func GetSortedSetValueCountByRange(key string, min string, max string) int64 {
	return redisClient.ZCount(getRealKey(key), min, max).Val()
}

// 自增
func SortedSetIncr(key string, increment float64, member string) {
	redisClient.ZIncrBy(getRealKey(key), increment, member)
}

// 自减
func SortedSetDecr(key string, increment float64, member string) {
	redisClient.ZIncrBy(getRealKey(key), increment, member)
}

func GetSortedSetIndexByValue(key string, member string) {
	redisClient.ZRank(key, member)
}

func RemoveSortedSetValue(key string, members ...interface{}) {
	redisClient.ZRem(getRealKey(key), members)
}

func GetScoreByValue(key, member string) float64 {
	return redisClient.ZScore(getRealKey(key), member).Val()
}
