package database

import (
	"log"
	"os"
	"time" // Import time

	"integracao-continua-pipelines/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")	

	stringDeConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
		
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(stringDeConexao))
		if err == nil {
			break
		}
		log.Printf("Aguardando banco de dados... (Tentativa %d/5)", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados após várias tentativas")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}