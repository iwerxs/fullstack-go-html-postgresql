package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

// Structure of Database, Table name, and Field names.
type Student struct {
	Name  string
	Klass string
	Grade string
}

func main() {
	//Serve static files from the 'static' dir
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// Serve image files from the 'assets' dir
	assets := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	connStr := "user=testusr password=testing dbname=testschdb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Server start listening on http://localhost:8080")
	// Routing Handlers for Frontend UI
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/insert.html", insertShowHandler)
	http.HandleFunc("/read.html", readShowHandler)
	http.HandleFunc("/update.html", updateShowHandler)
	http.HandleFunc("/delete.html", deleteShowHandler)

	// Function Handlers for Insert & Delete & Search & Update
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/update", updateHandler)

	// Start Server on port 8080
	http.ListenAndServe(":8080", nil)
}

// Handler Functions for redirect to frontend
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
	// parse the requesting from index.html
	tmp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}
func insertShowHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/insert.html")
	tmp, err := template.ParseFiles("templates/insert.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}
func insertHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testusr password=testing dbname=testschdb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Parse Input from the id form fields of the frontend UI
	name := r.FormValue("name")
	klass := r.FormValue("klass")
	grade := r.FormValue("grade")

	query := "INSERT INTO students (name, klass, grade) VALUES ($1, $2, $3)"

	_, err1 := db.Exec(query, name, klass, grade)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Connect to the Database
func readShowHandler(w http.ResponseWriter, r *http.Request) {

	connStr := "user=testusr password=testing dbname=testschdb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Read Records in Database
	result, err1 := db.Query("SELECT * FROM students")
	if err1 != nil {
		panic(err1)
	}
	defer result.Close()

	var students []Student
	// Load the data into the data structure
	for result.Next() {
		var student Student
		result.Scan(&student.Name, &student.Klass, &student.Grade)
		students = append(students, student)
	}

	tmp, err2 := template.ParseFiles("templates/read.html")
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, students) //pass the data to the frontend UI
}
func updateShowHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/update.html")
	tmp, err := template.ParseFiles("templates/update.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}
func deleteShowHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/delete.html")
	tmp, err := template.ParseFiles("templates/delete.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}

// delete Records from Database
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testusr password=testing dbname=testschdb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	name := r.FormValue("name")

	query := "DELETE FROM students WHERE name = $1"
	_, err1 := db.Exec(query, name)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testusr password=testing dbname=testschdb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	name := r.FormValue("name")

	query := "SELECT klass, grade FROM students WHERE name = $1"

	row := db.QueryRow(query, name)

	var klass, grade int
	err1 := row.Scan(&klass, &grade)
	if err1 != nil {
		http.Error(w, "Name not found", http.StatusNotFound)
		return
	}

	temp, err2 := template.ParseFiles("templates/update.html")
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Name  string
		Klass int
		Grade int
	}{
		Name:  name,
		Klass: klass,
		Grade: grade,
	}
	err3 := temp.Execute(w, data)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testusr password=testing dbname=testschdb host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	name := r.FormValue("name")
	klass := r.FormValue("klass")
	grade := r.FormValue("grade")

	query := "UPDATE students SET klass=$2, grade=$3 WHERE name=$1"
	_, err1 := db.Exec(query, name, klass, grade)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
