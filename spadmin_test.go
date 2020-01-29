package spadmin

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"google.golang.org/api/option"

	"github.com/gcpug/handy-spanner/fake"
)

const (
	testDatabaseName = "spadmin_test"
)

func TestCreateDatabase(t *testing.T) {
	dsn := os.Getenv("SPADMIN_SPANNER_DSN")
	if dsn == "" {
		t.Skip("env not set: SPADMIN_SPANNER_DSN")
	}

	ctx := context.Background()
	c, err := NewClient(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}

	// invalid database not found
	{
		exists, err := c.DatabaseExists(ctx, "invalid-db")
		if err != nil {
			t.Fatal(err)
		}
		assert.False(t, exists)
	}

	// create database & database exists
	{
		if err := c.CreateDatabase(ctx, testDatabaseName, nil); err != nil {
			t.Fatal(err)
		}
		exists, err := c.DatabaseExists(ctx, testDatabaseName)
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, exists)
	}

	// drop database
	{
		if err := c.DropDatabase(ctx, testDatabaseName); err != nil {
			t.Fatal(err)
		}
		exists, err := c.DatabaseExists(ctx, testDatabaseName)
		if err != nil {
			t.Fatal(err)
		}
		if exists {
			t.Fatalf("db: '%s' expected not found, but found", testDatabaseName)
		}
	}
}

func TestCreateDatabaseWithHandySpanner(t *testing.T) {
	dsn := "projects/fake/instances/fake"
	srv, conn, err := fake.Run()
	if err != nil {
		t.Fatal(err)
	}
	srv.Addr()
	defer srv.Stop()

	ctx := context.Background()
	c, err := NewClient(ctx, dsn, option.WithGRPCConn(conn))
	if err != nil {
		t.Fatal(err)
	}

	// invalid database not found
	{
		exists, err := c.DatabaseExists(ctx, "invalid-db")
		if err != nil {
			t.Fatal(err)
		}
		assert.False(t, exists)
	}

	// create database & database exists
	{
		if err := c.CreateDatabase(ctx, testDatabaseName, nil); err != nil {
			t.Fatal(err)
		}
		exists, err := c.DatabaseExists(ctx, testDatabaseName)
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, exists)
	}

	// drop database
	{
		if err := c.DropDatabase(ctx, testDatabaseName); err != nil {
			t.Fatal(err)
		}
		exists, err := c.DatabaseExists(ctx, testDatabaseName)
		if err != nil {
			t.Fatal(err)
		}
		if exists {
			t.Fatalf("db: '%s' expected not found, but found", testDatabaseName)
		}
	}
}
