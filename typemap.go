package dgwlib

// DefaultTypeMapCfg is the default type map configuration
var DefaultTypeMapCfg PgTypeMapConfig = PgTypeMapConfig{
	"string": TypeMap{
		DBTypes:        []string{"character", "character varying", "text", "money"},
		NotNullGoType:  "string",
		NullableGoType: "sql.NullString",
	},
	"time": TypeMap{
		DBTypes:        []string{"time with time zone", "time without time zone", "timestamp without time zone", "timestamp with time zone", "date"},
		NotNullGoType:  "time.Time",
		NullableGoType: "sql.NullTime",
	},
	"bool": TypeMap{
		DBTypes:        []string{"boolean"},
		NotNullGoType:  "bool",
		NullableGoType: "sql.NullBool",
	},
	"smallint": TypeMap{
		DBTypes:        []string{"smallint"},
		NotNullGoType:  "int16",
		NullableGoType: "sql.NullInt64",
	},
	"integer": TypeMap{
		DBTypes:        []string{"integer"},
		NotNullGoType:  "int",
		NullableGoType: "sql.NullInt64",
	},
	"bigint": TypeMap{
		DBTypes:        []string{"bigint"},
		NotNullGoType:  "int64",
		NullableGoType: "sql.NullInt64",
	},
	"smallserial": TypeMap{
		DBTypes:        []string{"smallserial"},
		NotNullGoType:  "uint16",
		NullableGoType: "sql.NullInt64",
	},
	"serial": TypeMap{
		DBTypes:        []string{"serial"},
		NotNullGoType:  "uint32",
		NullableGoType: "sql.NullInt64",
	},
	"real": TypeMap{
		DBTypes:        []string{"real"},
		NotNullGoType:  "float32",
		NullableGoType: "sql.NullFloat64",
	},
	"numeric": TypeMap{
		DBTypes:        []string{"numeric", "double precision"},
		NotNullGoType:  "float64",
		NullableGoType: "sql.NullFloat64",
	},
	"bytea": TypeMap{
		DBTypes:        []string{"bytea"},
		NotNullGoType:  "[]byte",
		NullableGoType: "[]byte",
	},
	"json": TypeMap{
		DBTypes:        []string{"json", "jsonb"},
		NotNullGoType:  "[]byte",
		NullableGoType: "[]byte",
	},
	"xml": TypeMap{
		DBTypes:        []string{"xml"},
		NotNullGoType:  "[]byte",
		NullableGoType: "[]byte",
	},
	"interval": TypeMap{
		DBTypes:        []string{"interval"},
		NotNullGoType:  "time.Duration",
		NullableGoType: "*time.Duration",
	},
	"default": TypeMap{
		DBTypes:        []string{"*"},
		NotNullGoType:  "interface{}",
		NullableGoType: "interface{}",
	},
}
