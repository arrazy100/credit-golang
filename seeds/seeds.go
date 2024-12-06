package main

import (
	"credit/config"
	"log"
)

func main() {
	configs, err := config.Load("config.dev.yaml")
	if err != nil {
		panic(err)
	}

	err = SeedInitial(configs)
	if err != nil {
		panic(err)
	}

	log.Println("Finished seeding")
}
