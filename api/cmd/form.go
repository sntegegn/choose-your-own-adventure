package main

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/sntegegn/choose-your-own-adventure/ui"
)

func (app *application) getData(choice string) (templateData, error) {
	if app.gopherStory == nil {
		err := app.readJSON(&app.gopherStory)
		if err != nil {
			return templateData{}, err
		}
	}

	data := templateData{
		Title:   app.gopherStory[choice].Title,
		Story:   app.gopherStory[choice].Story,
		Options: app.gopherStory[choice].Options,
	}
	return data, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, statusCode int, page, templateName string, choice string) error {
	data, err := app.getData(choice)
	if err != nil {
		return err
	}

	ts, err := template.New("story").ParseFS(ui.Files, "html/"+page)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = ts.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	buf.WriteTo(w)
	return nil
}

func (app *application) parseForm(r *http.Request, f *storyForm) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formdecoder.Decode(f, r.PostForm)
	if err != nil {
		return err
	}

	return nil
}
