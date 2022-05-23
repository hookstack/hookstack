package mongo

import (
	"context"
	"testing"

	"github.com/frain-dev/convoy/config"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
)

func getDSN() string {
	return "mongodb://localhost:27017/test_db"
	// return os.Getenv("TEST_MONGO_DSN")
}

func getConfig() config.Configuration {

	return config.Configuration{
		Database: config.DatabaseConfiguration{
			Type: config.MongodbDatabaseProvider,
			Dsn:  getDSN(),
		},
	}
}

func getDB(t *testing.T) (*mongo.Database, func()) {

	db, err := New(getConfig())
	require.NoError(t, err)

	return db.Client().(*mongo.Database), func() {
		require.NoError(t, db.Disconnect(context.Background()))
	}
}
