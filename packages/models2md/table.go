package models2md

import (
	"fmt"
	"github.com/datatug/datatug/packages/models"
	"io"
	"strings"
)

// EncodeTable encodes table summary to markdown file format
func (encoder) EncodeTable(w io.Writer, table *models.Table) error {
	var primaryKey string
	if table.PrimaryKey == nil {
		primaryKey = "*None*"
	} else {
		pkCols := make([]string, len(table.PrimaryKey.Columns))
		for i, pkCol := range table.PrimaryKey.Columns {
			pkCols[i] = fmt.Sprintf("**%v**", pkCol)
		}
		primaryKey = fmt.Sprintf("%v (%v)", table.PrimaryKey.Name, strings.Join(pkCols, ", "))
	}

	var foreignKeys string
	if len(table.ForeignKeys) == 0 {
		foreignKeys = "*None*"
	} else {
		fks := make([]string, len(table.ForeignKeys))
		for i, fk := range table.ForeignKeys {
			fks[i] = fmt.Sprintf("- %v (%v) `REFERENCES` [%v](../../../%v).[%v](../../../%v/tables/%v)",
				fk.Name,
				fmt.Sprintf("**%v**", strings.Join(fk.Columns, "**, **")),
				fk.RefTable.Schema, fk.RefTable.Schema, fk.RefTable.Name,
				fk.RefTable.Schema, fk.RefTable.Name,
			)
		}
		foreignKeys = strings.Join(fks, "\n")
	}

	var referencedBy string
	if len(table.ReferencedBy) == 0 {
		referencedBy = "*None*"
	} else {
		refBys := make([]string, len(table.ReferencedBy))
		for i, refBy := range table.ReferencedBy {
			refBys[i] = fmt.Sprintf("- [%v](../../../%v).[%v](../../../%v/tables/%v)", refBy.Schema, refBy.Schema, refBy.Name, refBy.Schema, refBy.Name)
		}
		referencedBy = strings.Join(refBys, "\n")
	}

	columns := make([]string, len(table.Columns))
	for i, c := range table.Columns {
		columns[i] = fmt.Sprintf("- `%v` %v", c.Name, c.DbType)
	}
	_, err := fmt.Fprintf(w, `# Table: [%v](..).%v

## Primary key
%v

## Foreign keys
%v

## Refenced by
%v

## Columns
%v

> Generated by free [DataTug.app](https://datatug.app)
`, table.Schema, table.Name, primaryKey, foreignKeys, referencedBy, strings.Join(columns, "\n"))
	return err
}