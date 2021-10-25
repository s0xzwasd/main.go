package main
import (
	"html/template"
	"log"
	"net/http"
)

func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//для простоты не обрабатываем ошибки
	t, _ := template.ParseFiles("test.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandlerFunc)
	http.HandleFunc("/login", loginHandlerFunc)
	log.Fatal(http.ListenAndServe("localhost:3050", nil))
}

func loginHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//Обрабатываем только POST-запрос
	if r.Method != "POST" {
		http.NotFound(w, r)
	}
	//для простоты не обрабатываем ошибки
	r.ParseForm()
	user := r.FormValue("user")
	password := r.FormValue("password")
	//Проверяем логин и пароль
	if !(user == "zaz600" && password == "123") {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	w.Write([]byte("hello " + user))
}
