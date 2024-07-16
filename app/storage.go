package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "host=localhost user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		db: db,
	}, err
}

func (s *PostgresStore) Init() error {
	err := s.CreateTable()
	return err
}

func (s *PostgresStore) CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
    	number serial,
    	balance integer,
    	create_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

func CreateAccount(account *Account) error {
	return nil
}

func DeleteAccount(id int) error {
	return nil
}

func UpdateAccount(account *Account) error {
	return nil
}

func GetAccountById(id int) error {
	return nil
}
