package constant

const (
	redisError     = "调用Redis发生异常------------------------->%s"
	redisDictKey   = "DICT_KEY_MONKEY_COOL"
	redisConfigKey = "CONFIG_KEY_MONKEY_COOL"
	mysqlError     = "调用Mysql发生异常------------------------->%s"
	redisTokenExpires = 3600
)

// RedisConstant Redis相关操作常量
type RedisConstant struct {
}

// GetRedisError Redis异常拼接常量
func (c RedisConstant) GetRedisError() string {
	return redisError
}

// GetRedisDictKey 获取Redis的Dict的key
func (c RedisConstant) GetRedisDictKey() string {
	return redisDictKey
}

// GetRedisConfigKey 获取redis的config的key
func (c RedisConstant) GetRedisConfigKey() string {
	return redisConfigKey
}

// GetRedisTokenExpires 获取redis的config的key
func (c RedisConstant) GetRedisTokenExpires() int {
	return redisTokenExpires
}

// MysqlConstant Mysql相关操作常量
type MysqlConstant struct {
}

// GetMysqlError Mysql异常拼接常量
func (c MysqlConstant) GetMysqlError() string {
	return mysqlError
}
