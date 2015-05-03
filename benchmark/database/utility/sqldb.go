package utility

import (
	"database/sql"
	"fmt"
)

type SqlDB struct {
	DB *sql.DB
}

func OpenDB(args ...string) *SqlDB {
	if len(args) >= 2 {
		db, _ := sql.Open(args[0], args[1])
		return &SqlDB{DB: db}
	} else {
		db, _ := sql.Open("mysql", "root:123@/boogr")
		return &SqlDB{DB: db}
	}
}

func (s *SqlDB) Execute(query string, args ...interface{}) {
	s.DB.Exec(query, args...)
}

func (s *SqlDB) ExecuteRows(query string, args ...interface{}) *sql.Rows {
	rows, _ := s.DB.Query(query, args...)
	return rows
}

func (s *SqlDB) QueryToList(query string, args ...interface{}) []interface{} {
	rows := s.ExecuteRows(query, args...)
	cols, _ := rows.Columns()

	val1 := make([]sql.RawBytes, len(cols))
	val2 := make([]interface{}, len(cols))
	list := make([]interface{}, 0)

	for rows.Next() {
		value := make(map[string]interface{})
		for i, _ := range cols {
			val2[i] = &val1[i]
		}
		rows.Scan(val2...)

		for i, _ := range cols {
			if val1[i] == nil {
				value[cols[i]] = nil
			} else {
				value[cols[i]] = string(val1[i])
			}
		}

		list = append(list, value)
	}
	return list
}

func (s *SqlDB) CloseDB() {
	s.DB.Close()
}

func sqldb_test() {
	fmt.Println("...")
}
