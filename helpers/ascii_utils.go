package utils

import (
	"html/template"
	"net/http"
	"strings"

	"bonus_1/bonus1"
)

var tmpl *template.Template

// SetTemplate allows main to pass the template to this package
func SetTemplate(t *template.Template) {
	tmpl = t
}

// Function to handle form requests
func GenerateBannerText(str, banner string) (string, error) {
	var result []string
	res := bonus1.ProcessInput(str)
	if res == nil {
		bannerText, err := bonus1.GenerateBanner(nil, str, banner)
		if err != nil {
			return "", ClassifyError(err)
		}
		result = append(result, bannerText)
	} else {
		for _, s := range res {
			bannerText, err := bonus1.GenerateBanner(res, s, banner)
			if err != nil {
				return "", ClassifyError(err)
			}
			result = append(result, bannerText)
		}
	}
	return strings.Join(result, "\n"), nil
}

func (e HttpError) Error() string {
	return e.Message
}

// Function to classify errors and determine the HTTP status code
func ClassifyError(err error) error {
	var status int
	switch {
	case err.Error() == "text not found":
		status = http.StatusNotFound
	case err.Error() == "invalid banner type":
		status = http.StatusNotFound
	case strings.Contains(err.Error(), "please provide printable characters"):
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	return HttpError{StatusCode: status, Message: err.Error()}
}

// Function to show error page
func ShowErrorPage(w http.ResponseWriter, errMsg string, statusCode int) {
	data := struct {
		Err string
	}{
		Err: errMsg,
	}
	w.WriteHeader(statusCode)
	if err := tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
		http.Error(w, "Error displaying the error page", http.StatusInternalServerError)
	}
}

type HttpError struct {
	StatusCode int
	Message    string
}
