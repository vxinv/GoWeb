package main
import (
    _"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)
	type DbWorker struct {
		//mysql data source name
		Dsn string
	}
	func main() {
		dbw := DbWorker{
			Dsn: "root:root@tcp(211.159.175.75:3306)/lixin",
		}
		db, err := sql.Open("mysql",dbw.Dsn)
		if err != nil {
			panic(err)
		}
		fmt.Println("hello world")
		defer db.Close()
	}
