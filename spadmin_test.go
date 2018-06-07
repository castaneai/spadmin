package spadmin

import (
	"context"
	"os"
	"testing"
)

const (
	testDatabaseName = "test-db"
)

func TestCreateDatabase(t *testing.T) {
	dsn := os.Getenv("SPADMIN_SPANNER_DSN")
	if dsn == "" {
		t.Skip("env not set: SPADMIN_SPANNER_DSN")
	}

	c, err := NewClient(dsn)
	if err != nil {
		t.Error(err)
	}

	ctx := context.Background()
	defer (func() {
		err := c.DropDatabase(ctx, testDatabaseName)
		if err != nil {
			t.Error(err)
		}
	})()

	if err := c.CreateDatabase(ctx, testDatabaseName, nil); err != nil {
		t.Error(err)
	}

	iexists, err := c.DatabaseExists(ctx, "invalid-db")
	if err != nil {
		t.Error(err)
	}
	if iexists {
		t.Error("db: 'invalid-db' must be not found, but found")
	}

	exists, err := c.DatabaseExists(ctx, testDatabaseName)
	if err != nil {
		t.Error(err)
	}
	if !exists {
		t.Errorf("db: '%s' must exists", testDatabaseName)
	}
}
