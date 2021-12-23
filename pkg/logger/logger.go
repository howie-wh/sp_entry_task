package logger

import (
	"entry_task/pkg/config"
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// Log 全局对象
var (
	Log *logrus.Logger
)

// InitLogger 初始化日志模块
func InitLogger() {
	logFilePath := config.GetLoggerCfg().LogPath
	logFileName := config.GetLoggerCfg().LogName
	logLevel := config.GetLoggerCfg().LogLevel

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if !gotool.FileUtils.Exists(fileName) {
		create, err := os.Create(fileName)
		if err != nil {
			gotool.Logs.ErrorLog().Println(err)
		}
		defer create.Close()
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	Log = logrus.New()

	// 设置输出
	Log.Out = src

	// 设置日志级别
	SetLogLevel(logLevel)

	// 为设置不同级别日志输出到不同的目的
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  SetRotateLogs(fileName+".debug"),
		logrus.DebugLevel: SetRotateLogs(fileName+".info"),
		logrus.WarnLevel:  SetRotateLogs(fileName+".warn"),
		logrus.ErrorLevel: SetRotateLogs(fileName+".error"),
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	Log.AddHook(lfHook)
	Log.SetReportCaller(true)
}

func SetRotateLogs(fileName string) *rotatelogs.RotateLogs {
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
		return nil
	}
	return logWriter
}

func SetLogLevel(logLevel string) {
	// 设置日志级别
	if Log == nil {
		return
	}
	if logLevel == "debug" {
		Log.SetLevel(logrus.DebugLevel)
	} else if logLevel == "info" {
		Log.SetLevel(logrus.InfoLevel)
	} else if logLevel == "warn" {
		Log.SetLevel(logrus.WarnLevel)
	} else {
		Log.SetLevel(logrus.ErrorLevel)
	}
}

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.GetLoggerCfg().LogPath
	logFileName := config.GetLoggerCfg().LogName
	logLevel := config.GetLoggerCfg().LogLevel

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if !gotool.FileUtils.Exists(fileName) {
		create, err := os.Create(fileName)
		if err != nil {
			gotool.Logs.ErrorLog().Println(err)
		}
		defer create.Close()
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	logger.SetReportCaller(true)
	// 设置日志级别
	if logLevel == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else if logLevel == "info" {
		logger.SetLevel(logrus.InfoLevel)
	} else if logLevel == "warn" {
		logger.SetLevel(logrus.WarnLevel)
	} else {
		logger.SetLevel(logrus.ErrorLevel)
	}

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.debug.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 新增 Hook
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Debug()
	}
}
