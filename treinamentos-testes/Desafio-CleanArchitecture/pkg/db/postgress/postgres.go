package postgress

import (
	"database/sql"

	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/wrferreira1003/Desafio-Clean-Architecture/config"
)

type PostgresDB struct {
	Config *config.Config
}

// Conecta com o banco de dados Postgres
// NewDBConnection cria e retorna uma nova conexão com o banco de dados PostgreSQL
func NewDBConnection(postgres *PostgresDB) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgres.Config.DBHost, postgres.Config.DBPort, postgres.Config.DBUser, postgres.Config.DBPassword, postgres.Config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	// Verifica se a conexão com o banco foi bem-sucedida
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com o banco de dados: %v", err)
		return nil, err
	}

	log.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")
	return db, nil
}
