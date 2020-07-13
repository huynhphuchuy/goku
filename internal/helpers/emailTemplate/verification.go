package template

import (
	"github.com/matcornic/hermes/v2"
)

// Button Struct
type Button struct {
	Color string
	Text  string
	Link  string
}

// Verification Struct
type Verification struct {
	Name         string
	Intros       []string
	Instructions string
	Button
	Outros []string
}

// Init Verification
func (c Verification) Init() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name:   c.Name,
			Intros: c.Intros,
			Actions: []hermes.Action{
				{
					Instructions: c.Instructions,
					Button: hermes.Button{
						Color: c.Button.Color,
						Text:  c.Button.Text,
						Link:  c.Button.Link,
					},
				},
			},
			Outros: c.Outros,
		},
	}
}
