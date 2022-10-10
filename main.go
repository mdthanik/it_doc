package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	fmt.Printf("main %s\n", os.Getenv("DB_NAME"))
}