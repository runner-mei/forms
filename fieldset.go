package forms

import (
	"bytes"
	"html/template"
)

// FieldSetType is a collection of fields grouped within a form.
type FieldSetType struct {
	name     string
	class    stringSet
	tags     stringSet
	fields   []FieldInterface
	fieldMap map[string]int
}

// Render translates a FieldSetType into HTML code and returns it as a template.HTML object.
func (f *FieldSetType) Render() template.HTML {
	var s string
	buf := bytes.NewBufferString(s)
	data := map[string]interface{}{
		"fields":  f.fields,
		"classes": f.class,
		"tags":    f.tags,
	}
	err := template.Must(template.ParseFiles(CreateURL("templates/fieldset.html"))).Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

// FieldSet creates and returns a new FieldSetType with the given name and list of
// Every method for FieldSetType objects returns the object itself, so that call can be chained.
func FieldSet(name string, elems ...FieldInterface) *FieldSetType {
	ret := &FieldSetType{
		name,
		stringSet{},
		stringSet{},
		elems,
		map[string]int{},
	}
	for i, elem := range elems {
		ret.fieldMap[elem.Name()] = i
	}
	return ret
}

// Field returns the field identified by name. It returns an empty field if it is missing.
func (f *FieldSetType) Field(name string) FieldInterface {
	ind, ok := f.fieldMap[name]
	if !ok {
		return &Field{}
	}
	return f.fields[ind].(FieldInterface)
}

// Name returns the name of the fieldset.
func (f *FieldSetType) Name() string {
	return f.name
}

// AddClass saves the provided class for the fieldset.
func (f *FieldSetType) AddClass(class string) *FieldSetType {
	f.class = f.class.add(class)
	return f
}

// RemoveClass removes the provided class from the fieldset, if it was present. Nothing is done if it was not originally present.
func (f *FieldSetType) RemoveClass(class string) *FieldSetType {
	f.class = f.class.remove(class)
	return f
}

// AddTag adds a no-value parameter (e.g.: "disabled", "checked") to the fieldset.
func (f *FieldSetType) AddTag(tag string) *FieldSetType {
	f.tags = f.tags.add(tag)
	return f
}

// RemoveTag removes a tag from the fieldset, if it was present.
func (f *FieldSetType) RemoveTag(tag string) *FieldSetType {
	f.tags = f.tags.remove(tag)
	return f
}

// Disable adds tag "disabled" to the fieldset, making it unresponsive in some environment (e.g.: Bootstrap).
func (f *FieldSetType) Disable() *FieldSetType {
	f.AddTag("disabled")
	return f
}

// Enable removes tag "disabled" from the fieldset, making it responsive.
func (f *FieldSetType) Enable() *FieldSetType {
	f.RemoveTag("disabled")
	return f
}
