package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/RipulHandoo/jwt/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBInstance() *database.Queries{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Could not load .env file")
	}

	db_url := os.Getenv("DB_URL")
	if db_url == ""{
		log.Fatal("Could not get db_url from .env file")
	}

	db, err := sql.Open("postgres",db_url)
	if err != nil {
		log.Fatal("Could not connect with database")
	}

	dbQueries := database.New(db)

		return dbQueries
}


var DbClient *database.Queries = DBInstance()