package setting

import (
	"time"
	"log"
	"github.com/go-ini/ini"
)

type App struct {
	RuntimePath string
	LogPath     string
}

type Sever struct {
	RunMode      string
	HttpAddr     string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Db struct {
	Type     string
	User     string
	Password string
	port     int
	Host     string
	Name     string
}

var (
	AppCfg    = new(App)
	ServerCfg = new(Sever)
	DbCfg     = new(Db)
)

func ConfigInit(mode *string) {
	var err error
	//配置文件
	var f *ini.File
	//传入环境
	if *mode == "pro" {
		f, err = ini.Load("conf/pro.ini")
		if err != nil {
			log.Fatal("pro config file load err:", err)
		}
	} else {
		f, err = ini.Load("conf/dev.ini")
		if err != nil {
			log.Fatal("dev config file load err:", err)
		}
	}
	//关系映射
	err = f.Section("app").MapTo(AppCfg)
	err = f.Section("server").MapTo(ServerCfg)
	err = f.Section("database").MapTo(DbCfg)
	if err != nil {
		log.Fatal("setting section to map err:", err)
	}
	ServerCfg.WriteTimeout = ServerCfg.WriteTimeout * time.Second
	ServerCfg.ReadTimeout = ServerCfg.ReadTimeout * time.Second
}
