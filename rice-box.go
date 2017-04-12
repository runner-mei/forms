package forms

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `baseform.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("<form{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\r\n\t{{ range .fields}}{{ .Render}}{{end}}\r\n</form>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/button.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{ define \"main\"}}<button type=\"{{.type}}\" name=\"{{.name}}\" class='btn {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}{{if eq .type \"submit\"}}btn-default{{end}}'{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</button>{{end}}"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/date.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/datetime.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/time.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/generic.tmpl`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{define \"generic\"}}<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n{{ if .label }}<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\r\n<input type=\"{{.type}}\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\r\n{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\r\n</div>\r\n{{end}}"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/hidden.html`,
		FileModTime: time.Unix(1491984455, 0),
		Content:     string("{{- define \"main\" -}}\r\n<input type=\"hidden\" name=\"{{.name}}\" class=\"{{range .classes}}{{.}} {{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\r\n{{- end -}}"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/input.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/mapinput.html`,
		FileModTime: time.Unix(1490170244, 0),
		Content:     string("{{ define \"main\"}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{ if .label }}\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\r\n  {{end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\r\n  {{if or .helptext .errors }}\r\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n    {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}\r\n  </span>\r\n  {{end}}\r\n  </div>\r\n</div>\r\n{{end}}"),
	}
	filee := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/number.html`,
		FileModTime: time.Unix(1490759833, 0),
		Content:     string("{{ define \"main\"}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{if .label }}\r\n  <label\r\n      class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{if .id}}\r\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\r\n  {{end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{if .id}} id=\"{{.id}}\" {{end}}\r\n           {{if .params}}\r\n           {{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{end}}\r\n           {{if .css}}\r\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{end}}\r\n           {{range $v :=.tags}} {{$v}}\r\n           {{end}}\r\n           {{if .value}}\r\n           value=\"{{.value}}\"\r\n           {{end}}>\r\n    {{if or .helptext .errors }}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}\r\n              </ul>{{end}}\r\n    </span>\r\n    {{end}}\r\n  </div>\r\n</div>\r\n{{end}}"),
	}
	filef := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/range.html`,
		FileModTime: time.Unix(1490759833, 0),
		Content:     string("{{ define \"main\"}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{if .label }}\r\n  <label\r\n      class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{if .id}}\r\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\r\n  {{end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{if .id}} id=\"{{.id}}\" {{end}}\r\n           {{if .params}}\r\n           {{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\r\n           {{end}}\r\n           {{if .css}}\r\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{end}}\r\n           {{range $v :=.tags}} {{$v}}\r\n           {{end}}\r\n           {{if .value}}\r\n           value=\"{{.value}}\"\r\n           {{end}}>\r\n    {{if or .helptext .errors }}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}\r\n              </ul>{{end}}\r\n    </span>\r\n    {{end}}\r\n  </div>\r\n</div>\r\n{{end}}"),
	}
	fileh := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/checkbox.html`,
		FileModTime: time.Unix(1490683458, 0),
		Content:     string("{{ define \"main\"}}\r\n{{ $p := . }}\r\n<div class=\"form-group\">\r\n\t<div class=\"col-lg-offset-{{default .labelWidth 2}} col-lg-{{default .controlWidth 9}}\">\r\n\t\t<div class=\"checkbox{{if .errors}} has-error{{end}}\">\r\n\t\t\t<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\">\r\n\t\t\t\t<input type=\"checkbox\" name=\"{{.name}}\"{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>\r\n\t\t\t\t{{.label}}\r\n\t\t\t</label>\r\n\t\t\t{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\r\n\t\t\t{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\r\n\t\t</div>\r\n\t</div>\r\n</div>\r\n{{end}}\r\n"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/radiobutton.html`,
		FileModTime: time.Unix(1490759833, 0),
		Content:     string("{{ define \"main\"}}{{ $p := . }}\r\n<div class=\"form-group\">\r\n{{if .label}}<label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">{{.label}}</label>{{end}}\r\n<!-- <div class=\"radio\"> -->\r\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    {{ range .choices }}\r\n    <label class=\"control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">\r\n      <input type=\"radio\" name=\"{{$p.name}}\"{{ if $p.classes }} class=\"{{range $p.classes}}{{.}} {{end}}\"{{end}} value=\"{{.Value}}\" id=\"{{$p.id}}\"{{if $p.params}}{{range $k2, $v2 := $p.params}} {{$k2}}=\"{{$v2}}\"{{end}}{{end}}{{if $p.css}} style=\"{{range $k2, $v2 := .css}}{{$k2}}: {{$v2}}; {{end}}\"{{end}}{{ if eq $p.value .Value}} checked{{end}}{{range $v2 := $p.tags}} {{$v2}}{{end}}>\r\n      {{.Label}}\r\n    </label>\r\n    {{end}}\r\n\t</div>\r\n<!-- </div> -->\r\n\r\n</div>{{end}}\r\n"),
	}
	filej := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/select.html`,
		FileModTime: time.Unix(1491983985, 0),
		Content:     string("{{ define \"main\"}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if .label -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n      {{- if .id}} for=\"{{.id}}\" {{end -}}>\r\n    {{.label}}\r\n  </label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <select name=\"{{.name}}\" class=\"form-control {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"\r\n            {{- if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end -}}\r\n            {{- if .css}}\r\n            style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n            {{- end -}}\r\n            {{- range $v :=.tags}} {{$v}}{{end}}>\r\n      {{- $p := . -}}\r\n      {{- range $k, $v := .choices -}}\r\n      {{- if $k -}}\r\n        <optgroup label=\"{{$k}}\">{{end -}}\r\n      {{- range $v -}}\r\n        <option value=\"{{.Value}}\"\r\n                {{- if strIn \"multiple\" $p.tags -}}\r\n                  {{- $id :=.Value -}}\r\n                  {{- range $k2, $p2 :=$p.multValues -}}\r\n                  {{- if eq $k2 $id}}selected{{end -}}\r\n                  {{- end -}}\r\n                {{- else -}}\r\n                  {{- if eq $p.value .Value}} selected{{end -}}\r\n                {{- end}}>{{raw .Label -}}\r\n        </option>\r\n      {{- end -}}\r\n      {{- if $k -}}\r\n        </optgroup>\r\n      {{- end -}}\r\n      {{- end -}}\r\n    </select>\r\n    {{- if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end -}}\r\n  </div>\r\n</div>\r\n{{- end}}\r\n"),
	}
	filek := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/static.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("{{define \"main\"}}<div class=\"form-group\">\r\n{{ if .label }}<label{{ if .labelClasses }} class=\"{{range .labelClasses}}{{.}} {{end}}\"{{end}}{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\r\n<p name=\"{{.name}}\" class=\"form-control-static {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</p>\r\n</div>{{end}}"),
	}
	filem := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/passwordinput.html`,
		FileModTime: time.Unix(1491984761, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if .label -}}\r\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\" {{if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"password\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}{{if .value}} value=\"{{.value}}\" {{end}}>\r\n    {{- if or .helptext .errors -}}\r\n    <span class=\"help-block\">\r\n      {{- if .helptext}}{{ .helptext }}{{end -}}\r\n      {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\r\n    </span>\r\n    {{- end -}}\r\n  </div>\r\n</div>{{end -}}"),
	}
	filen := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textareainput.html`,
		FileModTime: time.Unix(1491984751, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if .label -}}\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{- if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\r\n  {{- end}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\r\n  {{- if or .helptext .errors -}}\r\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\r\n  </span>\r\n  {{- end -}}\r\n  </div>\r\n</div>\r\n{{- end -}}"),
	}
	fileo := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textinput.html`,
		FileModTime: time.Unix(1491986196, 0),
		Content:     string("{{- define \"main\" -}}\r\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\r\n  {{- if .label -}}\r\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\r\n  {{- if .id -}}\r\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\r\n  {{- end -}}\r\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\r\n    <input type=\"text\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\r\n           {{- if .id}} id=\"{{.id}}\" {{end -}}\r\n           {{- if .params -}}\r\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}}\r\n           {{- end -}}\r\n           {{- if .css -}}\r\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\r\n           {{- end -}}\r\n           {{- range $v :=.tags}} {{$v -}}\r\n           {{- end -}}\r\n           {{- if .value -}}\r\n           value=\"{{.value}}\"\r\n           {{end}}>\r\n    {{- if or .helptext .errors -}}\r\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\r\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end -}}\r\n              </ul>{{end -}}\r\n    </span>\r\n    {{end -}}\r\n  </div>\r\n</div>\r\n{{- end -}}"),
	}
	filep := &embedded.EmbeddedFile{
		Filename:    `bootstrapform.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("<form role=\"form\"{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\r\n\t{{ range .fields}}{{ .Render}}{{end}}\r\n</form>"),
	}
	fileq := &embedded.EmbeddedFile{
		Filename:    `fieldset.html`,
		FileModTime: time.Unix(1490170239, 0),
		Content:     string("<fieldset{{if .classes }} class=\"{{range $v := .classes}} {{$v}}{{end}}\"{{end}}{{ if .tags}}{{range $v := .tags}} {{$v}}{{end}}{{end}}>\r\n\t{{range .fields}}{{ .Render }}{{end}}\r\n</fieldset>\r\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1490170239, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // baseform.html
			filep, // bootstrapform.html
			fileq, // fieldset.html

		},
	}
	dir3 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3`,
		DirModTime: time.Unix(1491984455, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // bootstrap3/button.html
			file9, // bootstrap3/generic.tmpl
			filea, // bootstrap3/hidden.html
			fileb, // bootstrap3/input.html
			filec, // bootstrap3/mapinput.html
			filek, // bootstrap3/static.html

		},
	}
	dir5 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/datetime`,
		DirModTime: time.Unix(1490170239, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file6, // bootstrap3/datetime/date.html
			file7, // bootstrap3/datetime/datetime.html
			file8, // bootstrap3/datetime/time.html

		},
	}
	dird := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/number`,
		DirModTime: time.Unix(1490759833, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filee, // bootstrap3/number/number.html
			filef, // bootstrap3/number/range.html

		},
	}
	dirg := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/options`,
		DirModTime: time.Unix(1491983985, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			fileh, // bootstrap3/options/checkbox.html
			filei, // bootstrap3/options/radiobutton.html
			filej, // bootstrap3/options/select.html

		},
	}
	dirl := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/text`,
		DirModTime: time.Unix(1491986196, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filem, // bootstrap3/text/passwordinput.html
			filen, // bootstrap3/text/textareainput.html
			fileo, // bootstrap3/text/textinput.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir3, // bootstrap3

	}
	dir3.ChildDirs = []*embedded.EmbeddedDir{
		dir5, // bootstrap3/datetime
		dird, // bootstrap3/number
		dirg, // bootstrap3/options
		dirl, // bootstrap3/text

	}
	dir5.ChildDirs = []*embedded.EmbeddedDir{}
	dird.ChildDirs = []*embedded.EmbeddedDir{}
	dirg.ChildDirs = []*embedded.EmbeddedDir{}
	dirl.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1490170239, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":                    dir1,
			"bootstrap3":          dir3,
			"bootstrap3/datetime": dir5,
			"bootstrap3/number":   dird,
			"bootstrap3/options":  dirg,
			"bootstrap3/text":     dirl,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"baseform.html":                       file2,
			"bootstrap3/button.html":              file4,
			"bootstrap3/datetime/date.html":       file6,
			"bootstrap3/datetime/datetime.html":   file7,
			"bootstrap3/datetime/time.html":       file8,
			"bootstrap3/generic.tmpl":             file9,
			"bootstrap3/hidden.html":              filea,
			"bootstrap3/input.html":               fileb,
			"bootstrap3/mapinput.html":            filec,
			"bootstrap3/number/number.html":       filee,
			"bootstrap3/number/range.html":        filef,
			"bootstrap3/options/checkbox.html":    fileh,
			"bootstrap3/options/radiobutton.html": filei,
			"bootstrap3/options/select.html":      filej,
			"bootstrap3/static.html":              filek,
			"bootstrap3/text/passwordinput.html":  filem,
			"bootstrap3/text/textareainput.html":  filen,
			"bootstrap3/text/textinput.html":      fileo,
			"bootstrapform.html":                  filep,
			"fieldset.html":                       fileq,
		},
	})
}
