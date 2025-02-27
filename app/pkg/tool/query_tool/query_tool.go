package query_tool

import (
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"nspark-cron-alarm.com/cron-alarm-server/app/pkg/tool/common_tool"
)

func generateWhereQueryStr(key string, columnValue any, compareType ColumnCompareType) (string, any) {
	var query string
	value := columnValue
	columnKey := key

	if !strings.Contains(columnKey, ".") {
		columnKey = fmt.Sprintf("`%s`", columnKey)
	}

	switch compareType {
	case EQUAL:
		query = fmt.Sprintf("%s = ?", columnKey)
	case NOT_EQUAL:
		query = fmt.Sprintf("%s != ?", columnKey)
	case GREATER:
		query = fmt.Sprintf("%s > ?", columnKey)
	case GREATER_EQUAL:
		query = fmt.Sprintf("%s >= ?", columnKey)
	case SMALLER:
		query = fmt.Sprintf("%s < ?", columnKey)
	case SMALLER_EQUAL:
		query = fmt.Sprintf("%s <= ?", columnKey)
	case BITWISE:
		query = fmt.Sprintf("%s & ?", columnKey)
	case IN:
		{
			inQuery, args, err := sqlx.In(
				fmt.Sprintf("%s IN(?)", columnKey),
				columnValue,
			)
			if err != nil {
				log.Panicln(err)
			}
			query = inQuery
			value = args
		}
	case NOT_IN:
		{
			inQuery, args, err := sqlx.In(
				fmt.Sprintf("%s NOT IN(?)", columnKey),
				columnValue,
			)
			if err != nil {
				log.Panicln(err)
			}
			query = inQuery
			value = args
		}
	}

	return query, value
}

func QueryBuilder(params QueryParams) (string, []any) {
	var query string
	var where string
	var set string
	var duplicate string
	var join string

	whereParams := make([]any, 0, len(params.Where))
	setParams := make([]any, 0, len(params.Set))
	dupParams := make([]any, 0, len(params.Duplicate))

	generateJoinQuery := func() error {
		joinStr := make([]string, 0, len(params.Join))
		for _, join := range params.Join {
			joinStr = append(joinStr, fmt.Sprintf("%s JOIN %s AS %s ON %s ", join.Type, join.Table, join.As, join.On))
		}
		join = strings.Join(joinStr, " ")
		return nil
	}

	generateSetQuery := func() error {
		if params.Action == SELECT || params.Action == DELETE {
			return nil
		}
		if len(params.Set) <= 0 {
			return nil
		}

		set = "SET"

		setList := make([]string, 0, len(params.Set))

		for k, v := range params.Set {
			setList = append(setList, fmt.Sprintf("`%s` = ?", k))
			setParams = append(setParams, v)
		}

		set = strings.Join([]string{
			set,
			strings.Join(setList, " , "),
		}, " ")
		return nil
	}

	generateWhereQuery := func() error {
		if params.Action == INSERT || params.Action == IGNORE {
			return nil
		}

		if len(params.Where) <= 0 {
			return nil
		}

		where = "WHERE"

		whereList := make([]string, 0, len(params.Where))

		for k, v := range params.Where {
			// statement 키가 입력된 경우 해당 로직 진행
			if k == "statement" {
				state, isString := v.(string)
				if isString {
					whereList = append(whereList, state)
					continue
				}
				stateList, isList := v.([]string)
				if isList {
					whereList = append(whereList, stateList...)
					continue
				}
				log.Panicln("invalidation statement query")
			}

			compareType := EQUAL

			var value any = v

			column, isCustomCompare := v.(CompareColumn)

			if isCustomCompare {
				compareType = column.CompareType
				value = column.Value
			} else if common_tool.IsSlice(v) {
				compareType = IN
			}

			queryStr, queryValue := generateWhereQueryStr(k, value, compareType)

			whereList = append(whereList, queryStr)
			if common_tool.IsSlice(queryValue) {
				whereParams = append(whereParams, queryValue.([]any)...)
			} else {
				whereParams = append(whereParams, queryValue)
			}
		}

		where = strings.Join([]string{
			where,
			strings.Join(whereList, " AND "),
		}, " ")
		return nil
	}

	generateDupQuery := func() error {
		if params.Action != DUPLICATE {
			return nil
		}

		if len(params.Duplicate) <= 0 {
			return nil
		}

		duplicate = "ON DUPLICATE KEY UPDATE"
		dupList := make([]string, 0, len(params.Duplicate))

		for k, v := range params.Duplicate {
			dupList = append(dupList, fmt.Sprintf("`%s` = ?", k))
			dupParams = append(dupParams, v)
		}

		duplicate = strings.Join([]string{
			duplicate,
			strings.Join(dupList, " , "),
		}, " ")
		return nil
	}

	switch params.Action {
	case UPDATE:
		query += fmt.Sprintf("UPDATE %s ", params.Table)
	case DELETE:
		query += fmt.Sprintf("DELETE FROM %s ", params.Table)
	case IGNORE:
		query += fmt.Sprintf("INSERT IGNORE INTO %s ", params.Table)
	case SELECT:
		query += "SELECT"
		var selectList string
		if len(params.Select) > 0 {
			selectList = strings.Join(params.Select, " , ")
		} else {
			selectList = "*"
		}
		query = strings.Join([]string{query, selectList, fmt.Sprintf("FROM %s ", params.Table), params.As}, " ")
	case INSERT:
		fallthrough
	case DUPLICATE:
		query += fmt.Sprintf("INSERT INTO %s ", params.Table)
	default:
		log.Panicln("not support query action")
	}

	common_tool.ParallelExec(generateJoinQuery, generateSetQuery, generateWhereQuery, generateDupQuery)

	query = strings.Join([]string{
		query,
		join,
		set,
		where,
		duplicate,
		params.Conditions,
	}, " ")

	var returnParams = make([]any, 0, len(setParams)+len(whereParams)+len(dupParams))

	returnParams = append(returnParams, setParams...)
	returnParams = append(returnParams, whereParams...)
	returnParams = append(returnParams, dupParams...)

	return query, returnParams
}

func QueryExecute(params QueryParams, db *sqlx.DB) {
	query, values := QueryBuilder(params)
	_, err := db.Exec(query, values...)

	if err != nil {
		log.Println(err)
	}
}

func QuerySelect[T any](params QueryParams, db *sqlx.DB) T {
	var returnValue T

	query, values := QueryBuilder(params)
	err := db.Select(&returnValue, query, values...)

	if err != nil {
		log.Panicln(err)
	}
	return returnValue
}

func QueryPrint(params QueryParams) string {
	query, values := QueryBuilder(params)

	for _, data := range values {
		data := data
		strData, isStr := data.(string)
		if isStr {
			data = fmt.Sprintf("'%s'", strData)
		}
		query = strings.Replace(query, "?", fmt.Sprintf("%v", data), 1)
	}
	return query

}
