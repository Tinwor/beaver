package server

import _ "github.com/lib/pq"
import "fmt"
import "database/sql"
import "log"

const (
	db_name     = "beaver"
	db_user     = "postgre"
	db_password = "password"
	db_port     = 8080
	db_host     = "localhost"
)

type UserServer struct {
}

func NewUserServer() {

}
func newDBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_password, db_name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		defer db.Close()
		log.Fatal(err.Error())
		return nil
	}
	return db
}
