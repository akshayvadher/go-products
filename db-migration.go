package main

import (
	"log"
	"training/go-products/config"
	"training/go-products/persistance"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func DbMigration() error {
	log.Println("Starting DB migrate")

	config.ReadConfig()
	connectionString := "mysql://" + persistance.GetConnectionString()

	m, err := migrate.New("file://sql", connectionString)
	if err != nil {
		log.Println("error creating migrate object " + err.Error())
		return err
	}
	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Println("No DB change found")
			return nil
		}
		log.Println("Error running DB migrate: " + err.Error())
		return err
	}

	return nil
}
