package config

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config *config

type config struct {
	Port int
	DB   struct {
		PostgreSql struct {
			Host         string
			Port         int
			DatabaseName string
			User         string
			Password     string
			TimeZone     string
		}
	}
}

func init() {
	Config = &config{}
	viper.AddConfigPath("./config")
	configFile := os.Getenv("LINX_CONFIG")
	viper.SetConfigName(configFile)
	log.Println(configFile)
	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Warningf("해당 환경:%s의 설정파일이 없습니다.", configFile)
		} else {
			log.Fatal(err)
		}
	} else {
		log.Infof("%s 설정 파일을 발견했습니다.", configFile)
		if err := viper.Unmarshal(Config); err != nil {
			log.Fatal(err)
		}
	}
	log.Info(Config)
}
