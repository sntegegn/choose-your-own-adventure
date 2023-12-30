package main

import "net/http"

type templateData struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

type storyForm struct {
	Choice string `form:"option"`
}

type story map[string]templateData

func (app *application) adventure(w http.ResponseWriter, r *http.Request) {
	err := app.render(w, r, http.StatusOK, "story.html", "story", "intro")
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) adventurePost(w http.ResponseWriter, r *http.Request) {
	var f storyForm
	err := app.parseForm(r, &f)
	if err != nil {
		app.serverError(w, err)
	}
	err = app.render(w, r, http.StatusOK, "story.html", "story", f.Choice)
	if err != nil {
		app.serverError(w, err)
	}
}
