package postgres

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // Driver de conexão com o Postgres
)

// Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
