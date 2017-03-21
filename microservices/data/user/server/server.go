package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	client "github.com/Tinwor/beaver/grpc/data/grpcuser"
)

const (
	db_name     = "beaver"
	db_user     = "postgre"
	db_password = "password"
	db_port     = 8080
	db_host     = "localhost"
)

type UserServer struct {
	loginQuery   *sql.Stmt
	newUserQuery *sql.Stmt
	db           *sql.DB
}

func NewUserServer() *UserServer {

	user := UserServer{db: newDBConnection()}
	stmt, err := user.db.Prepare("SELECT guid FROM users WHERE username = $1")
	if err != nil {
		defer user.db.Close()
		log.Fatal("Error preparing query")
	}
	user.loginQuery = stmt
	stmt, err = user.db.Prepare("INSERT INTO users(guid, username, email, password, salt) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		user.db.Close()
		log.Fatal("Error preparing query")
	}
	user.newUserQuery = stmt

	return &user
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
func (u *UserServer) UserLogin(context context.Context, in *client.LoginRequest) (*client.Response, error) {
	rows, err := u.loginQuery.Query(in.Username, in.Password)
	if err != nil {
		log.Println("Error executing query: ", err.Error())
		return &client.Response{
			Status: client.StatusResponse_SERVER_ERROR,
		}, err
	}
	if rows == nil {
		return &client.Response{
			Status: client.StatusResponse_FAILED,
		}, nil
	}
	var guid string
	for rows.Next() {
		rows.Scan(&guid)
	}
	return &client.Response{
		Status: client.StatusResponse_OK,
		Token:  guid,
	}, nil
}
func (u *UserServer) NewUser(context.Context, *client.RegisterUser) (*client.Response, error) {
	return nil, nil
}
