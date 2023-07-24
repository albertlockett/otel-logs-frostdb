package frostdbexporter

import (
	"context"
	"log"

	"github.com/polarsignals/frostdb"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
)

type frostdbExporter struct {
	columnstore *frostdb.ColumnStore
}

func (f *frostdbExporter) Start(ctx context.Context, host component.Host) error {
	// TODO
	return nil
}

func (f *frostdbExporter) Shutdown(ctx context.Context) error {
	// TODO
	return nil
}

func (f *frostdbExporter) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{
		MutatesData: false,
	}
}

type Record struct {
	Bodystring string
}

func (f *frostdbExporter) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	log.Printf("log received")

	// database, err := f.columnstore.DB(ctx, dbName)
	// if err != nil {
	// 	return err
	// }

	// log.Printf("got database %v", database)

	// table, err := database.GetTable(tableName)
	// if err != nil {
	// 	return err
	// }

	// log.Printf("got table %v", table)

	// for i := 0; i < logs.ResourceLogs().Len(); i++ {
	// 	resourceLogs := logs.ResourceLogs().At(i)
	// 	for j := 0; j < resourceLogs.ScopeLogs().Len(); j++ {
	// 		scopedLogs := resourceLogs.ScopeLogs().At(j)
	// 		for k := 0; k < scopedLogs.LogRecords().Len(); k++ {
	// 			logRecord := scopedLogs.LogRecords().At(k)
	// 			record := Record{
	// 				Bodystring: logRecord.Body().AsString(),
	// 			}

	// 			table.Write(ctx, record)
	// 		}
	// 	}
	// }

	return nil
}
