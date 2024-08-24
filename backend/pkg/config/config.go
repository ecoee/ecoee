package config

import (
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

const (
	mongoDBHost       = "MONGO_DB_HOST"
	mongoDBUserName   = "MONGO_DB_USER_NAME"
	mongoDBPassword   = "MONGO_DB_PASSWORD" // #nosec
	vertexAIProjectID = "VERTEX_AI_PROJECT_ID"
	vertexAILocation  = "VERTEX_AI_LOCATION"
)

type Config struct {
	MongoDBConfig  MongoDBConfig  `json:"mongo_db_config"`
	VertexAIConfig VertexAIConfig `json:"vertex_ai_config"`
}

type MongoDBConfig struct {
	MongoDBHost     string `json:"mongo_db_host"`
	MongoDBPort     string `json:"mongo_db_port"`
	MongoDBUserName string `json:"mongo_db_user_name"`
	MongoDBPassword string `json:"mongo_db_password"`
}

type VertexAIConfig struct {
	ProjectID string `json:"project_id"`
	Location  string `json:"location"`
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
		VertexAIConfig: VertexAIConfig{
			ProjectID: v.GetString(vertexAIProjectID),
			Location:  v.GetString(vertexAILocation),
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
	v.SetDefault(vertexAIProjectID, "ecoee-433110")
	v.SetDefault(vertexAILocation, "asia-northeast3")
}

func bindEnvironment(v *viper.Viper) {
	// mongoDB
	_ = v.BindEnv(mongoDBHost)
	_ = v.BindEnv(mongoDBUserName)
	_ = v.BindEnv(mongoDBPassword)

	// vertexAI
	_ = v.BindEnv(vertexAIProjectID)
	_ = v.BindEnv(vertexAILocation)
}
