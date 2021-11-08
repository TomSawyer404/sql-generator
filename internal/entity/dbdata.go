package entity

import (
	"strings"
)

type DbData struct {
	Table     string
	Fields    string
	Where     []WhereItem
	OrderBy   []OrderByItem
	PageIndex int
	PageSize  int
	DbType
}

func (this *DbData) GenWhere() string {
	s := make([]string, 0)
	for _, v := range this.Where {
		s = append(s, DbMapLeft[this.DbType])
		s = append(s, v.Field)
		s = append(s, DbMapRight[this.DbType])

		s = append(s, string(v.SqlOp))
		s = append(s, v.Value)
		//s = append(s, string(v.OrAnd))
		s = append(s, " ")
	}

	return strings.Join(s, "")
}

func (this *DbData) GenOrderBy() string {
	s := make([]string, 0)
	for _, v := range this.OrderBy {
		s = append(s, DbMapLeft[this.DbType])
		s = append(s, v.Field)
		s = append(s, DbMapRight[this.DbType])

		s = append(s, string(v.Order))
		s = append(s, ",")
	}

	s = s[0 : len(s)-1]
	return strings.Join(s, "")
}
