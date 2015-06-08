package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	create table cpeinfo (ppg_id text, cpe_id text, model text, olt_id text, frame integer, slot integer,  port integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into cpeinfo(ppg_id, cpe_id, model, olt_id, frame, slot,  port) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 16; i++ {
		_, err = stmt.Exec(
			fmt.Sprintf("ppg_id%03d", i),
			fmt.Sprintf("cpe_id%03d", i),
			"245",
			"lab:gpon1",
			0,
			2,
			i%5 == 0)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select * from cpeinfo where port = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println("EFA", rows)
		// var id int
		// var name string
		// rows.Scan(&id, &name)
		// fmt.Println(id, name)
	}
	rows.Close()

	// stmt, err = db.Prepare("select name from foo where id = ?")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer stmt.Close()
	// var name string
	// err = stmt.QueryRow("3").Scan(&name)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(name)

	// _, err = db.Exec("delete from foo")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// rows, err = db.Query("select id, name from foo")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	rows.Scan(&id, &name)
	// 	fmt.Println(id, name)
	// }
}
