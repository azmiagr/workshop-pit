package main

import (
	"log"
	"workshop-pit/internal/handler/rest"
	"workshop-pit/internal/repository"
	"workshop-pit/internal/service"
	"workshop-pit/pkg/config"
	"workshop-pit/pkg/database/mariadb"
)

func main() {
	config.LoadEnvironment()

	db, err := mariadb.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = mariadb.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	r := rest.NewRest(svc)
	r.MountEndpoint()
	r.Run()
}
