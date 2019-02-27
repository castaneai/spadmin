package spadmin

import (
	"context"
	"os"
	"testing"
)

const (
	testDatabaseName = "spadmin_test"
)

func TestCreateDatabase(t *testing.T) {
	dsn := os.Getenv("SPADMIN_SPANNER_DSN")
	if dsn == "" {
		t.Skip("env not set: SPADMIN_SPANNER_DSN")
	}

	c, err := NewClient(dsn)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	defer (func() {
		err := c.DropDatabase(ctx, testDatabaseName)
		if err != nil {
			t.Fatal(err)
		}
	})()

	if err := c.CreateDatabase(ctx, testDatabaseName, nil); err != nil {
		t.Fatal(err)
	}

	iexists, err := c.DatabaseExists(ctx, "invalid-db")
	if err != nil {
		t.Fatal(err)
	}
	if iexists {
		t.Fatal("db: 'invalid-db' must be not found, but found")
	}

	exists, err := c.DatabaseExists(ctx, testDatabaseName)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatalf("db: '%s' must exists", testDatabaseName)
	}
}
