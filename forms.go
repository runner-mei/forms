// This package provides form creation and rendering functionalities, as well as FieldSet definition.
// Two kind of forms can be created: base forms and Bootstrap3 compatible forms; even though the latters are automatically provided
// the required classes to make them render correctly in a Bootstrap environment, every form can be given custom parameters such as
// classes, id, generic parameters (in key-value form) and stylesheet options.
package forms

import (
	"html/template"
)

// Form methods: POST or GET.
const (
	POST = "POST"
	GET  = "GET"
)

// Form structure.
type Form struct {
	fields       []Renderable
	fieldMap     map[string]int
	containerMap map[string]string
	style        string
	template     *template.Template
	class        []string
	id           string
	params       map[string]string
	css          map[string]string
	method       string
	action       template.HTML
}

// BaseForm creates an empty form with no styling.
func BaseForm(method, action string) *Form {
	tmpl, err := template.ParseFiles(CreateURL("templates/baseform.html"))
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]Renderable, 0),
		make(map[string]int),
		make(map[string]string),
		BASE,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		template.HTML(action),
	}
}

// BootstrapForm creates an empty form compliant with Bootstrap3 CSS, both in structure and classes.
func BootstrapForm(method, action string) *Form {
	tmpl, err := template.ParseFiles(CreateURL("templates/baseform.html"))
	if err != nil {
		panic(err)
	}
	return &Form{
		make([]Renderable, 0),
		make(map[string]int),
		make(map[string]string),
		BOOTSTRAP,
		tmpl,
		[]string{},
		"",
		map[string]string{},
		map[string]string{},
		method,
		template.HTML(action),
	}
}
