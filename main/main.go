package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//1.创建db连接
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/zyd")
	//判空
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	//关闭资源
	defer func() {
		db.Close()
	}()
	//真正的开启
	db.Ping()

	//2.预处理（查询）
	stm, err := db.Prepare(" select * from student")
	//判空
	if err != nil {
		fmt.Println("预处理失败")
		return
	}
	//关闭资源
	defer func() {
		stm.Close()
	}()

	//3.获取查询结果
	rows, err := stm.Query()
	//判空
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	fmt.Println("aaaaaa",rows)
	//关闭资源
	// defer func() {
	// 	if rows != nil {
	// 		rows.Close()
	// 	}
	// }()
	//循环遍历结果
	for rows.Next() {
		var student_id int
		var  student_title, student_author,  studentdate string
		//分别赋值
		rows.Scan(&student_id,&student_title,&student_author,&studentdate)
		fmt.Println(student_title,"123")
	}
}