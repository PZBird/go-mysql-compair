package model

import "database/sql"

type DatabaseSchema struct {
	SchemaName string
	Tables     map[string]*Table
}

type Table struct {
	TableName    string
	Columns      []*Column
	PrimaryKeys  []*Column
	OtherColumns []*Column
}

type View struct {
	ViewName string
	Columns  []*Column
}

type Column struct {
	TableName              string
	ColumnName             string
	IsPrimaryKey           bool
	IsNullable             bool
	IsAutoIncrement        bool
	IsUnique               bool
	DataType               string
	CharacterMaximumLength sql.NullInt64
	NumericPrecision       sql.NullInt64
	NumericScale           sql.NullInt64
	ColumnType             string
	DefaultValue           string
	EnumValues             []string
}
