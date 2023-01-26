package main

import (
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Database connection
func connect() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Database connection failed")
		os.Exit(100)
	}else{
		log.Printf("Connection successful.\n")
	}

	if err := createSchema(db); err != nil{
		log.Fatalln("29:",err)
	}

	return db

}

func createSchema(db *pg.DB) error{
	models := []interface{}{
		(*Product)(nil),
		(*Store)(nil),
	}

	for _, model := range models{
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err!=nil{
			return err
		}
	}

	return nil
}