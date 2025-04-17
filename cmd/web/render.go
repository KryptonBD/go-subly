package main

import (
	"fmt"
	"html/template"
	"net/http"
	"subly/data"
	"time"
)

var pathToTemplate = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	User          *data.User
}

func (app *Config) render(responseWriter http.ResponseWriter, req *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplate),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplate),
	}

	var templateSlice []string

	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplate, t))
	templateSlice = append(templateSlice, partials...)

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)

	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(responseWriter, app.AppDefaultData(td, req)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (app *Config) AppDefaultData(td *TemplateData, req *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(req.Context(), "flash")
	td.Warning = app.Session.PopString(req.Context(), "warning")
	td.Error = app.Session.PopString(req.Context(), "error")

	if app.IsAuthenticated(req) {
		td.Authenticated = true
		user, ok := app.Session.Get(req.Context(), "user").(data.User)

		if !ok {
			app.ErrorLog.Println("Can't get user from session")
		} else {
			td.User = &user
		}
	}
	td.Now = time.Now()

	return td
}

func (app *Config) IsAuthenticated(req *http.Request) bool {
	return app.Session.Exists(req.Context(), "userId")
}
