package frostdbexporter

import (
	"context"

	"github.com/polarsignals/frostdb"
	schemapb "github.com/polarsignals/frostdb/gen/proto/go/frostdb/schema/v1alpha2"
)

const (
	dbName    = "otel-logs"
	tableName = "log_records"
)

func createColumnStore(ctx context.Context) (*frostdb.ColumnStore, error) {
	columnstore, err := frostdb.New()
	if err != nil {
		return nil, err
	}

	// TODO -- make this configurable?
	database, err := columnstore.DB(ctx, "otel-logs")
	if err != nil {
		return nil, err
	}

	schema := &schemapb.Schema{
		Root: &schemapb.Group{
			Name: tableName,
			Nodes: []*schemapb.Node{
				{
					Type: &schemapb.Node_Leaf{
						Leaf: &schemapb.Leaf{
							Name: "bodystring",
							StorageLayout: &schemapb.StorageLayout{
								Type:     schemapb.StorageLayout_TYPE_STRING,
								Encoding: schemapb.StorageLayout_ENCODING_PLAIN_UNSPECIFIED,
							},
						},
					},
				},
			},
		},
	}

	database.Table(tableName, frostdb.NewTableConfig(schema))

	return columnstore, nil

}
