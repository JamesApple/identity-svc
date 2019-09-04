package main

import (
	"log"
	"os"
	"root/postgres"
)

func main() {
	db := postgres.Connect()
	instrumentedDB := &postgres.QueryLogger{DB: db, Logger: log.New(os.Stdout, "[DB] ", 0)}

	err := postgres.Drop(instrumentedDB)
	if err != nil {
		log.Fatalf("Could not drop, %v", err)
	}

	err = postgres.Migrate(instrumentedDB)
	if err != nil {
		log.Fatalf("Could not migrate, %v", err)
	}

	r := &postgres.AccountRepo{QueryLogger: instrumentedDB}

	r.Find(1)

}
