package postgres

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

	reqScl := "INSERT INTO bank.users (id, balance) VALUES ($1, $2)"

	_, err := db.Conn.Exec(context.Background(),
		reqScl, id, balance.Money)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_id_key\" (SQLSTATE 23505)" {
			return service.AccountExistErr
		}
		return err
	}
	return nil
}

func (db *DB) ChangeBalance(id int, dif entity.Difference) error {
	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		log.Fatal("Unable to start transaction:", err)
	}

	_, err = db.ShowBalanceTx(tx, id)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	reqSql := "UPDATE bank.users SET balance=balance+$1 WHERE id=$2"
	_, err = tx.Exec(context.Background(), reqSql, dif.Quantity, id)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	tx.Commit(context.Background())

	return nil

}

func (db *DB) ShowBalance(id int) (int, error) {

	var user User
	reqSql := "SELECT id, balance FROM bank.users WHERE id=$1"
	err := db.Conn.QueryRow(context.Background(), reqSql, id).Scan(&user.ID, &user.Balance)
	if err != nil {
		log.Print(err)
		if err.Error() == "no rows in result set" {
			return 0, service.NoAccountErr
		}
		return 0, err
	}
	return user.Balance, nil
}

func (db *DB) ShowBalanceTx(tx pgx.Tx, id int) (int, error) {

	var user User
	reqSql := "SELECT id, balance FROM bank.users WHERE id=$1"
	err := tx.QueryRow(context.Background(), reqSql, id).Scan(&user.ID, &user.Balance)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return 0, service.NoAccountErr
		}
		return 0, err
	}
	return user.Balance, nil
}
