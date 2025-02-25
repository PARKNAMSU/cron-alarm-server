package query_tool

type QueryAction = int

type JoinType = string

type ColumnCompareType = int

type CompareColumn struct {
	CompareType ColumnCompareType
	Value       interface{}
}

const (
	EQUAL         ColumnCompareType = 0
	NOT_EQUAL     ColumnCompareType = 1
	IN            ColumnCompareType = 2
	NOT_IN        ColumnCompareType = 3
	GREATER       ColumnCompareType = 4
	GREATER_EQUAL ColumnCompareType = 5
	SMALLER       ColumnCompareType = 6
	SMALLER_EQUAL ColumnCompareType = 7
	BITWISE       ColumnCompareType = 8
)

const (
	INSERT    QueryAction = 0
	UPDATE    QueryAction = 1
	SELECT    QueryAction = 2
	DELETE    QueryAction = 3
	DUPLICATE QueryAction = 4
	IGNORE    QueryAction = 5
)

const (
	INNER JoinType = "INNER"
	LEFT  JoinType = "LEFT"
)

type QueryParams struct {
	Table      string
	Action     QueryAction
	Select     []string
	Where      map[string]interface{}
	Set        map[string]interface{}
	Conditions string
	Duplicate  map[string]interface{}
	As         string
	Join       []JoinParams
}

type JoinParams struct {
	Table string
	On    string
	Type  JoinType
	As    string
}
