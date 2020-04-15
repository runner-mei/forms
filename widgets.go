// This package contains the base logic for the creation and rendering of field widgets. Base widgets are defined for most input fields,
// both in classic and Bootstrap3 style; custom widgets can be defined and associated to a field, provided that they implement the
// WidgetInterface interface.
package forms

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	rice "github.com/GeertJohan/go.rice"
)

var MapToString = func(v interface{}) string {
	if v == nil {
		return ""
	}
	m, ok := v.(map[string]interface{})
	if !ok {
		panic(fmt.Errorf("MapToString: value is not map - %T `%v`", v, v))
	}
	var buf bytes.Buffer
	buf.WriteString("# 一行为一条记录，以等号分隔键和值\r\n")
	buf.WriteString("# 以 # 开始的行及空行将忽略\r\n")
	for k, v := range m {
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(fmt.Sprint(v))
		buf.WriteString("\r\n")
	}
	return buf.String()
}
var rootPath string
var templateFuncs template.FuncMap
var templateBox *rice.Box
var isDevMode bool

func Init(devMode bool, root string, funcs template.FuncMap) {
	isDevMode = devMode
	rootPath = root
	templateFuncs = funcs
	templateBox = rice.MustFindBox("templates")
}

// WidgetInterface defines the requirements for custom widgets.
type Widget interface {
	Render(data interface{}) string
}

// Simple widget object that gets executed at render time.
type widget struct {
	template *template.Template
}

var bytesCache sync.Pool
var templates atomic.Value

func init() {
	bytesCache.New = func() interface{} {
		return make([]byte, 0, 2*1024)
	}
}

// Render executes the internal template and returns the result as a template.HTML object.
func (w *widget) Render(data interface{}) string {
	bs := bytesCache.Get().([]byte)
	buf := bytes.NewBuffer(bs)
	err := w.template.ExecuteTemplate(buf, "main", data)
	if err != nil {
		bytesCache.Put(bs)
		panic(err)
	}
	s := buf.String()
	bytesCache.Put(bs)
	return s
}

func putTemplateToCache(name string, templ *template.Template) {
	o := templates.Load()
	if o == nil {
		templates.Store(map[string]*template.Template{name: templ})
		return
	}
	m, _ := o.(map[string]*template.Template)
	if m == nil {
		templates.Store(map[string]*template.Template{name: templ})
		return
	}
	newValue := map[string]*template.Template{name: templ}
	for k, v := range m {
		newValue[k] = v
	}
	templates.Store(newValue)
}

func getTemplateFromCache(name string) *template.Template {
	o := templates.Load()
	if o == nil {
		return nil
	}
	m, _ := o.(map[string]*template.Template)
	if m == nil {
		return nil
	}
	return m[name]
}

// BaseWidget creates a Widget based on style and inpuType parameters, both defined in the common package.
func BaseWidget(style, inputType string) Widget {
	if isDevMode {
		templ := loadTemplate(style, inputType)
		return &widget{template: templ}
	}
	name := style + "/" + inputType
	templ := getTemplateFromCache(name)
	if templ == nil {
		templ = loadTemplate(style, inputType)
		putTemplateToCache(name, templ)
	}
	return &widget{template: templ}
}

func loadTemplate(style, inputType string) *template.Template {
	var widgetFilename = style + "/generic.tmpl"
	switch inputType {
	case BUTTON:
		widgetFilename = style + "/button.html"
	case CHECKBOX:
		widgetFilename = style + "/options/checkbox.html"
	case TEXTAREA:
		widgetFilename = style + "/text/textareainput.html"
	case SELECT:
		widgetFilename = style + "/options/select.html"
	case MULI_SOURCE_SELECT:
		widgetFilename = style + "/options/mult_source_select.html"
	case PASSWORD:
		widgetFilename = style + "/text/passwordinput.html"
	case RADIO:
		widgetFilename = style + "/options/radiobutton.html"
	case TEXT:
		widgetFilename = style + "/text/textinput.html"
	case RANGE:
		widgetFilename = style + "/number/range.html"
	case NUMBER:
		widgetFilename = style + "/number/number.html"
	case RESET:
		widgetFilename = style + "/button.html"
	case SUBMIT:
		widgetFilename = style + "/button.html"
	case DATE:
		widgetFilename = style + "/datetime/date.html"
	case DATETIME:
		widgetFilename = style + "/datetime/datetime.html"
	case TIME:
		widgetFilename = style + "/datetime/time.html"
	case DATETIME_LOCAL:
		widgetFilename = style + "/datetime/datetime.html"
	case STATIC:
		widgetFilename = style + "/static.html"
	case CRON:
		widgetFilename = style + "/cron.html"
	case HIDDEN:
		widgetFilename = style + "/hidden.html"
	case MAP:
		widgetFilename = style + "/map.html"
	case SEARCH,
		TEL,
		URL,
		WEEK,
		COLOR,
		EMAIL,
		FILE,
		IMAGE,
		MONTH:
		widgetFilename = style + "/input.html"
	case "unmodifiable":
		widgetFilename = style + "/unmodifiable.html"
	default:
		widgetFilename = style + "/input.html"
	}

	if _, err := os.Stat(filepath.Join("templates", widgetFilename)); err == nil {
		return mustLoadTemplate(style, filepath.Join("templates", widgetFilename))
	}

	if _, err := os.Stat(filepath.Join(rootPath, "templates", widgetFilename)); err == nil {
		return mustLoadTemplate(style, filepath.Join(rootPath, "templates", widgetFilename))
	}

	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		styledURL := path.Join(p, "src/github.com/three-plus-three/forms", "templates", widgetFilename)
		if _, err := os.Stat(styledURL); err == nil {
			return mustLoadTemplate(style, styledURL)
		}
	}

	txt, err := templateBox.String(widgetFilename)
	if err != nil {
		panic(errors.New("load template(" + widgetFilename + ") fail, " + err.Error()))
	}
	templ, err := template.New(style).Funcs(defaultFuncs).Funcs(templateFuncs).Parse(txt)
	if err != nil {
		panic(errors.New("load template(" + widgetFilename + ") from rice-box fail, " + err.Error()))
	}
	return templ
}

func mustLoadTemplate(style, filename string) *template.Template {
	templ, err := template.New(style).Funcs(defaultFuncs).Funcs(templateFuncs).ParseFiles(filename)
	if err != nil {
		panic(errors.New("load template(" + filename + ") fail, " + err.Error()))
	}
	return templ
}

var gID int32

var defaultFuncs = template.FuncMap{
	"toOptionBoolean": func(v interface{}) bool {
		if v == nil {
			return false
		}
		if b, ok := v.(bool); ok {
			return b
		}
		if s, ok := v.(string); ok {
			if s == "true" || s == "on" || s == "yes" ||
				s == "True" || s == "On" || s == "Yes" ||
				s == "TRUE" || s == "ON" || s == "YES" {
				return true
			}
			return false
		}
		panic(fmt.Errorf("unknown type - %T %v", v, v))
	},
	"set": func(renderArgs map[string]interface{}, key string, value interface{}) template.JS {
		renderArgs[key] = value
		return template.JS("")
	},
	"append": func(renderArgs map[string]interface{}, key string, value interface{}) template.JS {
		if renderArgs[key] == nil {
			renderArgs[key] = []interface{}{value}
		} else {
			renderArgs[key] = append(renderArgs[key].([]interface{}), value)
		}
		return template.JS("")
	},

	"unique": func(renderArgs map[string]interface{}, key string) template.JS {
		o := renderArgs[key]
		if o == nil {
			return template.JS("")
		}
		values, ok := o.([]interface{})
		if !ok || len(values) <= 1 {
			return template.JS("")
		}
		newValues := make([]interface{}, 0, len(values))
		for idx, value := range values {
			if idx == 0 {
				newValues = append(newValues, value)
				continue
			}

			found := false
			for _, pre := range values[:idx] {
				if pre == value {
					found = true
					break
				}
			}

			if !found {
				newValues = append(newValues, value)
			}
		}
		renderArgs[key] = newValues
		return template.JS("")
	},

	"default": func(value, defvalue interface{}) interface{} {
		if nil == value {
			return defvalue
		}
		if s, ok := value.(string); ok && "" == s {
			return defvalue
		}
		return value
	},
	"form_date": func(value interface{}) string {
		t, ok := toTime(value)
		if !ok {
			return fmt.Sprint(value)
		}
		if t.IsZero() {
			return ""
		}
		return t.Format("2006-01-02")
	},
	"form_time": func(value interface{}) string {
		t, ok := toTime(value)
		if !ok {
			return fmt.Sprint(value)
		}
		if t.IsZero() {
			return ""
		}
		return t.Format("15:04")
	},
	"form_date_and_time": func(value interface{}) string {
		t, ok := toTime(value)
		if !ok {
			return fmt.Sprint(value)
		}
		if t.IsZero() {
			return ""
		}
		return t.Format("2006-01-02 15:04")
	},
	"form_datetime": func(value interface{}) string {
		t, ok := toTime(value)
		if !ok {
			return fmt.Sprint(value)
		}
		if t.IsZero() {
			return ""
		}
		return t.Format("2006-01-02 15:04:05")
	},
	"generateID": func() string {
		v := atomic.AddInt32(&gID, 1)
		return "widget_" + strconv.FormatInt(int64(v), 10)
	},
	"strIn": func(value string, values []string) bool {
		for _, v := range values {
			if v == value {
				return true
			}
		}
		return false
	},
	// Replaces newlines with <br>
	"nl2br": func(text string) template.HTML {
		return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
	},

	// Skips sanitation on the parameter.  Do not use with dynamic data.
	"raw": func(text string) template.HTML {
		return template.HTML(text)
	},

	// Skips sanitation on the parameter.  Do not use with dynamic data.
	"toInt": func(text interface{}) int {
		i, err := strconv.Atoi(fmt.Sprint(text))
		if err != nil {
			panic(err)
		}
		return i
	},
	"toString": func(v interface{}) string {
		if v == nil {
			return ""
		}
		return fmt.Sprint(v)
	},

	// Skips sanitation on the parameter.  Do not use with dynamic data.
	"mul": func(a, b int) int {
		return a * b
	},

	// Skips sanitation on the parameter.  Do not use with dynamic data.
	"add": func(a, b int) int {
		return a + b
	},
	"readOptValue": readOptValue,
	"readOptLabel": readOptLabel,
}

func toBoolean(v interface{}) bool {
	if v == nil {
		return false
	}
	if b, ok := v.(bool); ok {
		return b
	}
	if s, ok := v.(string); ok {
		if s == "true" || s == "on" || s == "yes" ||
			s == "True" || s == "On" || s == "Yes" ||
			s == "TRUE" || s == "ON" || s == "YES" {
			return true
		}
	}
	return false
}

func toTime(v interface{}) (time.Time, bool) {
	if t, ok := v.(time.Time); ok {
		return t, true
	}

	if t, ok := v.(*time.Time); ok {
		return *t, true
	}

	s, ok := v.(string)
	if !ok {
		return time.Time{}, false
	}
	if s == "" {
		return time.Time{}, true
	}
	for _, layout := range []string{
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02 15:04:05.999999999 -0700 -0700",
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700 -0700",
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05"} {
		if m, e := time.Parse(layout, s); nil == e {
			return m, true
		}
	}

	return time.Time{}, false
}

func readOptValue(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	switch value := v.(type) {
	case *InputChoice:
		return value.Value
	case InputChoice:
		return value.Value
	case [2]string:
		return value[0]
	case string:
		return value
	case map[string]interface{}:
		id := value["value"]
		if id == nil {
			id = value["id"]
		}
		return id
	case map[string]string:
		id, ok := value["value"]
		if !ok {
			id = value["id"]
		}
		return id
	case interface {
		OptionValue() interface{}
	}:
		return value.OptionValue()
	case interface {
		Value() interface{}
	}:
		return value.Value()
	default:
		rv := reflect.ValueOf(v)
		rValue := rv.FieldByName("OptionValue")
		if rValue.IsValid() {
			return rValue.Interface()
		}
		rValue = rv.FieldByName("Value")
		if rValue.IsValid() {
			return rValue.Interface()
		}
		panic(fmt.Errorf("read option value fail, unknown type is %T", v))
	}
}

func readOptLabel(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	switch value := v.(type) {
	case *InputChoice:
		return value.Label
	case InputChoice:
		return value.Label
	case *HierarchyChoice:
		return value.Label
	case HierarchyChoice:
		return value.Label
	case [2]string:
		return value[1]
	case string:
		return value
	case map[string]interface{}:
		text := value["label"]
		if text == nil {
			text = value["text"]
		}
		return text
	case map[string]string:
		text, ok := value["label"]
		if !ok {
			text = value["text"]
		}
		return text
	case interface {
		OptionLabel() string
	}:
		return value.OptionLabel()
	case interface {
		Label() string
	}:
		return value.Label()
	default:
		rv := reflect.ValueOf(v)
		rValue := rv.FieldByName("OptionLabel")
		if rValue.IsValid() {
			return rValue.Interface()
		}
		rValue = rv.FieldByName("Label")
		if rValue.IsValid() {
			return rValue.Interface()
		}
		panic(fmt.Errorf("read option label fail, unknown type is %T", v))
	}
}

func isOptionSet(v interface{}, isSlice bool) (bool, bool) {
	if v == nil {
		return false, false
	}

	switch value := v.(type) {
	case []*InputChoice:
		return true, false
	case []InputChoice:
		return true, false
	case []*HierarchyChoice:
		return true, true
	case []HierarchyChoice:
		return true, true
	case [][2]string:
		return true, false
	case []string:
		return true, false
	case []interface{}:
		if len(value) == 0 {
			return true, false
		}
		for _, v := range value {
			return isOptionSet(v, false)
		}
		return false, false
	case map[string]interface{}:
		if len(value) == 0 {
			return true, false
		}
		for _, v := range value {
			return isOptionSet(v, false)
		}
		return false, false
	case []map[string]interface{}:
		if len(value) == 0 {
			return true, false
		}
		object := value[0]
		_, ok := object["label"]
		if ok {
			_, ok = object["value"]
			if ok {
				return true, false
			}
			_, ok = object["id"]
			if ok {
				return true, false
			}
			return false, false
		}
		_, ok = object["text"]
		if ok {
			_, ok = object["value"]
			if ok {
				return true, false
			}
			_, ok = object["id"]
			if ok {
				return true, false
			}
			return true, false
		}
		return false, false
	case []map[string]string:
		if len(value) == 0 {
			return true, false
		}
		object := value[0]
		_, ok := object["label"]
		if ok {
			_, ok = object["value"]
			if ok {
				return true, false
			}
			_, ok = object["id"]
			if ok {
				return true, false
			}
			return false, false
		}
		_, ok = object["text"]
		if ok {
			_, ok = object["value"]
			if ok {
				return true, false
			}
			_, ok = object["id"]
			if ok {
				return true, false
			}
			return true, false
		}
		return false, false
	default:
		var rv reflect.Value
		if isSlice {
			rv = reflect.ValueOf(v)
			if rv.Kind() != reflect.Slice {
				return false, false
			}
			if rv.Len() == 0 {
				return isOptionSetByType(rv.Type(), true)
			}
			rv = rv.Index(0)
			if !rv.IsValid() {
				return false, false
			}
		}

		switch rv.Interface().(type) {
		case *InputChoice:
			return true, false
		case InputChoice:
			return true, false
		case [2]string:
			return true, false
		case interface {
			OptionLabel() string
			OptionValue() interface{}
		}:
			return true, false
		case interface {
			Label() string
			Value() interface{}
		}:
			return true, false
		}

		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		if rv.Kind() != reflect.Struct {
			return false, false
		}

		rValue := rv.FieldByName("OptionLabel")
		if rValue.IsValid() {
			rValue = rv.FieldByName("OptionValue")
			if rValue.IsValid() {
				return true, false
			}

			rValue = rv.FieldByName("Children")
			if rValue.IsValid() {
				valid, _ := isOptionSet(rValue.Interface(), true)
				return valid, true
			}

			return false, false
		}
		rValue = rv.FieldByName("Label")
		if rValue.IsValid() {
			rValue = rv.FieldByName("Value")
			if rValue.IsValid() {
				return true, false
			}

			rValue = rv.FieldByName("Children")
			if rValue.IsValid() {
				valid, _ := isOptionSet(rValue.Interface(), true)
				return valid, true
			}
			return false, false
		}
		return false, false
	}
}

var InputChoiceType = reflect.TypeOf(InputChoice{})
var InputChoicePtrType = reflect.TypeOf(&InputChoice{})

var HierarchyChoiceType = reflect.TypeOf(HierarchyChoice{})
var HierarchyChoicePtrType = reflect.TypeOf(&HierarchyChoice{})
var StrArrayPtrType = reflect.TypeOf([2]string{})

func isOptionSetByType(t reflect.Type, isSlice bool) (bool, bool) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if isSlice {
		if t.Kind() != reflect.Slice {
			return false, false
		}
		t = t.Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
	}

	switch t {
	case InputChoiceType:
		return true, false
	case InputChoicePtrType:
		return true, false
	case HierarchyChoiceType:
		return true, true
	case HierarchyChoicePtrType:
		return true, true
	case StrArrayPtrType:
		return true, false
		// case interface {
		// 	OptionLabel() string
		// 	OptionValue() interface{}
		// }:
		// 	return true, false
		// case interface {
		// 	Label() string
		// 	Value() interface{}
		// }:
		// 	return true, false
	}

	if t.Kind() != reflect.Struct {
		return false, false
	}

	if hasFunc(t, "OptionLabel") {
		if hasFunc(t, "OptionValue") {
			return true, false
		}
		return false, false
	}

	if hasFunc(t, "Label") {
		if hasFunc(t, "Value") {
			return true, false
		}
		return false, false
	}

	if ok, _ := hasField(t, "OptionLabel"); ok {
		if ok, _ := hasField(t, "OptionValue"); ok {
			return true, false
		}
		ok, vt := hasField(t, "Children")
		if ok {
			valid, _ := isOptionSetByType(vt, true)
			return valid, true
		}
		return false, false
	}

	if ok, _ := hasField(t, "Label"); ok {
		if ok, _ := hasField(t, "Value"); ok {
			return true, false
		}

		ok, vt := hasField(t, "Children")
		if ok {
			valid, _ := isOptionSetByType(vt, true)
			return valid, true
		}
		return false, false
	}
	return false, false
}

func hasFunc(t reflect.Type, name string) bool {
	_, ok := t.MethodByName(name)
	if ok {
		return true
	}
	return false
}

func hasField(t reflect.Type, name string) (bool, reflect.Type) {
	f, ok := t.FieldByName(name)
	if ok {
		return true, f.Type
	}
	return false, nil
}
