package entity

type SqlOp string

const (
	GE SqlOp = ">="
	GR SqlOp = ">"
	EQ SqlOp = "="
	LQ SqlOp = "<="
	LE SqlOp = "<"
)

type OrAnd string

const (
	OR   OrAnd = " OR "
	AND  OrAnd = " AND "
	NONE OrAnd = ""
)

type WhereItem struct {
	Field string
	Value string
	SqlOp
	OrAnd
}
