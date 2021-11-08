package api

import "sql-generator/internal/entity"

type IBuilder interface {
	GenSelect(*entity.DbData) string
	GenPage(*entity.DbData) string
}
