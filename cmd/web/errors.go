package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

type errorTemplateData struct {
	ErrorMessage string
	ErrorCode    int
}

func (app *application) renderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	data := &errorTemplateData{
		ErrorMessage: http.StatusText(status),
		ErrorCode:    status,
	}

	err := app.errorTemplate.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.logger.errorLog.Printf("Error rendering error template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.logger.errorLog.Print(trace)
	app.renderError(w, http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.renderError(w, http.StatusNotFound)
}

func (app *application) methodNotAllowed(w http.ResponseWriter) {
	app.renderError(w, http.StatusMethodNotAllowed)
}

func (app *application) loadTemplates() error {
	files := []string{
		"./ui/html/base.html",
		"./ui/html/pages/error.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		return fmt.Errorf("failed to load templates: %w", err)
	}

	app.errorTemplate = ts
	return nil
}
