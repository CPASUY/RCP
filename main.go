package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseFiles("index.ejs"))
var conection *sql.DB
var err error

func main() {
	conection, err = sql.Open("mysql", "root:@/rcp")
	if err != nil {
		panic(err.Error())
	}
	err = conection.Ping()
	if err != nil {
		panic(err.Error())
	}
	http.HandleFunc("/login", loadLogin)
	http.HandleFunc("/register", registerComp)
	http.HandleFunc("/index", loadUsers)
	http.HandleFunc("/", loadLogin)
	http.ListenAndServe(":8080", nil)
}

var u string

func registerComp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "register.ejs")
		return
	}
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmpwd := r.FormValue("confirmpwd")
	country := r.FormValue("country")
	fmt.Println("Pasword:" + password)
	fmt.Println("Confirm Pasword:" + confirmpwd)
	err := conection.QueryRow("SELECT Username FROM usuarios WHERE username=?", username).Scan(&u)
	if err == sql.ErrNoRows {
		fmt.Print("UNO")
		if len(username) > 5 && len(username) < 20 || len(password) > 0 {
			fmt.Print("DPS")
			if len(confirmpwd) > 0 || len(country) > 0 {
				fmt.Print("TRES")
				if password == confirmpwd {
					fmt.Print("CUATRO")
					_, err = conection.Exec("INSERT INTO usuarios(Username, Firstname, Lastname, Email,Password,Country) VALUES(?, ?,?,?,?,?)", username, firstname, lastname, email, password, country)
				}
			}
		}
	}
	http.Redirect(w, r, "/", 301)
	return
}

var us string
var pw string

func loadLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "login.ejs")
		return
	}
	username := r.FormValue("username")
	//password := r.FormValue("password")
	err := conection.QueryRow("SELECT Username, Password FROM usuarios WHERE Username=?", username).Scan(&us, &pw)
	if err != nil {
		http.Redirect(w, r, "/login", 301)
		return
	}
	http.Redirect(w, r, "/index", 301)
}
func loadUsers(w http.ResponseWriter, r *http.Request) {
	comps, err := conection.Query("SELECT Username, Firstname, Lastname FROM usuarios")

	if err != nil {
		http.Error(w, "ERROR", 500)
		return
	}
	competitor := Competitor{}
	competitors := []Competitor{}
	for comps.Next() {
		var u string
		var f string
		var l string
		err = comps.Scan(&u, &f, &l)
		if err != nil {
			http.Error(w, "ERROR", 500)
			return
		}
		competitor.username = u
		competitor.firstname = f
		competitor.lastname = l
		fmt.Println(u)
		fmt.Println(f)
		fmt.Println(l)

		competitors = append(competitors, competitor)

	}
	if err := templates.Execute(w, competitors); err != nil {
		fmt.Print("nulo")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Competitor struct {
	firstname, lastname, username string
}
