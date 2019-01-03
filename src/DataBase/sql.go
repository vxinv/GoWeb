package main
import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"net/http"
)
type DbWorker struct {
	//mysql data source name
	Dsn string
}

func sayhello(w http.ResponseWriter,r *http.Request)  {

}
func main() {
	var DB *sql.DB
    DB = InitDb()
    Insert(DB)
	Query(DB)
	defer DB.Close()

}
func Checkerr(e error,s string){
	if(e!=nil){
		fmt.Println(s)
		panic(e)
	}
}
func Insert(db *sql.DB)  {
	//开启事务
	stmt,err:=db.Prepare("INSERT user SET username=?,detailInfo=?,exedata=?,wendadata=?,historyMsg=?")
	Checkerr(err,"准备错误")
	res,err:=stmt.Exec("李鑫","15电气一班",0,0,0)
	Checkerr(err,"执行错误")
	fmt.Println(res.LastInsertId())
}
func InitDb() *sql.DB {
	dbw := DbWorker{
		Dsn: "root:root@tcp(211.159.175.75:3306)/lixin",
	}
	DB, err := sql.Open("mysql", dbw.Dsn)
	if err != nil {
		fmt.Println("初始化错误")
		panic(err)
	}
	//fmt.Println(DB)
	return DB;
}
func Query(db *sql.DB) {
	rows, err := db.Query("SELECT username * FROM user")
	Checkerr(err,"查询错误")
	for rows.Next() {
		var uid int
		var username string
		var detailInfo string
		var exedata int
		var wendadata int
		var historyMsg int
		err = rows.Scan(&uid, &username, &detailInfo, &exedata,&wendadata,&historyMsg)
		Checkerr(err,"读取错误")
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(historyMsg)
		fmt.Println(exedata)
	}
}
