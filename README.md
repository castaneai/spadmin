# spadmin

:bangbang: [google cloud spanner admin package](https://godoc.org/cloud.google.com/go/spanner/admin/database/apiv1) is in alpha. It is not stable.

Cloud Spanner Admin client wrapper


## Works with [handy-spanner](https://github.com/gcpug/handy-spanner)

```go

import (
	"github.com/castaneai/spadmin"
	"github.com/gcpug/handy-spanner/fake"

	"google.golang.org/api/option"
)

func main() {
	dsn := "projects/fake/instances/fake"
	srv, conn, err := fake.Run()
	if err != nil {
		t.Fatal(err)
	}
	srv.Addr()
	defer srv.Stop()

	ctx := context.Background()
	admin, err := spadmin.NewClient(ctx, dsn, option.WithGRPCConn(conn))
	...
}
```
