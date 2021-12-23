package main

import (
	"entry_task/admin/constant"
	"entry_task/admin/router"
	"entry_task/pkg/config"
	"entry_task/pkg/dao"
	"entry_task/pkg/jwt"
	"entry_task/pkg/logger"
	"flag"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var (
	port, mode string
	configPath string = "./admin/config/config-prod.ini"
)
func InitFilter() {
	whiteList := constant.GetWhiteList()
	for i := 0; i < len(whiteList); i++ {
		jwt.AddWhiteList(whiteList[i])
	}
	blackList := constant.GetBlackList()
	for i := 0; i < len(blackList); i++ {
		jwt.AddBlackList(blackList[i])
	}
}

func InitGin() {
	// 设置gin模式
	gin.SetMode(mode)

	// 设置gin日志输出到文件
	logFilePath := config.GetLoggerCfg().LogPath
	fileName := path.Join(logFilePath, "gin.log")
	if !gotool.FileUtils.Exists(fileName) {
		create, err := os.Create(fileName)
		if err != nil {
			logger.Log.WithFields(logrus.Fields{}).Error(err)
			return
		}
		defer create.Close()
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(src)
}

func init() {
	flag.StringVar(&port, "port", "3000", "server listening on, default 3000")
	flag.StringVar(&mode, "mode", "debug", "server running mode, default debug mode")
	// 配置初始化
	config.InitConfig(configPath)
	// 日志初始化
	logger.InitLogger()
	// 数据库初始化
	dao.InitDao()
	// 过滤条件初始化
	InitFilter()
	// 日志初始化
	logger.InitLogger()
	// gin初始化
	gin.SetMode(mode)
	//InitGin()
}

func main() {
	port := config.GetServerCfg().Port
	flag.Parse()

	r := router.Init()
	err := r.Run(port)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{}).Error(err)
	}
}
