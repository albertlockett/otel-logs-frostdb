package frostdbexporter

import (
	"fmt"

	schemapb "github.com/polarsignals/frostdb/gen/proto/go/frostdb/schema/v1alpha1"
)

// TODO clean this up this is messy
func anySchemaInternal(fieldName string, dynamic bool) []*schemapb.Column {
	return []*schemapb.Column{
		{
			Name: fmt.Sprintf("%s_string", fieldName),
			StorageLayout: &schemapb.StorageLayout{
				Type: schemapb.StorageLayout_TYPE_STRING,
			},
		},
		{
			Name: fmt.Sprintf("%s_bool", fieldName),
			StorageLayout: &schemapb.StorageLayout{
				Type: schemapb.StorageLayout_TYPE_BOOL,
			},
		},
		{
			Name: fmt.Sprintf("%s_int", fieldName),
			StorageLayout: &schemapb.StorageLayout{
				Type: schemapb.StorageLayout_TYPE_INT64,
			},
		},
	}
}

func anySchema(fieldName string) []*schemapb.Column {
	return anySchemaInternal(fieldName, false)
}

func keyValueSchema(fieldName string) []*schemapb.Column {
	return anySchemaInternal(fmt.Sprintf("%s_keyvalue", fieldName), true)
}
