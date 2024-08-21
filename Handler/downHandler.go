package Handler

import (
	"html/template"
	"net/http"
	"strconv"

	utils "bonus_1/helpers"
)

// DownHandler handles the download requests
func DownHandler(w http.ResponseWriter, r *http.Request) {
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
	fileType := r.FormValue("fileType")

	asciiArt, err := utils.GenerateBannerText(str, banner)
	if err != nil {
		if httpErr, ok := err.(utils.HttpError); ok {
			utils.ShowErrorPage(w, httpErr.Message, httpErr.StatusCode)
		} else {
			utils.ShowErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	if fileType == "html" {
		// Serve as HTML using a template
		w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.html")
		w.Header().Set("Content-Type", "text/html")

		// Load the template
		tmpl := template.Must(template.ParseFiles("Templates/ascii_art.html"))

		// Pass the ASCII Art to the template
		data := struct {
			AsciiArt string
		}{
			AsciiArt: asciiArt,
		}

		if err := tmpl.Execute(w, data); err != nil {
			utils.ShowErrorPage(w, "Error rendering template", http.StatusInternalServerError)
		}
	} else {
		// Serve as TXT (default)
		w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(asciiArt)))
		w.Write([]byte(asciiArt))
	}
}
