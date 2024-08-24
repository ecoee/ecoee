package mongo

import (
	"context"
	"ecoee/pkg/config"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"log/slog"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	appName     = "ecoee"
	ecoeeDB     = "ecoee"
	ecoeeTestDB = "ecoee_test"
	maxPoolSize = 20
	minPoolSize = 10
)

func NewDB(ctx context.Context, config config.Config) (*mongo.Database, error) {
	client, err := newDBClient(ctx, config)
	if err != nil {
		return nil, err
	}

	db := client.Database(ecoeeDB)
	return db, nil
}

func NewTestDB(ctx context.Context, config config.Config) *mongo.Database {
	client, _ := newDBClient(ctx, config)
	db := initTestDB(ctx, client)
	return db
}

func initTestDB(ctx context.Context, client *mongo.Client) *mongo.Database {
	db := client.Database(ecoeeTestDB)
	if err := db.Drop(ctx); err != nil {
		slog.Error(fmt.Sprintf("error while dropping testing DB: %v", err))
	}

	return db
}

func newDBClient(ctx context.Context, config config.Config) (*mongo.Client, error) {
	credential := options.Credential{
		Username: config.MongoDBConfig.MongoDBUserName,
		Password: config.MongoDBConfig.MongoDBPassword,
	}
	connectionURI := getConnectionURI(config)
	slog.Info(fmt.Sprintf("mongoDB connectionURI=%s", connectionURI))

	registryBuilder := bson.NewRegistry()
	registryBuilder.RegisterTypeMapEntry(bson.TypeEmbeddedDocument, reflect.TypeOf(bson.M{}))

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(connectionURI).
		SetAuth(credential).
		SetRetryWrites(false).
		SetAppName(appName).
		SetMinPoolSize(minPoolSize).
		SetMaxPoolSize(maxPoolSize).
		SetRegistry(registryBuilder))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open mongoDB client")
	}

	return client, nil
}

func getConnectionURI(config config.Config) string {
	return fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		config.MongoDBConfig.MongoDBUserName,
		config.MongoDBConfig.MongoDBPassword,
		config.MongoDBConfig.MongoDBHost,
	)
}
