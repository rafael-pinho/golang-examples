package pages

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates = template.Must(template.ParseFiles("views/edit.html", "views/view.html"))

type page struct {
	Title string
	Body  []byte
}

func (wikiPage *page) save() error {
	filename := "views/wiki/" + wikiPage.Title + ".html"
	return ioutil.WriteFile(filename, wikiPage.Body, 0600)
}

func loadPage(title string) (string, *page, error) {
	filename := "views/wiki/" + title + ".html"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", nil, err
	}

	return "", &page{Title: title, Body: body}, nil
}

func renderTemplate(writer http.ResponseWriter, templateName string, wikiPage *page) {
	err := templates.ExecuteTemplate(writer, templateName+".html", wikiPage)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
