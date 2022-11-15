package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// For SQL table "times"
type Times struct {
	Id       int
	Datetime time.Time
}

var datetime = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

var INSERT = fmt.Sprintf("INSERT INTO times (id, datetime) VALUES(0, '%s')",
	datetime.Format(time.RFC3339))

const SELECT = "SELECT * FROM times WHERE id = 0"

func main() {

	db, err := gorm.Open("postgres", "user=postgres password=mypassword dbname=demo sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	database := db.DB()
	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	if _, err = database.Exec("CREATE TABLE IF NOT EXISTS times (id integer, datetime timestamp without time zone)"); err != nil {
		panic(err)
	}

	if _, err = database.Exec(INSERT); err != nil {
		panic(err)
	}

	if _, err = database.Exec(INSERT); err != nil {
		panic(err)
	}

	rows := database.QueryRow(SELECT)
	fmt.Println(rows)
	t := Times{}
	rows.Scan(&t)
	fmt.Printf("Id: %d, Datetime: %s\n", t.Id, t.Datetime)

}
