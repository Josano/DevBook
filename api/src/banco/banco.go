package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexãoBanco)

	if erro != nil {
		return nil, erro
	}

	if error := db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}