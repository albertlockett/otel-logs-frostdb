package frostdbexporter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/memory"
	"github.com/polarsignals/frostdb"
	schemapb "github.com/polarsignals/frostdb/gen/proto/go/frostdb/schema/v1alpha2"
	"github.com/polarsignals/frostdb/query"
	"github.com/polarsignals/frostdb/query/logicalplan"
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

func startQueryPoller(columnstore *frostdb.ColumnStore) {

	// TODO
	go func() {
		for {
			ctx := context.Background()
			database, _ := columnstore.DB(ctx, dbName)
			// table, _ := database.GetTable(tableName)

			database.TableProvider().GetTable(tableName)

			time.Sleep(1 * time.Second)

			log.Println("checking how many records in db")

			engine := query.NewEngine(memory.DefaultAllocator, database.TableProvider())
			engine.ScanTable(tableName).
				Aggregate(
					[]logicalplan.Expr{
						logicalplan.Count(logicalplan.Col("bodystring")),
					},
					[]logicalplan.Expr{},
				).
				Execute(ctx, func(ctx context.Context, r arrow.Record) error {
					fmt.Println(r)
					return nil
				})

		}
	}()
}
