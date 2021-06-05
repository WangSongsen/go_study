/*
# @Time : 2021/6/5 17:10
# @Author : team_go
# @File : homework.go
# @Software: GoLand
*/

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"log"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "./data.db3"
)

type user struct {
	Username string
	Age      int
	Job      string
	Hobby    string
}

func checkErr(e error) bool {
	if e != nil {
		log.Println("err : ", e.Error())
		return true
	}
	return false
}

func createTable(db *sql.DB) error {
	sql := `create table if not exists "users" (
		"id" integer primary key autoincrement,
		"username" text not null,
		"age" integer not null,
		"job" text,
		"hobby" text
	)`
	_, err := db.Exec(sql)
	return err
}

func selectTable(db *sql.DB, id int) (string, error) {
	var name string
	var amp string
	err := db.QueryRow("select username from users where id = ?", id).Scan(&amp, name)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有查询到数据，这里分两种场景一种是数据本来就不存在(正常场景) 返回 nil 即可
			// 如果是用户的错误输入导致，这里应该是错误场景
			// 例如 业务已经规定 id 取值范围为 0-10 这时候就属于输入错误导致异常结果，应该错误返回
			if id < 0 || id > 10 {
				return amp, errors.Wrap(err, "error 查询失败，id超出限制")
			}
			return amp, nil
		} else {
			return amp, errors.Wrap(err, "error 查询失败")
		}
	}
	return amp, nil
}

func main() {
	db, err := sql.Open(dbDriverName, dbName)
	if checkErr(err) {
		return
	}
	defer db.Close()
	err = createTable(db)
	if checkErr(err) {
		return
	}
	name, err := selectTable(db, -1)
	if checkErr(err) {
		return
	}
	log.Println("name is :", name)

}
