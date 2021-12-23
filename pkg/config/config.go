package config

import (
	"github.com/Unknwon/goconfig"
	"github.com/druidcaesa/gotool"
	"log"
	"strconv"
	"time"
)

var Cfg *goconfig.ConfigFile

func InitConfig(filename string) {
	var err error
	Cfg, err = goconfig.LoadConfigFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return
}

type JwtConfig struct {
	TimeOut time.Duration //超时时间
	Issuer  string        //签证签发人
}

// FilePath 文件存储配置获取
type FilePath struct {
	Path string
}

// DbCfg Mysql数据库配置
type DbCfg struct {
	Username          string
	Password          string
	Database          string
	Host              string
	Port              string
	MaxIdleConnection int
	MaxOpenConnection int
	ShowType          string
}

// AppServer 应用程序配置
type AppServer struct {
	Port        string //App运行端口
}

// LoggerCfg 日志配置结构体
type LoggerCfg struct {
	LogPath string
	LogName string
	LogLevel string
}

// RedisCfg redis配置结构体
type RedisCfg struct {
	RedisHost string //地址
	Port      int64  //端口
	RedisPwd  string //密码
	RedisDB   int64  //数据库
	Timeout   int64  //超时时间
}

// GetMysqlCfg 读取mysql配置
func GetMysqlCfg() (mysql DbCfg) {
	mysql.Username, _ = Cfg.GetValue("mysql", "username")
	mysql.Password, _ = Cfg.GetValue("mysql", "password")
	mysql.Database, _ = Cfg.GetValue("mysql", "database")
	mysql.Host, _ = Cfg.GetValue("mysql", "host")
	mysql.Port, _ = Cfg.GetValue("mysql", "port")
	mysql.ShowType, _ = Cfg.GetValue("mysql", "sqlType")
	value, _ := Cfg.GetValue("mysql", "MaxIdleConnection")
	mysql.MaxIdleConnection, _ = strconv.Atoi(value)
	v, _ := Cfg.GetValue("mysql", "MaxOpenConnection")
	mysql.MaxOpenConnection, _ = strconv.Atoi(v)
	return mysql
}

// GetServerCfg 读取server配置
func GetServerCfg() (server AppServer) {
	server.Port, _ = Cfg.GetValue("app", "server")
	return server
}

// GetLoggerCfg 获取Logger配置
func GetLoggerCfg() (logger LoggerCfg) {
	logger.LogPath, _ = Cfg.GetValue("logger", "logPath")
	logger.LogName, _ = Cfg.GetValue("logger", "logName")
	logger.LogLevel, _ = Cfg.GetValue("logger", "logLevel")
	return logger
}

//GetRedisCfg 获取redis配置
func GetRedisCfg() (r RedisCfg) {
	r.RedisHost, _ = Cfg.GetValue("redis", "host")
	getValue, _ := Cfg.GetValue("redis", "port")
	r.Port, _ = strconv.ParseInt(getValue, 10, 32)
	r.RedisPwd, _ = Cfg.GetValue("redis", "password")
	db, _ := Cfg.GetValue("redis", "db")
	r.RedisDB, _ = strconv.ParseInt(db, 10, 32)
	value, _ := Cfg.GetValue("redis", "timeout")
	r.Timeout, _ = strconv.ParseInt(value, 10, 64)
	return r
}

// GetFilePath 获取文件存储位置
func GetFilePath() (g FilePath) {
	g.Path, _ = Cfg.GetValue("filePath", "path")
	return g
}

func GetJwtConfig() (j JwtConfig) {
	value, _ := Cfg.GetValue("jwt", "timeOut")
	issuer, _ := Cfg.GetValue("jwt", "issuer")
	if gotool.StrUtils.HasEmpty(value) {
		value = "60"
	}
	if gotool.StrUtils.HasEmpty(issuer) {
		value = "admin"
	}
	atoi, _ := strconv.ParseInt(value, 10, 64)
	j.TimeOut = time.Duration(atoi / 60)
	j.Issuer = issuer
	return j
}
