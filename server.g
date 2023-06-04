package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                 string
	Age                  uint16
	Money                int16
	AvgGrades, Happiness float64
	Hobbies              []string
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Username is %s. He is %d, and he has money %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"soccer", "ski", "saling"}}
	// bob.setNewName("Alex")
	// fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/HomePage.html")
	tmpl.Execute(w, bob)
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contactes")
}

func HandleRequest() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/contacts/", ContactPage)
	fmt.Println("I'm listening...")
	http.ListenAndServe(":8090", nil)
}

func main() {
	HandleRequest()
}
