package dao

import (
	"entry_task/pkg/common"
	"entry_task/pkg/config"
	"entry_task/pkg/logger"
	redisTool "entry_task/pkg/redistool"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"time"
)

// X 全局DB
var (
	SqlDB   *xorm.Engine
	RedisDB *redisTool.RedisClient
)

func InitDao() {
	var err error
	//配置mysql数据库
	mysql := config.GetMysqlCfg()
	jdbc := mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Database + "?charset=utf8&parseTime=True&loc=Local"
	SqlDB, _ = xorm.NewEngine(mysql.ShowType, jdbc)

	if err != nil {
		//log.Fatalf("db error: %#v\n", err.Error())
		logger.Log.WithFields(logrus.Fields{}).Error("db error: %#v\n", err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		//log.Fatalf("db connect error: %#v\n", err.Error())
		logger.Log.WithFields(logrus.Fields{}).Error("db connect error: %#v\n", err.Error())
	}
	SqlDB.SetMaxIdleConns(mysql.MaxIdleConnection)
	SqlDB.SetMaxOpenConns(mysql.MaxOpenConnection)
	_ = SqlDB.Sync2(
	//new(model.User),
	)
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				//log.Fatalf("db connect error: %#v\n", err.Error())
				logger.Log.WithFields(logrus.Fields{}).Error("db connect error: %#v\n", err.Error())
			}
		}
	}(SqlDB)
	SqlDB.ShowSQL(true)
	//初始化redis开始
	redisCfg := config.GetRedisCfg()
	redisOpt := common.RedisConnOpt{
		true,
		redisCfg.RedisHost,
		int32(redisCfg.Port),
		redisCfg.RedisPwd,
		int32(redisCfg.RedisDB),
		240,
	}

	RedisDB = redisTool.NewRedis(redisOpt)
	//配置redis结束
}