package main

import (
	"library/repository"

	"log"

	"github.com/joho/godotenv"
)



func main() {
	
	repo := repository.BookRepositoryImpl{}
	result, err := repo.New().Get(1)
	
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	println(result)
}