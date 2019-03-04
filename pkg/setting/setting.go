package setting

import (
	"log"

	"github.com/go-ini/ini"
)

type config struct {
	AppMode  string
	Protocol string
	HTTPPort int
	TCPPort  int
}

// AppConfig setting info
var AppConfig config
var configFile *ini.File

// Load config file
func Load() {
	var err error
	configFile, err = ini.Load("configs/main.ini")
	if err != nil {
		log.Fatalf("setting.Load, fail to parse 'configs/main.ini': %v", err)
	}

	AppConfig.AppMode = configFile.Section("").Key("app_mode").MustString("dev")
	AppConfig.Protocol = configFile.Section("server").Key("protocol").MustString("http")
	AppConfig.HTTPPort = configFile.Section("port").Key("http_port").MustInt(8999)
	AppConfig.TCPPort = configFile.Section("port").Key("tcp_port").MustInt(1314)
}
