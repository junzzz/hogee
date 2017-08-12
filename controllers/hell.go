package controllers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Person struct {
	ID      int
	Name    string
	Address string
	Rank    int
}

func HellWorld(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "postgres://junzzz:@localhost/postgres?sslmode=disable")
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

	//t := template.New("sample templete")
	t, err := template.ParseFiles("/Users/junzzz/src/github.com/junzzz/hogee/templates/index.html")
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Persons []Person
	}{
		Persons: persons,
	}
	err = t.Execute(w, data)
	//log.Println(err)
	//fmt.Fprintf(w, "%d, %s, %s, %d", id, name, address, rank)
}
