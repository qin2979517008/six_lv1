package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"log"
)
func main(){
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/one?charset=utf8")
	defer db.Close()
	 if(err != nil){
	 	fmt.Println("打开数据库失败")
	 	fmt.Println(err)
	 }
	insert(db)
	selectDB(db)
	update(db)
	remove(db)
}
//增加小明，30岁
func insert(db *sql.DB)  {
	stmt,err := db.Prepare("INSERT first (name,age) values (?,?)")
	if (err !=nil){
		log.Fatal(err)
	}
	  stmt.Exec("xiaoming",30)
	fmt.Println("插入成功")
}
//查询全表
func selectDB(db *sql.DB)  {
	stmt, err := db.Query("select * from first;")
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for stmt.Next(){
		var age int
		var name string
		err := stmt.Scan(&name,&age)
		if err != nil {
			 fmt.Println(err)
		}
		fmt.Println(name,age)
	}
}
//改，将小明的名字改为老王
func update(db *sql.DB)  {
	stmt1,err := db.Prepare("UPDATE first SET name=? WHERE age=?")
	stmt1.Exec("laowang",30)
	if (err !=nil){
		log.Fatal(err)
	}
}
//删除老王
func remove(db *sql.DB)  {
	stmt, err := db.Prepare(`DELETE FROM first WHERE name=?`)
	if (err !=nil){
		log.Fatal(err)
	}
	//执行删除操作
	  stmt.Exec("laowang")
	fmt.Println("删除成功")
}
