package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// RequestForSave ...
type RequestForSave struct {
	Date   string `json:"date" validate:"required"`
	Views  string `json:"views,omitempty" validate:"required"`
	Clicks string `json:"clicks,omitempty" validate:"required"`
	Cost   string `json:"cost,omitempty" validate:"required"`
}

func requiredIF(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

func (save *RequestForSave) Validate() error {
	return validation.ValidateStruct(
		save,
		validation.Field(&save.Date,
			validation.Required,
			validation.Date("2006-01-02").Error("'date' is not valid")),
	)
}

// RequestForShow ...
type RequestForShow struct {
	From      string `validate:"required"`
	To        string `validate:"required"`
	SortField string `validate:"required"`
}

func (show *RequestForShow) Validate() error {
	return validation.ValidateStruct(
		show,
		validation.Field(&show.From,
			validation.Required,
			validation.Date("2006-01-02").Error("'from' is not valid")),
		validation.Field(&show.To,
			validation.Required,
			validation.Date("2006-01-02").Error("'to' is not valid")),
		validation.Field(&show.SortField,
			validation.By(requiredIF(show.SortField != "")),
			validation.In(
				"event_date",
				"views",
				"clicks",
				"cost",
				"cpc",
				"cpm").Error(
				"field doesn`t exist, available fields [event_date, views, clicks, cost, cpc, cpm]")),
	)
}
