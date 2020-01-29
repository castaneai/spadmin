package spadmin

import (
	"context"
	"fmt"

	"google.golang.org/api/option"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type client struct {
	admin *database.DatabaseAdminClient
	dsn   string
}

// NewClient creates new spadmin client for Cloud Spanner
// dsn format: 'projects/<GCP_PROJECT_ID>/instances/<SPANNER_INSTANCE_ID>'
// @see https://cloud.google.com/spanner/docs/reference/rpc/google.spanner.admin.database.v1#google.spanner.admin.database.v1.CreateDatabaseMetadata
func NewClient(ctx context.Context, dsn string, opts ...option.ClientOption) (*client, error) {
	admin, err := database.NewDatabaseAdminClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &client{
		admin: admin,
		dsn:   dsn,
	}, nil
}

func (c *client) DatabaseExists(ctx context.Context, name string) (bool, error) {
	req := &adminpb.GetDatabaseRequest{
		Name: fmt.Sprintf("%s/databases/%s", c.dsn, name),
	}
	if _, err := c.admin.GetDatabase(ctx, req); err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (c *client) CreateDatabase(ctx context.Context, name string, statements []string) error {
	req := &adminpb.CreateDatabaseRequest{
		Parent:          c.dsn,
		CreateStatement: fmt.Sprintf("CREATE DATABASE `%s`", name),
		ExtraStatements: statements,
	}
	op, err := c.admin.CreateDatabase(ctx, req)
	if err != nil {
		return err
	}

	if _, err := op.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (c *client) DropDatabase(ctx context.Context, name string) error {
	req := &adminpb.DropDatabaseRequest{
		Database: fmt.Sprintf("%s/databases/%s", c.dsn, name),
	}
	return c.admin.DropDatabase(ctx, req)
}
