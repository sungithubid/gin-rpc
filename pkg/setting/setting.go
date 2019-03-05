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

	AppConfig.AppMode = configFile.Section("app").Key("app_mode").In("prod", []string{"prod", "dev", "debug"})
	AppConfig.Protocol = configFile.Section("server").Key("protocol").In("http", []string{"http", "https", "tcp"})
	AppConfig.HTTPPort = configFile.Section("server").Key("http_port").MustInt()
	AppConfig.TCPPort = configFile.Section("server").Key("tcp_port").MustInt()
}
