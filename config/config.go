package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig(file string) *Config {
	viper.SetConfigFile(file)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("load config error: %v", err)
	}

	viper.SetDefault("app.name", "Exam API")
	viper.SetDefault("server.port", ServerPort)

	viper.SetDefault("db.main.host", MariadbHost)
	viper.SetDefault("db.main.port", MariadbPort)

	cfg := &Config{
		App: appConfig{
			Name:       viper.GetString("app.name"),
			Mode:       viper.GetString("app.mode"),
			PrefixPath: viper.GetString("app.prefix.path"),
		},
		Server: serverCfg{
			Port: viper.GetString("server.port"),
		},
		DBMain: MariaDBConfig{
			Host:      viper.GetString("db.main.host"),
			Port:      viper.GetString("db.main.port"),
			User:      viper.GetString("db.main.username"),
			Password:  viper.GetString("db.main.password"),
			Database:  viper.GetString("db.main.database"),
			Migration: viper.GetBool("db.main.migration"),
		},
	}

	return cfg
}
