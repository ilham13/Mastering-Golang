package basic

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID    string
	Name  string
	Age   int
	Grade int
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SqlQuery() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
	}

}

func SqlQueryRow() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = Student{}
	var id = "E001"
	err = db.
		QueryRow("select id, name, grade from tb_student where id = ?", id).
		Scan(&result.ID, &result.Name, &result.Grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result.Name)
}

func SqlExec() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tb_student values (?, ?, ?, ?)", "G021", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("UPDATE tb_student SET age = ? WHERE grade = ?", 59, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("DELETE FROM tb_student WHERE grade = ?", 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
