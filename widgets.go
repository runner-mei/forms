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
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	rice "github.com/GeertJohan/go.rice"
)

var MapToString = func(v interface{}) string {
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
	w.template.ExecuteTemplate(buf, "main", data)
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

var g_id int32 = 0

var defaultFuncs = template.FuncMap{
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
		v := atomic.AddInt32(&g_id, 1)
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
}

func toTime(v interface{}) (time.Time, bool) {
	if t, ok := v.(time.Time); ok {
		return t, true
	}

	s, ok := v.(string)
	if !ok {
		return time.Time{}, false
	}
	for _, layout := range []string{
		"2006-01-02 15:04:05.999999999 -0700 MST",
		"2006-01-02 15:04:05 -0700 MST",
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02",
		"2006-01-02 15:04:05"} {
		if m, e := time.Parse(layout, s); nil == e {
			return m, true
		}
	}

	return time.Time{}, false
}
