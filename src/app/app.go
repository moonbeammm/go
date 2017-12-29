package app

import (
	"io"
	"net/http"
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// path: /app
func AppHelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s",req.URL)

	// 打开mysql
	db, err := sql.Open("mysql", "root:123456@tcp(47.52.250.43:3306)/sgx?parseTime=true")
    if err != nil{
        log.Fatal(err)
	}
	// 关闭mysql
	defer db.Close()

	// 打开sgx.T_Account这个表
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sgx.T_Account(world varchar(50))")
    if err != nil{
        log.Fatalln(err)
	}
	
	// 查询表
	rows, err := db.Query("SELECT * FROM T_Account")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}
	// 关闭表
	defer rows.Close()

	for rows.Next() {
		var id int
		var account, pwd, accesskey string
		rows.Scan(&id, &account, &pwd, &accesskey)
		fmt.Println("id:", id, "account:", account, "password:", pwd)
	}
	io.WriteString(w, "hello, world.")
}