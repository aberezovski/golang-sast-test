package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Session struct {
	user string
}

type Todo struct {
	Title string "json:title"
	Done  bool   "json:done"
}

func (t Todo) ToString() string {
	bytes, _ := json.Marshal(t)
	return string(bytes)
}

func getTodos() []Todo {
	todos := make([]Todo, 3)
	raw, _ := ioutil.ReadFile("./todos.json")
	json.Unmarshal(raw, &todos)
	return todos

}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Todos!",
			Todos:     getTodos(),
		}

		tmpl.Execute(w, data)

	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func DirTraversal(w http.ResponseWriter, r *http.Request, session *Session) {
	for _, f := range r.MultipartForm.File["file"] {
		input1, _ := f.Open()
		folderPath := "./list1/images/" + "/"
		for {
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				break
			}
			folderPath = "./list1/images/" + "/"
		}
		input2, _ := os.Create(folderPath + f.Filename)
		io.Copy(input2, input1)
		input1.Close()
		input2.Close()
	}
}
