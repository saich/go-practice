// main
package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	data, err := ioutil.ReadFile(filename)
	var p *Page
	if err == nil {
		p = &Page{Title: title, Body: data}
	}
	return p, err
}

// Parse the template files and keep it in cache
// template.Must panics in case of error (instead of just returning)
var templates = template.Must(template.ParseFiles("tmpl/view.html", "tmpl/edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Validate that title can contain only valid characters
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

func getTitle(rawTitle string, w http.ResponseWriter, r *http.Request) (title string, err error) {
	if !titleValidator.MatchString(rawTitle) {
		http.NotFound(w, r)
		err = errors.New("Invalid Title")
	}
	title = rawTitle
	return
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	const lenPath = len("/view/")
	title, err := getTitle(r.URL.Path[lenPath:], w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	const lenPath = len("/save/")
	title, err := getTitle(r.URL.Path[lenPath:], w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	const lenPath = len("/edit/")
	title, err := getTitle(r.URL.Path[lenPath:], w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func main() {

	err := os.MkdirAll("data/", 0755)
	if err != nil {
		return
	}

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/edit/", editHandler)

	http.ListenAndServe(":8080", nil)
}
