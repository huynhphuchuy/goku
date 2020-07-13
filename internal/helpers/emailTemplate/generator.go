package template

import (
	"os"

	"github.com/matcornic/hermes/v2"
)

// Export Helper
func Export(e hermes.Email) (string, string) {

	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: os.Getenv("PRODUCT_NAME"),
			Link: os.Getenv("PRODUCT_URL"),
			// Optional product logo
			Logo: os.Getenv("PRODUCT_LOGO"),
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(e)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(e)
	if err != nil {
		panic(err)
	}
	return emailBody, emailText
}
