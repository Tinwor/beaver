package server

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	client "github.com/Tinwor/beaver/grpc/data/grpcuser"
	_ "github.com/lib/pq"
	"github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	db_name     = "beaver"
	db_user     = "postgres"
	db_password = "password"
	db_port     = 32768
	db_host     = "localhost"
)

type UserServer struct {
	loginQuery        *sql.Stmt
	newUserQuery      *sql.Stmt
	checkRegistration *sql.Stmt
	db                *sql.DB
}

func NewUserServer() *UserServer {
	rand.Seed(time.Now().UnixNano())
	user := UserServer{db: newDBConnection()}
	stmt, err := user.db.Prepare("SELECT guid FROM users WHERE username = $1")
	if err != nil {
		defer user.db.Close()
		log.Fatal("Error preparing query " + err.Error())
	}
	user.loginQuery = stmt
	stmt, err = user.db.Prepare("INSERT INTO users(guid, username, email, password, salt, created) VALUES($1,$2,$3,$4,$5, $6)")
	if err != nil {
		defer user.db.Close()
		log.Fatal("Error preparing query " + err.Error())
	}
	user.newUserQuery = stmt
	stmt, err = user.db.Prepare("SELECT * FROM users WHERE username = $1 OR email = $2")
	if err != nil {
		defer user.db.Close()
		log.Fatal("Error preparing query: " + err.Error())
	}
	user.checkRegistration = stmt
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
func (u *UserServer) NewUser(context context.Context, in *client.RegisterUser) (*client.Response, error) {
	guid := uuid.NewV4().String()
	rows, err := u.checkRegistration.Query(in.Username, in.Email)
	if err != nil {
		log.Println("Error executing query: ", err.Error())
		return &client.Response{
			Status: client.StatusResponse_SERVER_ERROR,
		}, err
	}
	if rows == nil {

		salt := randStringRunes(6)
		password, err := bcrypt.GenerateFromPassword([]byte(salt+in.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing password: ", err.Error())
			return &client.Response{
				Status: client.StatusResponse_SERVER_ERROR,
			}, err
		}
		_, err = u.newUserQuery.Exec(guid, in.Username, in.Email, password, salt, time.Now())
		if err != nil {
			log.Println("Error inserting new user: " + err.Error())
			return &client.Response{
				Status: client.StatusResponse_SERVER_ERROR,
			}, err
		}

	} else {
		return &client.Response{
			Status: client.StatusResponse_CREDENTIAL_EXIST,
		}, nil
	}
	return &client.Response{
		Status: client.StatusResponse_OK,
		Token:  guid,
	}, nil

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
