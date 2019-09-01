package main

import (
	"identity"
	"identity/postgres"
	"log"
)

func main() {
	pg := postgres.Open()
	accountRepo := postgres.NewAccountRepo(pg)

	acc := &identity.Account{
		Email: "james@jamesapple.com",
	}

	err := accountRepo.FindByEmail(acc)
	if err != nil {
		log.Fatalf("Main Error %v", err)
	}
	log.Print(acc.Authenticate("Tes12345"))
	log.Print(acc)
}
