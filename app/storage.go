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
	GetAccounts() error
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

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account := new(Account)

		account, err := scanFromRow(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	query := "DELETE FROM accounts WHERE id = $1"

	_, err := s.db.Query(query, id)

	return err
}

func UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	query := "SELECT * FROM accounts WHERE id = $1"
	row, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	account := new(Account)
	if row.Next() {
		account, err = scanFromRow(row)
	} else {
		return nil, fmt.Errorf("record not found with id %d", id)
	}

	if err != nil {
		return nil, err
	}

	return account, nil
}

func scanFromRow(rows *sql.Rows) (*Account, error) {
	account := Account{}
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)

	return &account, err
}
