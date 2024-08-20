package Handler

import (
	"net/http"

	utils "bonus_1/helpers"
)

// FormHandler processes the form submissions
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		utils.ShowErrorPage(w, "Error: Invalid URL path.", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodPost {
		utils.ShowErrorPage(w, "Error: Unsupported HTTP method. Please use POST.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		utils.ShowErrorPage(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	str := r.FormValue("text")
	banner := r.FormValue("Banner")

	if len(str) > 1000 {
		data := Inputs{
			Message: str,
			Banner:  banner,
			Success: false,
			Err:     "Error: Message exceeds 1000 characters.",
		}
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		if err := tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
			utils.ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		}
		return
	}

	asciiArt, err := utils.GenerateBannerText(str, banner)
	if err != nil {
		if httpErr, ok := err.(utils.HttpError); ok {
			utils.ShowErrorPage(w, httpErr.Message, httpErr.StatusCode)
		} else {
			utils.ShowErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	data := Inputs{
		Message:  str,
		Banner:   banner,
		Success:  true,
		AsciiArt: asciiArt,
	}
	w.WriteHeader(http.StatusOK)
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		utils.ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
	}
}
