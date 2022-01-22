package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
)

type DbInfo struct {
	TYPE     string
	USER     string
	PASSWORD string
	DB_HOST  string
	NAME     string
}

type ServerInfo struct {
	HTTP_PORT string
}

func GetServerConfig() *ServerInfo {
	cfg, err := ini.Load("configFile/config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v \n", err)
		os.Exit(1)
	}
	d := new(ServerInfo)
	_ = cfg.Section("server").MapTo(d)
	return d
}

func GetDbConfig() *DbInfo {
	cfg, err := ini.Load("configFile/config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v \n", err)
		os.Exit(1)
	}
	d := new(DbInfo)
	_ = cfg.Section("database").MapTo(d)
	return d
}

func GetLogPath() string {
	timeObj := time.Now()
	datetime := timeObj.Format("2006-01-02-15-04-05")
	return "log/Course_Select" + datetime + ".log"
}

func GetLogFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
