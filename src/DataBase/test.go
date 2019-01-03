package main
import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)
func main() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)
	// 插入数据
	stmt, err := db.Prepare("INSERT user_info SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("test", " 研发部门", "2017-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update user_info set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("test", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM user_info")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from user_info where uid=?")
	checkErr(err)
	//res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

