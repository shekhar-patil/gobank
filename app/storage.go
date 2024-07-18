package app

import (
	"database/sql"
	"fmt"
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
    	created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO accounts
	(first_name, last_name, number, balance, created_at)
	values ($1, $2, $3, $4, $5)`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
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
