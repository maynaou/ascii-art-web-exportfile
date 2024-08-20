package Handler

import "html/template"

// Inputs structure for template data
type Inputs struct {
	Message  string // Message to convert to ASCII art
	Banner   string // Banner template to use for ASCII art
	Success  bool   // Indicates if ASCII art generation succeeded
	Err      string // Error message
	AsciiArt string // Generated ASCII art
}

var tmpl *template.Template

// SetTemplate allows main to pass the template to this package
func VetTemplate(t *template.Template) {
	tmpl = t
}
