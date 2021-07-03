package template

import (
	"html/template"

	"github.com/marcelosilva-dev/curso-go-web-avancado/lib/contx"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": contx.I18n,
		}}
}
