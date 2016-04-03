package table

import (
	"../index"
)

var TableList []*Table

func CreateTable(tableName string) *Table {
	return &Table{
		Name:      tableName,
		FirstPage: NewPage(),
	}
}

func (t *Table) AddFiled(fieldName string, fixedSize bool, size int, ttype int, keyType int) {
	ind := index.CreateIndex(keyType, "db", t.Name, fieldName+"_key")
	filed := Field{
		Name:      fieldName,
		FixedSize: fixedSize,
		Size:      size,
		Type:      ttype,
		Index:     ind,
	}
	t.Fields = append(t.Fields, filed)
}
