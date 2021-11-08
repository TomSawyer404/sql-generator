//! This is a simple implementation of interface IBuilder

package mssql

import (
	"sql-generator/internal/entity"
	"strconv"
	"strings"
)

type MsSql struct{}

/// Implementation of the IBuilder interface
/// Example:
///   SELECT age FROM my_table WHERE id==1;
func (this *MsSql) GenSelect(da *entity.DbData) string {
	if da == nil {
		return ``
	}

	s := make([]string, 0)
	s = append(s, "SELECT ")
	fields := strings.Split(da.Fields, ",")
	for _, field := range fields {
		s = append(s, entity.DbMapLeft[da.DbType])
		s = append(s, field)
		s = append(s, entity.DbMapRight[da.DbType])
		s = append(s, ",")
	}

	s = s[0 : len(s)-1]
	s = append(s, " FROM ")
	s = append(s, entity.DbMapLeft[da.DbType])
	s = append(s, da.Table)
	s = append(s, entity.DbMapRight[da.DbType])

	s = append(s, " WHERE ")
	if da.Where == nil {
		s = append(s, "1=1")
	} else {
		s = append(s, da.GenWhere())
	}

	if da.OrderBy != nil {
		s = append(s, " ORDER BY ")
		s = append(s, da.GenOrderBy())
	}

	return strings.Join(s, "")
}

func (this *MsSql) GenPage(da *entity.DbData) string {
	if da == nil {
		return ""
	}

	s := make([]string, 0)
	str := this.GenSelect(da)

	// limit (curPage-1)*pageSize,pageSize
	s = append(s, str)
	s = append(s, " offset ")

	i := (da.PageIndex - 1) * da.PageSize

	s = append(s, strconv.Itoa(i))
	s = append(s, " rows fetch next ")
	s = append(s, strconv.Itoa(da.PageSize))
	s = append(s, " rows only ")

	return strings.Join(s, "")
}
