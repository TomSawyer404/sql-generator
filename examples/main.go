package main

import (
	"fmt"
	"sql-generator/internal"
	"sql-generator/internal/entity"
)

func main() {
	db := entity.DbData{}
	db.Table = "Student"
	db.Fields = "ID,Name,AGE"

	where := make([]entity.WhereItem, 0)
	where = append(where, entity.WhereItem{
		Field: "ID",
		Value: "'0906'",
		SqlOp: entity.EQ,
		OrAnd: entity.AND,
	})
	db.Where = where

	orderby := make([]entity.OrderByItem, 0)
	orderby = append(orderby, entity.OrderByItem{
		Field: "ID",
		Order: entity.ASC,
	})
	db.OrderBy = orderby

	db.DbType = entity.MSSQL
	sqlBuilder := internal.SqlBuilder{}

	sql0 := sqlBuilder.GenSelectSQL(&db)
	fmt.Println(sql0)

	db.PageIndex = 1
	db.PageSize = 20

	sql0 = sqlBuilder.GenPageSQL(&db)
	fmt.Println(sql0)

	db.DbType = entity.MYSQL
	sql0 = sqlBuilder.GenSelectSQL(&db)
	fmt.Println(sql0)
	sql0 = sqlBuilder.GenPageSQL(&db)
	fmt.Println(sql0)
}
