package main

import (
	"fmt"
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main(){

	//загрузка .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//читаем переменные 
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	//формируем строку подключения
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user,password,dbname,host,port,
	)

	db,err := sql.Open("postgres",connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//проверка соединения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PostgreSQL")
}
