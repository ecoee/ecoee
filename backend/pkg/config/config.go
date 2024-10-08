package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
)

const (
	mongoDBHost            = "MONGO_DB_HOST"
	mongoDBUserName        = "MONGO_DB_USER_NAME"
	mongoDBPassword        = "MONGO_DB_PASSWORD" // #nosec
	gcpAPIKey              = "GCP_API_KEY"
	vertexAIProjectID      = "VERTEX_AI_PROJECT_ID"
	vertexAILocation       = "VERTEX_AI_LOCATION"
	cloudStorageBucketName = "CLOUD_STORAGE_BUCKET_NAME"
	storage                = "STORAGE"
	vertexAI               = "VERTEX"
)

type Config struct {
	MongoDBConfig MongoDBConfig `json:"mongo_db_config"`
	GCPConfig     GCPConfig     `json:"gcp_config"`
}

type MongoDBConfig struct {
	MongoDBHost     string `json:"mongo_db_host"`
	MongoDBPort     string `json:"mongo_db_port"`
	MongoDBUserName string `json:"mongo_db_user_name"`
	MongoDBPassword string `json:"mongo_db_password"`
}

type GCPConfig struct {
	ProjectID          string             `json:"project_id"`
	Location           string             `json:"location"`
	APIKey             string             `json:"api_key"`
	CloudStorageConfig CloudStorageConfig `json:"cloud_storage_config"`
	Storage            string             `json:"storage"`
	Vertex             string             `json:"vertex"`
}

type CloudStorageConfig struct {
	BucketName string `json:"bucket_name"`
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
		GCPConfig: GCPConfig{
			ProjectID: v.GetString(vertexAIProjectID),
			Location:  v.GetString(vertexAILocation),
			APIKey:    v.GetString(gcpAPIKey),
			CloudStorageConfig: CloudStorageConfig{
				BucketName: v.GetString(cloudStorageBucketName),
			},
			Storage: v.GetString(storage),
			Vertex:  v.GetString(vertexAI),
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
	v.SetDefault(gcpAPIKey, "AIzaSyCjZcxY4Q1AdDgFxR83e5j6cgfrP4duz_o")
	v.SetDefault(vertexAIProjectID, "ecoee-433110")
	v.SetDefault(vertexAILocation, "asia-northeast3")
	v.SetDefault(cloudStorageBucketName, "ecoee-assessment")
}

func bindEnvironment(v *viper.Viper) {
	// mongoDB
	_ = v.BindEnv(mongoDBHost)
	_ = v.BindEnv(mongoDBUserName)
	_ = v.BindEnv(mongoDBPassword)

	// GCP
	_ = v.BindEnv(gcpAPIKey)
	_ = v.BindEnv(vertexAIProjectID)
	_ = v.BindEnv(vertexAILocation)
	_ = v.BindEnv(cloudStorageBucketName)
	_ = v.BindEnv(storage)
	_ = v.BindEnv(vertexAI)
}
