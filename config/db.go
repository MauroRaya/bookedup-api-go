package config

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "app.db")

	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS Usuario (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Nome TEXT NOT NULL,
			Email TEXT NOT NULL UNIQUE
		);
	`)

	if err != nil {
		panic(fmt.Sprintf("Erro ao criar tabela: %v", err))
	}
}
