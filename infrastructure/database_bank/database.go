package database_bank

import (
	"Bankirka/config"
	"Bankirka/internal/entity"
	"Bankirka/internal/service"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

type DB struct {
	Conn *pgx.Conn
}

type User struct {
	ID      int
	Balance int
}

func NewDB() *DB {

	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}
	conn, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	fmt.Println("Успешное подключение к базе данных!")
	return &DB{Conn: conn}
}

func (db *DB) CreatePerson(id int, balance entity.Balance) error {

	var userID int64
	var exists bool

	err := db.Conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM bank.users WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		return err
	}
	if exists == true {
		return service.AccountExistErr
	}

	err = db.Conn.QueryRow(context.Background(),
		"INSERT INTO bank.users (id, balance) VALUES ($1, $2) RETURNING id", id, balance.Money).Scan(&userID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ChangeBalance(id int, dif entity.Difference) error {
	bal, err := db.ShowBalance(id)
	if err != nil {
		return err
	}

	bal = bal + dif.Quantity

	_, err = db.Conn.Exec(context.Background(), "UPDATE bank.users SET balance=$1 WHERE id=$2", bal, id)
	if err != nil {
		return err
	}
	return nil

}

func (db *DB) ShowBalance(id int) (int, error) {

	var user User
	var exists bool

	err := db.Conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM bank.users WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if exists == false {
		return 0, service.NoAccountErr
	}

	err = db.Conn.QueryRow(context.Background(), "SELECT id, balance FROM bank.users WHERE id=$1", id).Scan(&user.ID, &user.Balance)
	if err != nil {
		return 0, err
	}
	return user.Balance, nil
}
