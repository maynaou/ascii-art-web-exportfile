package Handler

import (
	"net/http"

	utils "bonus_1/helpers"
)

// HomeHandler serves the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		utils.ShowErrorPage(w, "Error: Invalid URL path.", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		data := Inputs{}
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			utils.ShowErrorPage(w, "Error executing template", http.StatusInternalServerError)
		}
		return
	} else {
		utils.ShowErrorPage(w, "Error: Unsupported HTTP method. Please use GET.", http.StatusMethodNotAllowed)
		return
	}
}
