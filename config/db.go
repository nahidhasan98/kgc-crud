package config

import (
	"context"
	"time"

	"github.com/nahidhasan98/kgc-crud/errorhandling"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Connect function for connenting to DB
func DBConnect() (*mongo.Database, context.Context, context.CancelFunc) {
	dbName := databaseName
	dbConnectionString := databaseConnectionString

	dbClient, err := mongo.NewClient(options.Client().ApplyURI(dbConnectionString))
	errorhandling.Check(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = dbClient.Connect(ctx)
	errorhandling.Check(err)

	err = dbClient.Ping(ctx, readpref.Primary())
	errorhandling.Check(err)

	//return db
	return dbClient.Database(dbName), ctx, cancel
}
