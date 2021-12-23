package common

// RedisConnOpt connect redis options
type RedisConnOpt struct {
	Enable   bool
	Host     string
	Port     int32
	Password string
	Index    int32
	TTL      int32
}

// RedisData 存储数据结构
type RedisData struct {
	Key    string
	Field  string
	Value  string
	Expire int64
}

type RedisDataArray []*RedisData