package main

import (
	"root/identity"
	"root/postgres"
	"root/validator"
)

func main() {
	ar := postgres.NewAccountRepo(postgres.Open())
	s := identity.NewService(ar)
	v := validator.New()
	server := identity.NewServer(s, v)

	server.Serve()
}
