package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var templates = template.Must(template.ParseFiles("index.html"))
var conection *sql.DB

func main() {
	conection, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/Rcp")
	if err != nil {
		panic(err.Error())
	}
	defer conection.Close()

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
		fmt.Print("post")
		http.ServeFile(w, r, "register.html")
		return
	}
	fmt.Print("postcuerpo")
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	password := r.FormValue("password")
	confirmpwd := r.FormValue("confirmPassword")
	country := r.FormValue("country")
	err := conection.QueryRow("SELECT Username FROM usuarios WHERE username=?", username).Scan(&u)
	if err == sql.ErrNoRows {
		if err == nil {
			insertuser, err := conection.Prepare("INSERT INTO usuarios(Username,Firstname,Lastname,Email,Password,ConfirmPwd,Country) VALUES(username,firstname,lastname,password,confirmpwd,country)")
			if err != nil {
				panic(err.Error())
			} else {
				if len(firstname) > 0 || len(lastname) > 0 {
					if len(username) > 5 && len(username) < 20 || len(password) > 0 {
						if len(confirmpwd) > 0 || len(country) > 0 {
							if password == confirmpwd {
								insertuser.Exec(username, firstname, lastname, password, confirmpwd, country)
							}
						}
					}
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
		http.ServeFile(w, r, "login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	err := conection.QueryRow("SELECT Username, Password FROM usuarios WHERE Username=? AND Password=?", username, password).Scan(&us, &pw)
	fmt.Print("erorrrrrrrrrrrrrrrr")
	if err != nil {
		fmt.Print("primero")
		http.Redirect(w, r, "/login", 301)
		return
	} else {
		fmt.Print("segundo")
		http.Redirect(w, r, "/index", 301)
	}
}
func loadUsers(w http.ResponseWriter, r *http.Request) {

	users, err := conection.Query("SELECT Username, Firstname, Lastname FROM usuarios")

	if err != nil {
		http.Error(w, "ERROR", 500)
		return
	}
	competitor := Competitor{}
	competitors := []Competitor{}
	for users.Next() {
		var u string
		var f string
		var l string
		err = users.Scan(&u, &f, &l)
		if err != nil {
			http.Error(w, "ERROR", 500)
			return
		}
		competitor.username = u
		competitor.firstname = f
		competitor.lastname = l
		competitors = append(competitors, competitor)
	}
	if err := templates.Execute(w, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Competitor struct {
	firstname  string
	lastname   string
	username   string
	password   string
	confirmPwd string
	country    string
}
