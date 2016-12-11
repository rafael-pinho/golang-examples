package pages

import (
	"net/http"
)

func View(writer http.ResponseWriter, request *http.Request) {
	title, wikiPage, err := loadPage(request.URL.Path[len("/view/"):])

	if err != nil {
		http.Redirect(writer, request, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(writer, "view", wikiPage)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	title, wikiPage, err := loadPage(request.URL.Path[len("/edit/"):])

	if err != nil {
		wikiPage = &(page{Title: title})
	}

	renderTemplate(writer, "edit", wikiPage)
}

func Save(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/save/"):]
	body := request.FormValue("body")
	wikiPage := &(page{Title: title, Body: []byte(body)})

	wikiPage.save()
	http.Redirect(writer, request, "/view/"+title, http.StatusFound)
}
