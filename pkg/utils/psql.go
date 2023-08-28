package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func StringSliceToDollarPsqlArray(len_ int, offset int) string {

	var values []string

	for i := 1; i < len_+1; i++ {
		values = append(values, fmt.Sprintf("$%d", i+offset))
	}

	return " (" + strings.Join(values, ", ") + " ) "
}

func CloseConnection(connection *sql.Rows) error {
	if connection == nil {
		return errors.New("connection is nil")
	}
	return connection.Close()
}

func StringSliceToDollarPsqlArrayWithTTL(len_ int, Ttl []interface{}) string {

	var values []string
	var count int
	var num int

	for i := 1; i < len_+1; i++ {
		num = len_ + i + 1 - count
		//if Ttl[i-1] != "" {
		//	values = append(values,
		//		fmt.Sprintf("($1, (SELECT segment_id FROM public.segment WHERE segment_name = $%d), $%d)", i+1, num))
		//
		//} else {
		//	values = append(values,
		//		fmt.Sprintf("($1, (SELECT segment_id FROM public.segment WHERE segment_name = $%d), NULL)", i+1))
		//	count += 1
		//}
		values = append(values,
			fmt.Sprintf("\n ($1, (SELECT segment_id FROM public.segment WHERE segment_name = $%d), $%d)", i+1, num))
	}
	return strings.Join(values, ", \n")
}

func InserrtRowsDollarParams(len_ int) string {

	var values []string

	for i := 1; i < len_+1; i++ {
		values = append(values,
			fmt.Sprintf("($1, $%d, $2, $3)", i+3))
	}

	return strings.Join(values, ", ")
}

func AddNulls(ttl []interface{}) []interface{} {
	for i, val := range ttl {
		if val == "" {
			ttl[i] = sql.NullTime{Valid: false}
		}
	}
	return ttl
}
