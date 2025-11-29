package main

import (
	"log"
	"workshop-pit/internal/handler/rest"
	"workshop-pit/internal/repository"
	"workshop-pit/internal/service"
	"workshop-pit/pkg/bcrypt"
	"workshop-pit/pkg/config"
	"workshop-pit/pkg/database/mariadb"
	"workshop-pit/pkg/jwt"
	"workshop-pit/pkg/middleware"
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
	bcrypt := bcrypt.Init()
	jwt := jwt.Init()
	svc := service.NewService(repo, bcrypt, jwt)
	m := middleware.Init(svc, jwt)
	r := rest.NewRest(svc, m)
	r.MountEndpoint()
	r.Run()
}
