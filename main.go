package main

import (
  "fmt"
  "net/http"
)

type User struct {
  name string
  age uint16
  money int16
  avgGrades, happiness float64
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name %s, He is %d and he has money %d.",
  u.name, u.age, u.money)
}

func (u *User) setNewName(name string) {
  u.name = name
}

func HomePage(w http.ResponseWriter, r *http.Request) {
  user := User{"Bob", 12, -12, 4.3, 0.9}
  user.setNewName("Alex")
  fmt.Fprintf(w, user.getAllInfo())
}

func ContactsPage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Contacts Page")
}

func HandleRequest(){
  http.HandleFunc("/", HomePage)
  http.HandleFunc("/contacts/", ContactsPage)
  http.ListenAndServe(":8081", nil)
}

func main() {
  HandleRequest()
}
