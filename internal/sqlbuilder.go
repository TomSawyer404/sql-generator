package internal

import (
	"sql-generator/internal/entity"
	"sql-generator/internal/mssql"
	"sql-generator/internal/mysql"
	"sql-generator/pkg/api"
)

type SqlBuilder struct{}
type _sqlBuilder struct {
	api.IBuilder
}

func (this *SqlBuilder) GenSelectSQL(da *entity.DbData) string {
	if da == nil {
		return ""
	}

	switch da.DbType {
	case entity.MYSQL:
		{
			mysql := mysql.MySql{}
			sqlBuilder := _sqlBuilder{&mysql}
			sql := sqlBuilder.GenSelect(da)
			return sql
		}
	case entity.MSSQL:
		{
			mssql := mssql.MsSql{}
			sqlBuilder := _sqlBuilder{&mssql}
			sql := sqlBuilder.GenSelect(da)
			return sql
		}
	}

	return ""
}

func (this *SqlBuilder) GenPageSQL(da *entity.DbData) string {
	if da == nil {
		return ""
	}

	switch da.DbType {
	case entity.MYSQL:
		{
			mysql := mysql.MySql{}
			sqlBuilder := _sqlBuilder{&mysql}
			sql := sqlBuilder.GenPage(da)
			return sql
		}
	case entity.MSSQL:
		{
			mssql := mssql.MsSql{}
			sqlBuilder := _sqlBuilder{&mssql}
			sql := sqlBuilder.GenPage(da)
			return sql
		}
	}
	return ""
}
