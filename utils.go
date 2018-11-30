package forms

import (
	"bytes"
	"html/template"
	"reflect"
)

// FormElement interface defines a form object (usually a Field or a FieldSet) that can be rendered as a template.HTML object.
type Renderable interface {
	Render(theme string) template.HTML
}

// Render executes the internal template and renders the form, returning the result as a template.HTML object embeddable
// in any other template.
func (f *Form) Render(theme string) template.HTML {
	var s string
	buf := bytes.NewBufferString(s)
	data := map[string]interface{}{
		"fields":  f.fields,
		"classes": f.class,
		"id":      f.id,
		"params":  f.params,
		"css":     f.css,
		"method":  f.method,
		"action":  f.action,
	}
	err := f.template.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return template.HTML(buf.String())
}

// Elements adds the provided elements to the form.
func (f *Form) Elements(elems ...Renderable) *Form {
	for _, e := range elems {
		t := reflect.TypeOf(e)
		switch {
		case t.Implements(reflect.TypeOf((*FieldInterface)(nil)).Elem()):
			f.addField(e.(FieldInterface))
		case reflect.ValueOf(e).Type().String() == "*forms.FieldSetType":
			f.addFieldSet(e.(*FieldSetType))
		}
	}
	return f
}

func (f *Form) addField(field FieldInterface) *Form {
	field.SetTheme(f.style)
	f.fields = append(f.fields, field)
	f.fieldMap[field.Name()] = len(f.fields) - 1
	return f
}

func (f *Form) addFieldSet(fs *FieldSetType) *Form {
	for _, v := range fs.fields {
		v.SetTheme(f.style)
		f.containerMap[v.Name()] = fs.name
	}
	f.fields = append(f.fields, fs)
	f.fieldMap[fs.Name()] = len(f.fields) - 1
	return f
}

// RemoveElement removes an element (identified by name) from the Form.
func (f *Form) RemoveElement(name string) *Form {
	ind, ok := f.fieldMap[name]
	if !ok {
		return f
	}
	delete(f.fieldMap, name)
	f.fields = append(f.fields[:ind], f.fields[ind+1:]...)
	return f
}

// AddClass associates the provided class to the Form.
func (f *Form) AddClass(class string) *Form {
	f.class = append(f.class, class)
	return f
}

// RemoveClass removes the given class (if present) from the Form.
func (f *Form) RemoveClass(class string) *Form {
	ind := -1
	for i, v := range f.class {
		if v == class {
			ind = i
			break
		}
	}

	if ind != -1 {
		f.class = append(f.class[:ind], f.class[ind+1:]...)
	}
	return f
}

// SetId set the given id to the form.
func (f *Form) SetId(id string) *Form {
	f.id = id
	return f
}

// SetParam adds the given key-value pair to form parameters list.
func (f *Form) SetParam(key, value string) *Form {
	f.params[key] = value
	return f
}

// DeleteParm removes the parameter identified by key from form parameters list.
func (f *Form) DeleteParam(key string) *Form {
	delete(f.params, key)
	return f
}

// AddCss add a CSS value (in the form of option-value - e.g.: border - auto) to the form.
func (f *Form) AddCss(key, value string) *Form {
	f.css[key] = value
	return f
}

// RemoveCss removes CSS style from the form.
func (f *Form) RemoveCss(key string) *Form {
	delete(f.css, key)
	return f
}

// Field returns the field identified by name. It returns an empty field if it is missing.
func (f *Form) Field(name string) FieldInterface {
	ind, ok := f.fieldMap[name]
	if !ok || !reflect.TypeOf(f.fields[ind]).Implements(reflect.TypeOf((*FieldInterface)(nil)).Elem()) {
		if v, ok2 := f.containerMap[name]; ok2 {
			return f.FieldSet(v).Field(name)
		}
		return &Field{}
	}
	return f.fields[ind].(FieldInterface)
}

// Field returns the field identified by name. It returns an empty field if it is missing.
func (f *Form) FieldSet(name string) *FieldSetType {
	ind, ok := f.fieldMap[name]
	if !ok || reflect.ValueOf(f.fields[ind]).Type().String() != "*forms.FieldSetType" {
		return &FieldSetType{}
	}
	return f.fields[ind].(*FieldSetType)
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
