package utils

import (
	"fmt"
)

func SelectSQL(outputField, findingField string) string {
	query := fmt.Sprintf("SELECT %v FROM %v WHERE %v = $1", outputField, DB_TABLE_NAME, findingField)
	return query
}

func InsertSQL(field1, field2 string) string {
	query := fmt.Sprintf("INSERT INTO %v (%v, %v) VALUES ($1, $2)", DB_TABLE_NAME, field1, field2)
	return query
}
