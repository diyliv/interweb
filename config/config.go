package config

import "github.com/spf13/viper"

type Config struct {
	Postgres Postgres
	Telegram Telegram
}

type Postgres struct {
	Host            string
	Port            string
	Login           string
	Password        string
	DB              string
	ConnMaxLifeTime int
	MaxOpenConn     int
	MaxIdleConn     int
}

type Telegram struct {
	Token string
}

func ReadConfig(cfgName, cfgType, cfgPath string) *Config {
	var cfg Config

	viper.SetConfigName(cfgName)
	viper.SetConfigType(cfgType)
	viper.AddConfigPath(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
