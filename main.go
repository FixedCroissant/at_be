// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"main/morestrings"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//Structs
type City struct {
	Id         int
	Name       string
	Population int
}

//Database connection information
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := ""
	dbPass := ""
	dbName := ""
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//Get Template Locations
var tmpl = template.Must(template.ParseGlob("templates/*" + ".html"))

//Index
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Cities ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	ourCity := City{}
	res := []City{}

	//Get Values
	for selDB.Next() {
		var id int
		var name string
		var population int
		err = selDB.Scan(&id, &name, &population)
		if err != nil {
			panic(err.Error())
		}

		//Get information
		ourCity.Id = id
		ourCity.Name = name
		ourCity.Population = population

		//Sent to response the loop of data from the query.
		res = append(res, ourCity)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

//New creation
func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//Show method
func Show(w http.ResponseWriter, r *http.Request) {
	//Get connection
	db := dbConn()

	//Get Parameter
	nId := r.URL.Query().Get("id")

	//Bind value
	res, err := db.Query("SELECT * FROM Cities WHERE id=?", nId)

	if err != nil {
		log.Fatal(err)
	}

	//Get information out of our query.
	for res.Next() {
		//Save value in a Struct
		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}

		//Print to log.
		//fmt.Printf("%v\n", city)

		tmpl.ExecuteTemplate(w, "Show", city)
	}
	defer db.Close()

}

//Insert method
func Insert(w http.ResponseWriter, r *http.Request) {
	//Get connections
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		population := r.FormValue("population")
		insForm, err := db.Prepare("INSERT INTO Cities(name, population) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		//Execute
		insForm.Exec(name, population)
		//Write to the log the change
		log.Println("INSERT: Name: " + name + " | Population: " + population)
	}
	defer db.Close()

	//Call function 2 within mail.
	//Standard message
	//mail.SendMailOurMail()

	//Custom message.
	//mail.SendMailCustom("Custom Message")

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Cities WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	city := City{}

	for selDB.Next() {
		var id int
		var population int
		var name string
		err = selDB.Scan(&id, &name, &population)
		if err != nil {
			panic(err.Error())
		}
		city.Id = id
		city.Name = name
		city.Population = population
	}
	tmpl.ExecuteTemplate(w, "Edit", city)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		population := r.FormValue("population")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Cities SET name=?, population=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, population, id)
		//Write the the log  the update.
		log.Println("UPDATE: Name: " + name + " | Population: " + population)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Cities WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//Run application.
func main() {

	//Call function 1
	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))

	log.Println("Server started on: http://localhost:8081")

	//Static Files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Handle routes
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	//End routes

	log.Fatal(http.ListenAndServe(":8081", nil))
}
