package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// create table
	cmd := `CREATE TABLE IF NOT EXISTS person(
						name STRING,
						age INT)
					`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// insert data
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// // _, err = DbConnection.Exec(cmd, "Nancy", 20)
	// _, err = DbConnection.Exec(cmd, "Bob", 22)
	// // _, err = DbConnection.Exec(cmd, "Ken", 25)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 100, "Ken")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// マルチプルセレクト
	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// fmt.Println(rows.Next())
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// シングルセレクト
	// cmd = "SELECT * FROM person WHERE age = ?"
	// row := DbConnection.QueryRow(cmd, 210)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Panicln(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "Bob")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

}
