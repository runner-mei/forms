// This package provides all the input fields logic and customization methods.
package forms

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

// Field is a generic type containing all data associated to an input field.
type Field struct {
	ctx            interface{}
	fieldType      string
	Widget         Widget // Public Widget field for widget customization
	name           string
	classes        stringSet
	id             string
	params         map[string]string
	css            map[string]string
	label          string
	labelClasses   stringSet
	tags           stringSet
	valueLoaded    bool
	value          string
	helptext       string
	errors         []string
	additionalData map[string]interface{}
}

// FieldInterface defines the interface an object must implement to be used in a form. Every method returns a FieldInterface object
// to allow methods chaining.
type FieldInterface interface {
	Name() string
	Render() template.HTML
	AddClass(class string) FieldInterface
	AddData(key, value string) FieldInterface
	RemoveClass(class string) FieldInterface
	AddTag(class string) FieldInterface
	RemoveTag(class string) FieldInterface
	SetID(id string) FieldInterface
	SetParam(key, value string) FieldInterface
	DeleteParam(key string) FieldInterface
	AddCSS(key, value string) FieldInterface
	RemoveCSS(key string) FieldInterface
	SetStyle(style string) FieldInterface
	SetLabel(label string) FieldInterface
	AddLabelClass(class string) FieldInterface
	RemoveLabelClass(class string) FieldInterface
	SetValue(value interface{}) FieldInterface
	Disabled() FieldInterface
	Enabled() FieldInterface
	SetHelptext(text string) FieldInterface
	AddError(err string) FieldInterface
	MultipleChoice() FieldInterface
	SingleChoice() FieldInterface
	AddSelected(opt ...string) FieldInterface
	RemoveSelected(opt string) FieldInterface
	SetSelectChoices(choices map[string][]InputChoice) FieldInterface
	SetRadioChoices(choices []InputChoice) FieldInterface
	SetText(text string) FieldInterface
}

// FieldWithType creates an empty field of the given type and identified by name.
func FieldWithType(ctx interface{}, name, t string) *Field {
	return &Field{
		ctx:            ctx,
		fieldType:      t,
		Widget:         nil,
		name:           name,
		classes:        stringSet{},
		id:             "",
		params:         map[string]string{},
		css:            map[string]string{},
		label:          "",
		labelClasses:   stringSet{},
		tags:           stringSet{},
		value:          "",
		helptext:       "",
		errors:         []string{},
		additionalData: map[string]interface{}{},
	}
}

// FieldWithTypeWithCtx creates an field of the given type and identified by name.
func FieldWithTypeWithCtx(ctx interface{}, name, label, typ string) *Field {
	field := FieldWithType(ctx, name, typ)
	field.SetLabel(label)
	return field
}

// SetStyle sets the style (e.g.: BASE, BOOTSTRAP) of the field, correctly populating the Widget field.
func (f *Field) SetStyle(style string) FieldInterface {
	f.Widget = BaseWidget(style, f.fieldType)
	return f
}

// Name returns the name of the field.
func (f *Field) Name() string {
	return strings.TrimSuffix(f.name, "[]")
}

func (f *Field) loadValueIfNeed() {
	if f.valueLoaded {
		return
	}

	if f.ctx == nil {
		return
	}

	name := f.Name()
	ctx := f.ctx.(map[string]interface{})

	value := revel.NewField(name, ctx)
	f.setValue(value)

	if value.Error != nil {
		f.AddError(value.Error.Message)
	}
}

type FieldValue interface {
	Flash() string
	FlashArray() []string
	Value() interface{}
}

type fieldValue struct {
	value interface{}
}

func (fv fieldValue) Flash() string {
	return ""
}
func (fv fieldValue) FlashArray() []string {
	return nil
}
func (fv fieldValue) Value() interface{} {
	return fv.value
}

func (f *Field) setValue(value FieldValue) {
	if MAP == f.fieldType {
		if flashValue := value.Flash(); "" != flashValue {
			f.SetText(flashValue)
		} else {
			value := value.Value()
			if value == nil {
				f.SetText("")
			} else {
				f.SetText(MapToString(value))
			}
		}
	} else if TEXTAREA == f.fieldType {
		if flashValue := value.Flash(); "" != flashValue {
			f.SetText(flashValue)
		} else {
			value := value.Value()
			if value == nil {
				f.SetText("")
			} else {
				f.SetText(fmt.Sprint(value))
			}
		}
	} else if SELECT == f.fieldType && f.IsMultipleChoice() {
		if flashArray := value.FlashArray(); len(flashArray) > 0 {
			f.AddSelected(flashArray...)
		} else {
			opts := toStringArray(value.Value(), nil)
			if len(opts) != 0 {
				f.AddSelected(opts...)
			}
		}
	} else {
		if flashValue := value.Flash(); "" != flashValue {
			f.value = flashValue
		} else {
			val := value.Value()
			if b, ok := val.(byte); ok {
				f.value = string(b)
			} else {
				if value.Value() != nil {
					f.value = fmt.Sprint(value.Value())
				}
			}
		}
	}
	f.valueLoaded = true
}

func (f *Field) dataForRender() map[string]interface{} {
	f.loadValueIfNeed()

	safeParams := make(map[template.HTMLAttr]string)
	for k, v := range f.params {
		safeParams[template.HTMLAttr(k)] = v
	}
	data := map[string]interface{}{
		"ctx_parent":   f.ctx,
		"classes":      f.classes,
		"id":           f.id,
		"name":         f.name,
		"params":       safeParams,
		"css":          f.css,
		"type":         f.fieldType,
		"label":        f.label,
		"labelClasses": f.labelClasses,
		"tags":         f.tags,
		"value":        f.value,
		"helptext":     f.helptext,
		"errors":       f.errors,
	}
	for k, v := range f.additionalData {
		data[k] = v
	}
	return data
}

// Render packs all data and executes widget render method.
func (f *Field) Render() template.HTML {
	if f.Widget == nil {
		f.Widget = BaseWidget("bootstrap3", f.fieldType)
	}

	if f.Widget != nil {
		data := f.dataForRender()
		return template.HTML(f.Widget.Render(data))
	}
	return template.HTML("field template is not found.")
}

func (f *Field) String() string {
	return string(f.Render())
}

// AddClass adds a class to the field.
func (f *Field) AddClass(class string) FieldInterface {
	f.classes = f.classes.add(class)
	return f
}

// AddData adds a k/v to the additional data.
func (f *Field) AddData(key, value string) FieldInterface {
	f.additionalData[key] = value
	return f
}

// RemoveClass removes a class from the field, if it was present.
func (f *Field) RemoveClass(class string) FieldInterface {
	f.classes = f.classes.remove(class)
	return f
}

// SetID associates the given id to the field, overwriting any previous id.
func (f *Field) SetID(id string) FieldInterface {
	f.id = id
	return f
}

// SetLabel saves the label to be rendered along with the field.
func (f *Field) SetLabel(label string) FieldInterface {
	f.label = label
	return f
}

// AddLabelClass allows to define custom classes for the label.
func (f *Field) AddLabelClass(class string) FieldInterface {
	f.labelClasses = f.labelClasses.add(class)
	return f
}

// RemoveLabelClass removes the given class from the field label.
func (f *Field) RemoveLabelClass(class string) FieldInterface {
	f.labelClasses = f.labelClasses.remove(class)
	return f
}

// SetParam adds a parameter (defined as key-value pair) in the field.
func (f *Field) SetParam(key, value string) FieldInterface {
	f.params[key] = value
	return f
}

// SetIntParam adds a parameter (defined as key-value pair) in the field.
func (f *Field) SetIntParam(key string, value int) FieldInterface {
	f.params[key] = strconv.FormatInt(int64(value), 10)
	return f
}

// DeleteParam removes a parameter identified by key from the field.
func (f *Field) DeleteParam(key string) FieldInterface {
	delete(f.params, key)
	return f
}

// AddCSS adds a custom CSS style the field.
func (f *Field) AddCSS(key, value string) FieldInterface {
	f.css[key] = value
	return f
}

// RemoveCSS removes CSS options identified by key from the field.
func (f *Field) RemoveCSS(key string) FieldInterface {
	delete(f.css, key)
	return f
}

// Disabled add the "disabled" tag to the field, making it unresponsive in some environments (e.g. Bootstrap).
func (f *Field) Disabled() FieldInterface {
	f.AddTag("disabled")
	return f
}

// Enabled removes the "disabled" tag from the field, making it responsive.
func (f *Field) Enabled() FieldInterface {
	f.RemoveTag("disabled")
	return f
}

// AddTag adds a no-value parameter (e.g.: checked, disabled) to the field.
func (f *Field) AddTag(tag string) FieldInterface {
	f.tags = f.tags.add(tag)
	return f
}

// RemoveTag removes a no-value parameter from the field.
func (f *Field) RemoveTag(tag string) FieldInterface {
	f.tags = f.tags.remove(tag)
	return f
}

// HasTag has a no-value parameter in the field.
func (f *Field) HasTag(tag string) bool {
	return f.tags.has(tag)
}

// SetValue sets the value parameter for the field.
func (f *Field) SetValue(value interface{}) FieldInterface {
	f.setValue(fieldValue{value})
	return f
}

// SetHelptext saves the field helptext.
func (f *Field) SetHelptext(text string) FieldInterface {
	f.helptext = text
	return f
}

// AddError adds an error string to the field. It's valid only for Bootstrap forms.
func (f *Field) AddError(err string) FieldInterface {
	f.errors = append(f.errors, err)
	return f
}

// MultipleChoice configures the SelectField to accept and display multiple choices.
// It has no effect if type is not SELECT.
func (f *Field) MultipleChoice() FieldInterface {
	if f.fieldType == SELECT {
		f.AddTag("multiple")
		// fix name if necessary
		if !strings.HasSuffix(f.name, "[]") {
			f.name = f.name + "[]"
		}
	}
	return f
}

// IsMultipleChoice is a multiple choices.
func (f *Field) IsMultipleChoice() bool {
	return f.HasTag("multiple")
}

// SingleChoice configures the Field to accept and display only one choice (valid for SelectFields only).
// It has no effect if type is not SELECT.
func (f *Field) SingleChoice() FieldInterface {
	if f.fieldType == SELECT {
		f.RemoveTag("multiple")
		if strings.HasSuffix(f.name, "[]") {
			f.name = strings.TrimSuffix(f.name, "[]")
		}
	}
	return f
}

// AddSelected If the field is configured as "multiple", AddSelected adds a selected value to the field (valid for SelectFields only).
// It has no effect if type is not SELECT.
func (f *Field) AddSelected(opt ...string) FieldInterface {
	if f.fieldType == SELECT {
		for _, v := range opt {
			f.additionalData["multValues"].(map[string]struct{})[v] = struct{}{}
		}
	}
	return f
}

// RemoveSelected If the field is configured as "multiple", AddSelected removes the selected value from the field (valid for SelectFields only).
// It has no effect if type is not SELECT.
func (f *Field) RemoveSelected(opt string) FieldInterface {
	if f.fieldType == SELECT {
		if _, ok := f.additionalData["multValues"]; ok {
			delete(f.additionalData["multValues"].(map[string]struct{}), opt)
		}
	}
	return f
}

// SetSelectChoices takes as input a dictionary whose key-value entries are defined as follows: key is the group name (the empty string
// is the default group that is not explicitly rendered) and value is the list of choices belonging to that group.
// Grouping is only useful for Select fields, while groups are ignored in Radio fields.
// It has no effect if type is not SELECT.
func (f *Field) SetSelectChoices(choices map[string][]InputChoice) FieldInterface {
	if f.fieldType == SELECT {
		f.additionalData["choices"] = choices
	}
	return f
}

// SetRadioChoices sets an array of InputChoice objects as the possible choices for a radio field. It has no effect if type is not RADIO.
func (f *Field) SetRadioChoices(choices []InputChoice) FieldInterface {
	if f.fieldType == RADIO {
		f.additionalData["choices"] = choices
	}
	return f
}

// SetText saves the provided text as content of the field, usually a TextAreaField.
func (f *Field) SetText(text string) FieldInterface {
	if f.fieldType == BUTTON ||
		f.fieldType == SUBMIT ||
		f.fieldType == RESET ||
		f.fieldType == STATIC ||
		f.fieldType == TEXTAREA {
		f.additionalData["text"] = text
	}
	return f
}

var (
	FieldFuncs = template.FuncMap{
		"f_addClass": func(class string, field FieldInterface) FieldInterface {
			field.AddClass(class)
			return field
		},
		"f_removeClass": func(class string, field FieldInterface) FieldInterface {
			field.RemoveClass(class)
			return field
		},
		"f_addTag": func(class string, field FieldInterface) FieldInterface {
			field.AddTag(class)
			return field
		},
		"f_removeTag": func(class string, field FieldInterface) FieldInterface {
			field.RemoveTag(class)
			return field
		},
		"f_setId": func(id string, field FieldInterface) FieldInterface {
			field.SetID(id)
			return field
		},
		"f_addParams": func(key, value string, field FieldInterface) FieldInterface {
			field.SetParam(key, value)
			return field
		},
		"f_removeParams": func(key string, field FieldInterface) FieldInterface {
			field.DeleteParam(key)
			return field
		},
		"f_addCss": func(key, value string, field FieldInterface) FieldInterface {
			field.AddCSS(key, value)
			return field
		},
		"f_removeCss": func(key string, field FieldInterface) FieldInterface {
			field.RemoveCSS(key)
			return field
		},
		"f_setStyle": func(style string, field FieldInterface) FieldInterface {
			field.SetStyle(style)
			return field
		},
		"f_setLabel": func(label string, field FieldInterface) FieldInterface {
			field.SetLabel(label)
			return field
		},
		"f_addLabelClass": func(class string, field FieldInterface) FieldInterface {
			field.AddLabelClass(class)
			return field
		},
		"f_removeLabelClass": func(class string, field FieldInterface) FieldInterface {
			field.RemoveLabelClass(class)
			return field
		},
		"f_setValue": func(value interface{}, field FieldInterface) FieldInterface {
			field.SetValue(value)
			return field
		},
		"f_setText": func(text string, field FieldInterface) FieldInterface {
			field.SetText(text)
			return field
		},
		"f_disabled": func(field FieldInterface) FieldInterface {
			field.Disabled()
			return field
		},
		"f_enabled": func(field FieldInterface) FieldInterface {
			field.Enabled()
			return field
		},
		"f_setEditMode": func(mode bool, field FieldInterface) FieldInterface {
			if mode {
				field.DeleteParam("readonly")
			} else {
				field.SetParam("readonly", "readonly")
			}
			return field
		},
		"f_helpText": func(text string, field FieldInterface) FieldInterface {
			field.SetHelptext(text)
			return field
		},
		"f_addError": func(err string, field FieldInterface) FieldInterface {
			field.AddError(err)
			return field
		},
		"f_addData": func(key, value string, field FieldInterface) FieldInterface {
			field.AddData(key, value)
			return field
		},
		"f_multiple": func(field FieldInterface) FieldInterface {
			field.MultipleChoice()
			return field
		},
		"f_selected": func(options []string, field FieldInterface) FieldInterface {
			field.AddSelected(options...)
			return field
		},
		"render": func(field FieldInterface) template.HTML {
			return field.Render()
		},
		//"f_addError": func(err string, field FieldInterface) FieldInterface {
		//	field.AddError(err)
		//	return field
		//},

		// SingleChoice() FieldInterface
		// RemoveSelected(opt string) FieldInterface
		// SetSelectChoices(choices map[string][]InputChoice) FieldInterface
		// SetRadioChoices(choices []InputChoice) FieldInterface
	}
)

func toStringArray(value interface{}, defValue []string) []string {
	if value == nil {
		return defValue
	}
	switch vv := value.(type) {
	case []int:
		results := make([]string, 0, len(vv))
		for _, v := range vv {
			results = append(results, strconv.FormatInt(int64(v), 10))
		}
		return results
	case []int64:
		results := make([]string, 0, len(vv))
		for _, v := range vv {
			results = append(results, strconv.FormatInt(v, 10))
		}
		return results
	case []string:
		return vv
	case []interface{}:
		results := make([]string, 0, len(vv))
		for _, v := range vv {
			results = append(results, fmt.Sprint(v))
		}
		return results
	default:
		rValue := reflect.ValueOf(value)
		if rValue.Kind() == reflect.Slice {
			results := make([]string, 0, rValue.Len())
			for idx := 0; idx < rValue.Len(); idx++ {
				results = append(results, fmt.Sprint(rValue.Index(idx).Interface()))
			}
			return results
		}
	}

	return defValue
}
