package main

import (
	"encoding/json"
	"net/http"

	"github.com/sntegegn/choose-your-own-adventure/internal/data"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.logger.Error(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func (app *application) readJSON(sd *story) error {
	val, err := data.JsonData.ReadFile(app.cfg.jsonPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(val, sd); err != nil {
		return err
	}
	return nil
}
