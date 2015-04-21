package main

import (
	"./utility"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// sql := utility.OpenDB()
	// sql := utility.OpenDB("mysql", "root:123@/benchmark")
	// sql := utility.OpenDB("mssql", "Server=10.10.5.36;Database=seego;User Id=sa;Password=123")
	sql := utility.OpenDB("postgres", "user=postgres dbname=benchmark password=1234")
	defer sql.CloseDB()

	time1 := time.Now()

	sql.Execute("delete from benchmark1")

	length := 500
	segment := 100

	row := 0
	for row < length {
		sql.Execute("insert into benchmark1(id, name) values (?,?)", (row + 1), "data "+strconv.Itoa(row))
		row = row + 1

		time2 := time.Now()

		if row%segment == 0 {
			delay := time2.Sub(time1)
			fmt.Println(row/segment, delay)

			time1 = time2
		}
	}
}
