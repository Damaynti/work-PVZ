package db

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)


func NewOn(ctx context.Context)(*Database,error) {
	
	err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }

pool,err:=pgxpool.Connect(ctx,generateDsn())
if err!=nil{
	return nil,err
}
return newDatabase(pool),nil
}

func generateDsn()string{
	
	host := os.Getenv("POSTGRES_HOST")
    portStr := os.Getenv("POSTGRES_PORT")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DBNAME")

    port, err := strconv.Atoi(portStr)
    if err != nil {
        panic("Invalid port value")
    }

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
}