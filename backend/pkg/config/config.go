package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
)

const (
	mongoDBHost     = "MONGO_DB_HOST"
	mongoDBUserName = "MONGO_DB_USER_NAME"
	mongoDBPassword = "MONGO_DB_PASSWORD" // #nosec
)

type Config struct {
	MongoDBConfig MongoDBConfig `json:"mongo_db_config"`
}

type MongoDBConfig struct {
	MongoDBHost     string `json:"mongo_db_host"`
	MongoDBPort     string `json:"mongo_db_port"`
	MongoDBUserName string `json:"mongo_db_user_name"`
	MongoDBPassword string `json:"mongo_db_password"`
}

func NewConfig(v *viper.Viper) Config {
	setDefaults(v)
	bindEnvironment(v)

	return Config{
		MongoDBConfig: MongoDBConfig{
			MongoDBHost:     v.GetString(mongoDBHost),
			MongoDBUserName: v.GetString(mongoDBUserName),
			MongoDBPassword: v.GetString(mongoDBPassword),
		},
	}
}

func (c Config) Log() {
	slog.Info(fmt.Sprintf("config=%v", c))
}

func setDefaults(v *viper.Viper) {
	v.SetDefault(mongoDBHost, "ecoee.ykgcpvf.mongodb.net")
	v.SetDefault(mongoDBUserName, "ecoee")
	v.SetDefault(mongoDBPassword, "ecoee")
}

func bindEnvironment(v *viper.Viper) {
	// mongoDB
	_ = v.BindEnv(mongoDBHost)
	_ = v.BindEnv(mongoDBUserName)
	_ = v.BindEnv(mongoDBPassword)
}
