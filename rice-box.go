package forms

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `baseform.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("<form{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\n\t{{ range .fields}}{{ .Render}}{{end}}\n</form>"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/button.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}<button type=\"{{.type}}\" name=\"{{.name}}\" class='btn {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}{{if eq .type \"submit\"}}btn-default{{end}}'{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</button>{{end}}"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/date.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/datetime.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/time.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/generic.tmpl`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"generic\"}}<div class=\"form-group{{if .errors}} has-error{{end}}\">\n{{ if .label }}<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\n<input type=\"{{.type}}\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\n{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\n</div>\n{{end}}"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/hidden.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"main\"}}\n<input type=\"hidden\" name=\"{{.name}}\" class=\"{{range .classes}}{{.}} {{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\n{{end}}"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/input.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filed := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/number.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filee := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/range.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	fileg := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/checkbox.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}<div class=\"checkbox{{if .errors}} has-error{{end}}\">\n<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\">\n\t<input type=\"checkbox\" name=\"{{.name}}\"{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>\n\t{{.label}}\n</label>\n{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\n</div>\n{{end}}\n"),
	}
	fileh := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/radiobutton.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ $p := . }}\n<div class=\"form-group\">\n{{if .label}}<label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">{{.label}}</label>{{end}}\n<!-- <div class=\"radio\"> -->\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\n    {{ range .choices }}\n    <label class=\"control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">\n      <input type=\"radio\" name=\"{{$p.name}}\"{{ if $p.classes }} class=\"{{range $p.classes}}{{.}} {{end}}\"{{end}} value=\"{{.ID}}\" id=\"{{$p.id}}\"{{if $p.params}}{{range $k2, $v2 := $p.params}} {{$k2}}=\"{{$v2}}\"{{end}}{{end}}{{if $p.css}} style=\"{{range $k2, $v2 := .css}}{{$k2}}: {{$v2}}; {{end}}\"{{end}}{{ if eq $p.value .ID}} checked{{end}}{{range $v2 := $p.tags}} {{$v2}}{{end}}>\n      {{.Val}}\n    </label>\n    {{end}}\n\t</div>\n<!-- </div> -->\n\n</div>{{end}}\n"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/select.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{if .label}}\n  <label\n      class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n      {{if .id}} for=\"{{.id}}\" {{end}}>\n    {{.label}}\n  </label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <select name=\"{{.name}}\" class=\"form-control {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"\n            {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}\n            {{if .css}}\n            style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\n            {{end}}\n            {{range $v :=.tags}} {{$v}}{{end}}>\n      {{ $p := .}}\n      {{range $k, $v := .choices}}\n      {{ if $k }}\n        <optgroup label=\"{{$k}}\">{{end}}\n      {{range $v}}\n        <option value=\"{{.ID}}\"\n                {{if strIn \"multiple\" $p.tags }}\n                  {{$id :=.ID}}\n                  {{range $k2, $p2 :=$p.multValues}}\n                  {{if eq $k2 $id}}selected{{end}}\n                  {{end}}\n                {{else}}\n                  {{ if eq $p.value .ID}} selected{{end}}\n                {{end}}>{{.Val}}\n        </option>\n      {{end}}\n      {{if $k}}\n        </optgroup>\n      {{end}}\n      {{end}}\n    </select>\n    {{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n    {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\n  </div>\n</div>\n{{end}}\n"),
	}
	filej := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/static.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"main\"}}<div class=\"form-group\">\n{{ if .label }}<label{{ if .labelClasses }} class=\"{{range .labelClasses}}{{.}} {{end}}\"{{end}}{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\n<p name=\"{{.name}}\" class=\"form-control-static {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</p>\n</div>{{end}}"),
	}
	filel := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/passwordinput.html`,
		FileModTime: time.Unix(1489737066, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{if .label }}\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\" {{if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <input type=\"password\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}{{if .value}} value=\"{{.value}}\" {{end}}>\n    {{if or .helptext .errors }}\n    <span class=\"help-block\">\n      {{if .helptext}}{{ .helptext }}{{end}}\n      {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}\n    </span>\n    {{end}}\n  </div>\n</div>{{end}}"),
	}
	filem := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textareainput.html`,
		FileModTime: time.Unix(1489736870, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{ if .label }}\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\n  {{if or .helptext .errors }}\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n    {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}\n  </span>\n  {{end}}\n  </div>\n</div>\n{{end}}"),
	}
	filen := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textinput.html`,
		FileModTime: time.Unix(1489736939, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{if .label }}\n  <label\n      class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{if .id}}\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <input type=\"text\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\n           {{if .id}} id=\"{{.id}}\" {{end}}\n           {{if .params}}\n           {{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\n           {{end}}\n           {{if .css}}\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\n           {{end}}\n           {{range $v :=.tags}} {{$v}}\n           {{end}}\n           {{if .value}}\n           value=\"{{.value}}\"\n           {{end}}>\n    {{if or .helptext .errors }}\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}\n              </ul>{{end}}\n    </span>\n    {{end}}\n  </div>\n</div>\n{{end}}"),
	}
	fileo := &embedded.EmbeddedFile{
		Filename:    `bootstrapform.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("<form role=\"form\"{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\n\t{{ range .fields}}{{ .Render}}{{end}}\n</form>"),
	}
	filep := &embedded.EmbeddedFile{
		Filename:    `fieldset.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("<fieldset{{if .classes }} class=\"{{range $v := .classes}} {{$v}}{{end}}\"{{end}}{{ if .tags}}{{range $v := .tags}} {{$v}}{{end}}{{end}}>\n\t{{range .fields}}{{ .Render }}{{end}}\n</fieldset>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // baseform.html
			fileo, // bootstrapform.html
			filep, // fieldset.html

		},
	}
	dir3 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // bootstrap3/button.html
			file9, // bootstrap3/generic.tmpl
			filea, // bootstrap3/hidden.html
			fileb, // bootstrap3/input.html
			filej, // bootstrap3/static.html

		},
	}
	dir5 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/datetime`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file6, // bootstrap3/datetime/date.html
			file7, // bootstrap3/datetime/datetime.html
			file8, // bootstrap3/datetime/time.html

		},
	}
	dirc := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/number`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filed, // bootstrap3/number/number.html
			filee, // bootstrap3/number/range.html

		},
	}
	dirf := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/options`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			fileg, // bootstrap3/options/checkbox.html
			fileh, // bootstrap3/options/radiobutton.html
			filei, // bootstrap3/options/select.html

		},
	}
	dirk := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/text`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filel, // bootstrap3/text/passwordinput.html
			filem, // bootstrap3/text/textareainput.html
			filen, // bootstrap3/text/textinput.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir3, // bootstrap3

	}
	dir3.ChildDirs = []*embedded.EmbeddedDir{
		dir5, // bootstrap3/datetime
		dirc, // bootstrap3/number
		dirf, // bootstrap3/options
		dirk, // bootstrap3/text

	}
	dir5.ChildDirs = []*embedded.EmbeddedDir{}
	dirc.ChildDirs = []*embedded.EmbeddedDir{}
	dirf.ChildDirs = []*embedded.EmbeddedDir{}
	dirk.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1489735114, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":                    dir1,
			"bootstrap3":          dir3,
			"bootstrap3/datetime": dir5,
			"bootstrap3/number":   dirc,
			"bootstrap3/options":  dirf,
			"bootstrap3/text":     dirk,
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
			"bootstrap3/number/number.html":       filed,
			"bootstrap3/number/range.html":        filee,
			"bootstrap3/options/checkbox.html":    fileg,
			"bootstrap3/options/radiobutton.html": fileh,
			"bootstrap3/options/select.html":      filei,
			"bootstrap3/static.html":              filej,
			"bootstrap3/text/passwordinput.html":  filel,
			"bootstrap3/text/textareainput.html":  filem,
			"bootstrap3/text/textinput.html":      filen,
			"bootstrapform.html":                  fileo,
			"fieldset.html":                       filep,
		},
	})
}
