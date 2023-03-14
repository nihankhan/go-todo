package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nihankhan/go-todo/config"
	"github.com/nihankhan/go-todo/models"
)

var (
	id        int
	item      string
	completed int

	view = template.Must(template.ParseFiles("./templates/index.html"))
	db   = config.Connect()
	_    = config.CreateDB()
)

func Index(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Query(`SELECT * FROM gotodo.todos`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for stmt.Next() {
		err = stmt.Scan(&id, &item, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")

	_, err := db.Exec(`INSERT INTO todos (item) VALUE(?)`, item)

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
