// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hive_metastore

import (
	"bytes"
	"context"
	"fmt"
	"reflect"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

const DDL_TIME = "transient_lastDdlTime"
const HIVE_FILTER_FIELD_OWNER = "hive_filter_field_owner__"
const HIVE_FILTER_FIELD_PARAMS = "hive_filter_field_params__"
const HIVE_FILTER_FIELD_LAST_ACCESS = "hive_filter_field_last_access__"
const IS_ARCHIVED = "is_archived"
const ORIGINAL_LOCATION = "original_location"
const IS_IMMUTABLE = "immutable"
const META_TABLE_COLUMNS = "columns"
const META_TABLE_COLUMN_TYPES = "columns.types"
const BUCKET_FIELD_NAME = "bucket_field_name"
const BUCKET_COUNT = "bucket_count"
const FIELD_TO_DIMENSION = "field_to_dimension"
const META_TABLE_NAME = "name"
const META_TABLE_DB = "db"
const META_TABLE_LOCATION = "location"
const META_TABLE_SERDE = "serde"
const META_TABLE_PARTITION_COLUMNS = "partition_columns"
const META_TABLE_PARTITION_COLUMN_TYPES = "partition_columns.types"
const FILE_INPUT_FORMAT = "file.inputformat"
const FILE_OUTPUT_FORMAT = "file.outputformat"
const META_TABLE_STORAGE = "storage_handler"
const TABLE_IS_TRANSACTIONAL = "transactional"
const TABLE_NO_AUTO_COMPACT = "no_auto_compaction"
const TABLE_TRANSACTIONAL_PROPERTIES = "transactional_properties"

func init() {
}
