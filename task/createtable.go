package task
import (
	"database/sql"
	_ "github.com/lib/pq"
    "os"
    "log"
    "fmt"
)

func CreateTable() {
	url := os.Getenv("DATABASE_URL")
	db,err := sql.Open("postgres",url)
	if err != nil {
		log.Fatal(err)
		return
	}

    defer db.Close()
    
	createTb := ` 
	CREATE TABLE IF NOT EXISTS customer ( 
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
		);
		`
		_, err = db.Exec(createTb) 
		if err != nil {
			log.Fatal("can't create table", err)
		}
    fmt.Println("create table success")
}