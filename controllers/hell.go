package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Person struct {
	ID      int
	Name    string
	Address string
	Rank    int
}

func HellWorld(w http.ResponseWriter, r *http.Request) {
	dbString := getDBString()
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT id, name, address, rank FROM users;")
	persons := []Person{}
	for rows.Next() {
		p := Person{}
		err := rows.Scan(&p.ID, &p.Name, &p.Address, &p.Rank)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, p)
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Persons []Person
	}{
		Persons: persons,
	}
	err = t.Execute(w, data)
	//fmt.Fprint(w, "hell world")
}

func getDBString() string {
	return fmt.Sprintf("postgres://%s:@%s:%s/%s?sslmode=%s",
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_SSLMODE"),
	)
}
