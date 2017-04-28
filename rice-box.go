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
	file5 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/cron.html`,
		FileModTime: time.Unix(1491557134, 0),
		Content:     string("{{ define \"main\"}}\n\n<script>\n    $(function () {\n        $('#cornExpression a').click(function (e) {\n            e.preventDefault();\n            $(this).tab('show')\n        });\n//        function offset(){\n//            var ele = $(\"#offset\");\n//            return ele.val()?\"_\"+ele.val():\"\"\n//        }\n        var cornExpression = [];\n        $(\"#month select\").on(\"change\", function () {\n            var month_monthDay = $(\"#month_monthDay\").val(),\n                    month_hour = $(\"#month_hour\").val(),\n                    month_minute = $(\"#month_minute\").val(),\n                    cornExpression = [\"0\", month_minute, month_hour, month_monthDay, \"*\", \"?\" ];\n            $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\n        });\n        $(\"#week select\").on(\"change\", function () {\n            var week_weekDay = $(\"#week_weekDay\").val(),\n                    week_hour = $(\"#week_hour\").val(),\n                    week_minute = $(\"#week_minute\").val(),\n                    cornExpression = [\"0\", week_minute, week_hour, \"?\", \"*\", week_weekDay ];\n            $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\n        });\n        $(\"#day select\").on(\"change\", function () {\n            var day_hour = $(\"#day_hour\").val(),\n                    day_minute = $(\"#day_minute\").val(),\n                    cornExpression = [\"0\", day_minute, day_hour, \"*\", \"*\", \"?\" ];\n            $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\n        });\n        $(\"#hour select\").on(\"change\", function () {\n            var hour_hour = $(\"#hour_hour\").val(),\n                    cornExpression = [\"0\", \"0\", \"0/\" + hour_hour, \"*\", \"*\", \"?\" ];\n            $(\"[name='{{.name}}']\").val(cornExpression.join(\" \"));\n        });\n        $(\"#offset\").on(\"change\", function () {\n            $(\".tab-content .active select\").trigger(\"change\")\n        });\n        $(\"#cornExpression\").on(\"change\", \"select\", function () {\n            $(\"[name='{{.name}}']\").trigger(\"change\")\n        });\n    })\n</script>\n<div class=\"control-group{{if .errors}} has-error{{end}}\">{{ if .label }}\n    <label class=\"control-label{{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\" {{if .id}}\n           for=\"{{.id}}\" {{end}}>{{.label}}</label>{{end}}\n    <div class=\"controls\">\n        <div class=\"list-group-item\">\n            <ul class=\"nav nav-tabs\" role=\"tablist\" id=\"cornExpression\">\n                <li role=\"presentation\" class=\"active\"><a href=\"#month\" style=\"color:#555\">每月</a></li>\n                <li role=\"presentation\"><a href=\"#week\" style=\"color:#555\">每周</a></li>\n                <li role=\"presentation\"><a href=\"#day\"  style=\"color:#555\">每日</a></li>\n                <li role=\"presentation\"><a href=\"#hour\" style=\"color:#555\">每时</a></li>\n            </ul>\n\n            <div class=\"tab-content\">\n                <div role=\"tabpanel\" class=\"tab-pane active\" id=\"month\">\n                    <span>每月</span>\n                    <select id=\"month_monthDay\">\n                        <option value=\"1\">1日</option>\n                        <option value=\"2\">2日</option>\n                        <option value=\"3\">3日</option>\n                        <option value=\"4\">4日</option>\n                        <option value=\"5\">5日</option>\n                        <option value=\"6\">6日</option>\n                        <option value=\"7\">7日</option>\n                        <option value=\"8\">8日</option>\n                        <option value=\"9\">9日</option>\n                        <option value=\"10\">10日</option>\n                        <option value=\"11\">11日</option>\n                        <option value=\"12\">12日</option>\n                        <option value=\"13\">13日</option>\n                        <option value=\"14\">14日</option>\n                        <option value=\"15\">15日</option>\n                        <option value=\"16\">16日</option>\n                        <option value=\"17\">17日</option>\n                        <option value=\"18\">18日</option>\n                        <option value=\"19\">19日</option>\n                        <option value=\"20\">20日</option>\n                        <option value=\"21\">21日</option>\n                        <option value=\"22\">22日</option>\n                        <option value=\"23\">23日</option>\n                        <option value=\"24\">24日</option>\n                        <option value=\"25\">25日</option>\n                        <option value=\"26\">26日</option>\n                        <option value=\"27\">27日</option>\n                        <option value=\"28\">28日</option>\n                        <option value=\"29\">29日</option>\n                        <option value=\"30\">30日</option>\n                        <option value=\"30\">31日</option>\n                    </select>\n                    <span>的</span>\n                    <select id=\"month_hour\">\n                        <option value=\"0\">0时</option>\n                        <option value=\"1\">1时</option>\n                        <option value=\"2\">2时</option>\n                        <option value=\"3\">3时</option>\n                        <option value=\"4\">4时</option>\n                        <option value=\"5\">5时</option>\n                        <option value=\"6\">6时</option>\n                        <option value=\"7\">7时</option>\n                        <option value=\"8\">8时</option>\n                        <option value=\"9\">9时</option>\n                        <option value=\"10\">10时</option>\n                        <option value=\"11\">11时</option>\n                        <option value=\"12\">12时</option>\n                        <option value=\"13\">13时</option>\n                        <option value=\"14\">14时</option>\n                        <option value=\"15\">15时</option>\n                        <option value=\"16\">16时</option>\n                        <option value=\"17\">17时</option>\n                        <option value=\"18\">18时</option>\n                        <option value=\"19\">19时</option>\n                        <option value=\"20\">20时</option>\n                        <option value=\"21\">21时</option>\n                        <option value=\"22\">22时</option>\n                        <option value=\"23\">23时</option>\n                    </select>\n                    <span>:</span>\n                    <select id=\"month_minute\">\n                        <option value=\"0\">0分</option>\n                        <option value=\"1\">1分</option>\n                        <option value=\"2\">2分</option>\n                        <option value=\"3\">3分</option>\n                        <option value=\"4\">4分</option>\n                        <option value=\"5\">5分</option>\n                        <option value=\"6\">6分</option>\n                        <option value=\"7\">7分</option>\n                        <option value=\"8\">8分</option>\n                        <option value=\"9\">9分</option>\n                        <option value=\"10\">10分</option>\n                        <option value=\"11\">11分</option>\n                        <option value=\"12\">12分</option>\n                        <option value=\"13\">13分</option>\n                        <option value=\"14\">14分</option>\n                        <option value=\"15\">15分</option>\n                        <option value=\"16\">16分</option>\n                        <option value=\"17\">17分</option>\n                        <option value=\"18\">18分</option>\n                        <option value=\"19\">19分</option>\n                        <option value=\"20\">20分</option>\n                        <option value=\"21\">21分</option>\n                        <option value=\"22\">22分</option>\n                        <option value=\"23\">23分</option>\n                        <option value=\"24\">24分</option>\n                        <option value=\"25\">25分</option>\n                        <option value=\"26\">26分</option>\n                        <option value=\"27\">27分</option>\n                        <option value=\"28\">28分</option>\n                        <option value=\"29\">29分</option>\n                        <option value=\"30\">30分</option>\n                        <option value=\"31\">31分</option>\n                        <option value=\"32\">32分</option>\n                        <option value=\"33\">33分</option>\n                        <option value=\"34\">34分</option>\n                        <option value=\"35\">35分</option>\n                        <option value=\"36\">36分</option>\n                        <option value=\"37\">37分</option>\n                        <option value=\"38\">38分</option>\n                        <option value=\"39\">39分</option>\n                        <option value=\"40\">40分</option>\n                        <option value=\"41\">41分</option>\n                        <option value=\"42\">42分</option>\n                        <option value=\"43\">43分</option>\n                        <option value=\"44\">44分</option>\n                        <option value=\"45\">45分</option>\n                        <option value=\"46\">46分</option>\n                        <option value=\"47\">47分</option>\n                        <option value=\"48\">48分</option>\n                        <option value=\"49\">49分</option>\n                        <option value=\"50\">50分</option>\n                        <option value=\"51\">51分</option>\n                        <option value=\"52\">52分</option>\n                        <option value=\"53\">53分</option>\n                        <option value=\"54\">54分</option>\n                        <option value=\"55\">55分</option>\n                        <option value=\"56\">56分</option>\n                        <option value=\"57\">57分</option>\n                        <option value=\"58\">58分</option>\n                        <option value=\"59\">59分</option>\n                    </select>\n                </div>\n                <div role=\"tabpanel\" class=\"tab-pane\" id=\"week\">\n                    <span>每周:</span>\n                    <select id=\"week_weekDay\">\n                        <option value=\"SUN\">星期日</option>\n                        <option value=\"MON\">星期一</option>\n                        <option value=\"TUE\">星期二</option>\n                        <option value=\"WED\">星期三</option>\n                        <option value=\"THU\">星期四</option>\n                        <option value=\"FIR\">星期五</option>\n                        <option value=\"SAT\">星期六</option>\n                    </select>\n                    <span>的</span>\n                    <select id=\"week_hour\">\n                        <option value=\"0\">0时</option>\n                        <option value=\"1\">1时</option>\n                        <option value=\"2\">2时</option>\n                        <option value=\"3\">3时</option>\n                        <option value=\"4\">4时</option>\n                        <option value=\"5\">5时</option>\n                        <option value=\"6\">6时</option>\n                        <option value=\"7\">7时</option>\n                        <option value=\"8\">8时</option>\n                        <option value=\"9\">9时</option>\n                        <option value=\"10\">10时</option>\n                        <option value=\"11\">11时</option>\n                        <option value=\"12\">12时</option>\n                        <option value=\"13\">13时</option>\n                        <option value=\"14\">14时</option>\n                        <option value=\"15\">15时</option>\n                        <option value=\"16\">16时</option>\n                        <option value=\"17\">17时</option>\n                        <option value=\"18\">18时</option>\n                        <option value=\"19\">19时</option>\n                        <option value=\"20\">20时</option>\n                        <option value=\"21\">21时</option>\n                        <option value=\"22\">22时</option>\n                        <option value=\"23\">23时</option>\n                    </select>\n                    <span>:</span>\n                    <select id=\"week_minute\">\n                        <option value=\"0\">0分</option>\n                        <option value=\"1\">1分</option>\n                        <option value=\"2\">2分</option>\n                        <option value=\"3\">3分</option>\n                        <option value=\"4\">4分</option>\n                        <option value=\"5\">5分</option>\n                        <option value=\"6\">6分</option>\n                        <option value=\"7\">7分</option>\n                        <option value=\"8\">8分</option>\n                        <option value=\"9\">9分</option>\n                        <option value=\"10\">10分</option>\n                        <option value=\"11\">11分</option>\n                        <option value=\"12\">12分</option>\n                        <option value=\"13\">13分</option>\n                        <option value=\"14\">14分</option>\n                        <option value=\"15\">15分</option>\n                        <option value=\"16\">16分</option>\n                        <option value=\"17\">17分</option>\n                        <option value=\"18\">18分</option>\n                        <option value=\"19\">19分</option>\n                        <option value=\"20\">20分</option>\n                        <option value=\"21\">21分</option>\n                        <option value=\"22\">22分</option>\n                        <option value=\"23\">23分</option>\n                        <option value=\"24\">24分</option>\n                        <option value=\"25\">25分</option>\n                        <option value=\"26\">26分</option>\n                        <option value=\"27\">27分</option>\n                        <option value=\"28\">28分</option>\n                        <option value=\"29\">29分</option>\n                        <option value=\"30\">30分</option>\n                        <option value=\"31\">31分</option>\n                        <option value=\"32\">32分</option>\n                        <option value=\"33\">33分</option>\n                        <option value=\"34\">34分</option>\n                        <option value=\"35\">35分</option>\n                        <option value=\"36\">36分</option>\n                        <option value=\"37\">37分</option>\n                        <option value=\"38\">38分</option>\n                        <option value=\"39\">39分</option>\n                        <option value=\"40\">40分</option>\n                        <option value=\"41\">41分</option>\n                        <option value=\"42\">42分</option>\n                        <option value=\"43\">43分</option>\n                        <option value=\"44\">44分</option>\n                        <option value=\"45\">45分</option>\n                        <option value=\"46\">46分</option>\n                        <option value=\"47\">47分</option>\n                        <option value=\"48\">48分</option>\n                        <option value=\"49\">49分</option>\n                        <option value=\"50\">50分</option>\n                        <option value=\"51\">51分</option>\n                        <option value=\"52\">52分</option>\n                        <option value=\"53\">53分</option>\n                        <option value=\"54\">54分</option>\n                        <option value=\"55\">55分</option>\n                        <option value=\"56\">56分</option>\n                        <option value=\"57\">57分</option>\n                        <option value=\"58\">58分</option>\n                        <option value=\"59\">59分</option>\n                    </select>\n\n                </div>\n                <div role=\"tabpanel\" class=\"tab-pane\" id=\"day\">\n                    <span>每天 的</span>\n                    <select id=\"day_hour\">\n                        <option value=\"0\">0时</option>\n                        <option value=\"1\">1时</option>\n                        <option value=\"2\">2时</option>\n                        <option value=\"3\">3时</option>\n                        <option value=\"4\">4时</option>\n                        <option value=\"5\">5时</option>\n                        <option value=\"6\">6时</option>\n                        <option value=\"7\">7时</option>\n                        <option value=\"8\">8时</option>\n                        <option value=\"9\">9时</option>\n                        <option value=\"10\">10时</option>\n                        <option value=\"11\">11时</option>\n                        <option value=\"12\">12时</option>\n                        <option value=\"13\">13时</option>\n                        <option value=\"14\">14时</option>\n                        <option value=\"15\">15时</option>\n                        <option value=\"16\">16时</option>\n                        <option value=\"17\">17时</option>\n                        <option value=\"18\">18时</option>\n                        <option value=\"19\">19时</option>\n                        <option value=\"20\">20时</option>\n                        <option value=\"21\">21时</option>\n                        <option value=\"22\">22时</option>\n                        <option value=\"23\">23时</option>\n                    </select>\n                    <span>:</span>\n                    <select id=\"day_minute\">\n                        <option value=\"0\">0分</option>\n                        <option value=\"1\">1分</option>\n                        <option value=\"2\">2分</option>\n                        <option value=\"3\">3分</option>\n                        <option value=\"4\">4分</option>\n                        <option value=\"5\">5分</option>\n                        <option value=\"6\">6分</option>\n                        <option value=\"7\">7分</option>\n                        <option value=\"8\">8分</option>\n                        <option value=\"9\">9分</option>\n                        <option value=\"10\">10分</option>\n                        <option value=\"11\">11分</option>\n                        <option value=\"12\">12分</option>\n                        <option value=\"13\">13分</option>\n                        <option value=\"14\">14分</option>\n                        <option value=\"15\">15分</option>\n                        <option value=\"16\">16分</option>\n                        <option value=\"17\">17分</option>\n                        <option value=\"18\">18分</option>\n                        <option value=\"19\">19分</option>\n                        <option value=\"20\">20分</option>\n                        <option value=\"21\">21分</option>\n                        <option value=\"22\">22分</option>\n                        <option value=\"23\">23分</option>\n                        <option value=\"24\">24分</option>\n                        <option value=\"25\">25分</option>\n                        <option value=\"26\">26分</option>\n                        <option value=\"27\">27分</option>\n                        <option value=\"28\">28分</option>\n                        <option value=\"29\">29分</option>\n                        <option value=\"30\">30分</option>\n                        <option value=\"31\">31分</option>\n                        <option value=\"32\">32分</option>\n                        <option value=\"33\">33分</option>\n                        <option value=\"34\">34分</option>\n                        <option value=\"35\">35分</option>\n                        <option value=\"36\">36分</option>\n                        <option value=\"37\">37分</option>\n                        <option value=\"38\">38分</option>\n                        <option value=\"39\">39分</option>\n                        <option value=\"40\">40分</option>\n                        <option value=\"41\">41分</option>\n                        <option value=\"42\">42分</option>\n                        <option value=\"43\">43分</option>\n                        <option value=\"44\">44分</option>\n                        <option value=\"45\">45分</option>\n                        <option value=\"46\">46分</option>\n                        <option value=\"47\">47分</option>\n                        <option value=\"48\">48分</option>\n                        <option value=\"49\">49分</option>\n                        <option value=\"50\">50分</option>\n                        <option value=\"51\">51分</option>\n                        <option value=\"52\">52分</option>\n                        <option value=\"53\">53分</option>\n                        <option value=\"54\">54分</option>\n                        <option value=\"55\">55分</option>\n                        <option value=\"56\">56分</option>\n                        <option value=\"57\">57分</option>\n                        <option value=\"58\">58分</option>\n                        <option value=\"59\">59分</option>\n                    </select>\n                </div>\n                <div role=\"tabpanel\" class=\"tab-pane\" id=\"hour\">\n                    <span>每隔</span>\n                    <select id=\"hour_hour\">\n                        <option value=\"1\">1时</option>\n                        <option value=\"2\">2时</option>\n                        <option value=\"3\">3时</option>\n                        <option value=\"4\">4时</option>\n                        <option value=\"5\">5时</option>\n                        <option value=\"6\">6时</option>\n                        <option value=\"7\">7时</option>\n                        <option value=\"8\">8时</option>\n                        <option value=\"9\">9时</option>\n                        <option value=\"10\">10时</option>\n                        <option value=\"11\">11时</option>\n                        <option value=\"12\">12时</option>\n                        <option value=\"13\">13时</option>\n                        <option value=\"14\">14时</option>\n                        <option value=\"15\">15时</option>\n                        <option value=\"16\">16时</option>\n                        <option value=\"17\">17时</option>\n                        <option value=\"18\">18时</option>\n                        <option value=\"19\">19时</option>\n                        <option value=\"20\">20时</option>\n                        <option value=\"21\">21时</option>\n                        <option value=\"22\">22时</option>\n                        <option value=\"23\">23时</option>\n                    </select>\n                    <span>小时</span>\n                </div>\n            </div>\n        </div>\n\n\n    </div>\n</div>\n<div class=\"control-group\">\n    <div class=\"controls\">\n        <input type=\"text\" readonly name=\"{{.name}}\" class=\"{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\n               {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\"\n               {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}{{ if .value}} value=\"{{.value}}\" {{end}}>\n        <div></div>\n        {{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}{{if .errors}}<ul>{{range .errors }}\n        <li>{{.}}</li>\n        {{end}}\n    </ul>{{end}}</span>{{end}}\n    </div>\n</div>\n\n\n{{end}}"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/date.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/datetime.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/datetime/time.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/generic.tmpl`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"generic\"}}<div class=\"form-group{{if .errors}} has-error{{end}}\">\n{{ if .label }}<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\n<input type=\"{{.type}}\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\n{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\n</div>\n{{end}}"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/hidden.html`,
		FileModTime: time.Unix(1491998145, 0),
		Content:     string("{{- define \"main\" -}}\n<input type=\"hidden\" name=\"{{.name}}\" class=\"{{range .classes}}{{.}} {{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{ if .value}} value=\"{{.value}}\"{{end}}>\n{{- end -}}"),
	}
	filec := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/input.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{ define \"main\"}}{{ template \"generic\" . }}{{ end }}"),
	}
	filed := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/mapinput.html`,
		FileModTime: time.Unix(1493283111, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{ if .label }}\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\n  {{- if or .helptext .errors }}\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n    {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}\n  </span>\n  {{- end}}\n  </div>\n</div>\n{{end}}"),
	}
	filef := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/number.html`,
		FileModTime: time.Unix(1490698434, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{if .label }}\n  <label\n      class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{if .id}}\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\n           {{if .id}} id=\"{{.id}}\" {{end}}\n           {{if .params}}\n           {{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\n           {{end}}\n           {{if .css}}\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\n           {{end}}\n           {{range $v :=.tags}} {{$v}}\n           {{end}}\n           {{if .value}}\n           value=\"{{.value}}\"\n           {{end}}>\n    {{if or .helptext .errors }}\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}\n              </ul>{{end}}\n    </span>\n    {{end}}\n  </div>\n</div>\n{{end}}"),
	}
	fileg := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/number/range.html`,
		FileModTime: time.Unix(1490692390, 0),
		Content:     string("{{ define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{if .label }}\n  <label\n      class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{if .id}}\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <input type=\"number\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\n           {{if .id}} id=\"{{.id}}\" {{end}}\n           {{if .params}}\n           {{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}\n           {{end}}\n           {{if .css}}\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\n           {{end}}\n           {{range $v :=.tags}} {{$v}}\n           {{end}}\n           {{if .value}}\n           value=\"{{.value}}\"\n           {{end}}>\n    {{if or .helptext .errors }}\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}\n              </ul>{{end}}\n    </span>\n    {{end}}\n  </div>\n</div>\n{{end}}"),
	}
	filei := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/checkbox.html`,
		FileModTime: time.Unix(1490760963, 0),
		Content:     string("{{ define \"main\"}}\n{{ $p := . }}\n<div class=\"form-group\">\n\t<div class=\"col-lg-offset-{{default .labelWidth 2}} col-lg-{{default .controlWidth 9}}\">\n\t\t<div class=\"checkbox{{if .errors}} has-error{{end}}\">\n\t\t\t<label class=\"control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\">\n\t\t\t\t<input type=\"checkbox\" name=\"{{.name}}\"{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>\n\t\t\t\t{{.label}}\n\t\t\t</label>\n\t\t\t{{if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end}}\n\t\t\t{{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end}}\n\t\t</div>\n\t</div>\n</div>\n{{end}}\n"),
	}
	filej := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/radiobutton.html`,
		FileModTime: time.Unix(1490698951, 0),
		Content:     string("{{ define \"main\"}}{{ $p := . }}\n<div class=\"form-group\">\n{{if .label}}<label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">{{.label}}</label>{{end}}\n<!-- <div class=\"radio\"> -->\n\t<div class=\"col-lg-{{default .controlWidth 9}}\">\n    {{ range .choices }}\n    <label class=\"control-label {{ if $p.labelClasses }}{{range $p.labelClasses}} {{.}}{{end}}{{end}}\">\n      <input type=\"radio\" name=\"{{$p.name}}\"{{ if $p.classes }} class=\"{{range $p.classes}}{{.}} {{end}}\"{{end}} value=\"{{.Value}}\" id=\"{{$p.id}}\"{{if $p.params}}{{range $k2, $v2 := $p.params}} {{$k2}}=\"{{$v2}}\"{{end}}{{end}}{{if $p.css}} style=\"{{range $k2, $v2 := .css}}{{$k2}}: {{$v2}}; {{end}}\"{{end}}{{ if eq $p.value .Value}} checked{{end}}{{range $v2 := $p.tags}} {{$v2}}{{end}}>\n      {{.Label}}\n    </label>\n    {{end}}\n\t</div>\n<!-- </div> -->\n\n</div>{{end}}\n"),
	}
	filek := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/options/select.html`,
		FileModTime: time.Unix(1493079556, 0),
		Content:     string("{{- define \"main\"}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{- if .label}}\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n      {{- if .id}} for=\"{{.id}}\" {{end -}}>\n    {{.label}}\n  </label>\n  {{- end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <select name=\"{{.name}}\" class=\"form-control {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"\n            {{- if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end -}}\n            {{- if .css}}\n            style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\n            {{- end -}}\n            {{- range $v :=.tags}} {{$v}}{{end}}>\n      {{- $p := . -}}\n      {{- range $k, $v := .choices -}}\n      {{- if $k -}}\n        <optgroup label=\"{{$k}}\">{{end -}}\n      {{- range $v -}}\n        <option value=\"{{.Value}}\"\n                {{- if strIn \"multiple\" $p.tags -}}\n                  {{- $id :=.Value -}}\n                  {{- range $k2, $p2 :=$p.multValues -}}\n                  {{- if eq $k2 $id}}selected{{end -}}\n                  {{- end -}}\n                {{- else -}}\n                  {{- if eq $p.value .Value}} selected{{end -}}\n                {{- end}}>{{raw .Label -}}\n        </option>\n      {{- end -}}\n      {{- if $k -}}\n        </optgroup>\n      {{- end -}}\n      {{- end -}}\n    </select>\n    {{- if or .helptext .errors }}<span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end}}</span>{{end -}}\n  </div>\n</div>\n{{- end}}\n"),
	}
	filel := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/static.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("{{define \"main\"}}<div class=\"form-group\">\n{{ if .label }}<label{{ if .labelClasses }} class=\"{{range .labelClasses}}{{.}} {{end}}\"{{end}}{{if .id}} for=\"{{.id}}\"{{end}}>{{.label}}</label>{{end}}\n<p name=\"{{.name}}\" class=\"form-control-static {{ if .classes }}{{range .classes}}{{.}} {{end}}{{end}}\"{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}}{{range $v := .tags}} {{$v}}{{end}}>{{.text}}</p>\n</div>{{end}}"),
	}
	filen := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/passwordinput.html`,
		FileModTime: time.Unix(1491998145, 0),
		Content:     string("{{- define \"main\" -}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{- if .label -}}\n  <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\" {{if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{- end -}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <input type=\"password\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}{{if .value}} value=\"{{.value}}\" {{end}}>\n    {{- if or .helptext .errors -}}\n    <span class=\"help-block\">\n      {{- if .helptext}}{{ .helptext }}{{end -}}\n      {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\n    </span>\n    {{- end -}}\n  </div>\n</div>{{end -}}"),
	}
	fileo := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textareainput.html`,
		FileModTime: time.Unix(1491998145, 0),
		Content:     string("{{- define \"main\" -}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{- if .label -}}\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{- if .id}} for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{- end}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n  <textarea name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\" {{if .id}} id=\"{{.id}}\" {{end}}{{if .params}}{{range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\" {{end}}{{range $v :=.tags}} {{$v}}{{end}}>{{.text}}</textarea>\n  {{- if or .helptext .errors -}}\n  <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\n    {{- if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end}}</ul>{{end -}}\n  </span>\n  {{- end -}}\n  </div>\n</div>\n{{- end -}}"),
	}
	filep := &embedded.EmbeddedFile{
		Filename:    `bootstrap3/text/textinput.html`,
		FileModTime: time.Unix(1493282778, 0),
		Content:     string("{{- define \"main\" -}}\n<div class=\"form-group{{if .errors}} has-error{{end}}\">\n  {{- if .label -}}\n    <label class=\"col-lg-{{default .labelWidth 2}} control-label {{ if .labelClasses }}{{range .labelClasses}} {{.}}{{end}}{{end}}\"\n  {{- if .id -}}\n      for=\"{{.id}}\" {{end}}>{{.label}}</label>\n  {{- end -}}\n  <div class=\"col-lg-{{default .controlWidth 9}}\">\n    <input type=\"text\" name=\"{{.name}}\" class=\"form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}\"\n           {{- if .id}} id=\"{{.id}}\" {{end -}}\n           {{- if .params -}}\n           {{- range $k, $v :=.params}} {{$k}}=\"{{$v}}\" {{end -}}\n           {{- end -}}\n           {{- if .css -}}\n           style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"\n           {{- end -}}\n           {{- range $v :=.tags}} {{$v -}}\n           {{- end -}}\n           {{- if .value -}}\n           value=\"{{.value}}\"\n           {{end}}>\n    {{- if or .helptext .errors -}}\n    <span class=\"help-block\">{{if .helptext}}{{ .helptext }}{{end -}}\n            {{if .errors}}<ul>{{ range .errors }}<li>{{.}}</li>{{end -}}\n              </ul>{{end -}}\n    </span>\n    {{end -}}\n  </div>\n</div>\n{{- end -}}\n"),
	}
	fileq := &embedded.EmbeddedFile{
		Filename:    `bootstrapform.html`,
		FileModTime: time.Unix(1489735114, 0),
		Content:     string("<form role=\"form\"{{if .name}} name=\"{{.name}}\"{{end}}{{ if .classes }} class=\"{{range .classes}}{{.}} {{end}}\"{{end}}{{if .id}} id=\"{{.id}}\"{{end}}{{if .params}}{{range $k, $v := .params}} {{$k}}=\"{{$v}}\"{{end}}{{end}}{{if .css}} style=\"{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}\"{{end}} method=\"{{.method}}\" action=\"{{.action}}\">\n\t{{ range .fields}}{{ .Render}}{{end}}\n</form>"),
	}
	filer := &embedded.EmbeddedFile{
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
			fileq, // bootstrapform.html
			filer, // fieldset.html

		},
	}
	dir3 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3`,
		DirModTime: time.Unix(1493283754, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file4, // bootstrap3/button.html
			file5, // bootstrap3/cron.html
			filea, // bootstrap3/generic.tmpl
			fileb, // bootstrap3/hidden.html
			filec, // bootstrap3/input.html
			filed, // bootstrap3/mapinput.html
			filel, // bootstrap3/static.html

		},
	}
	dir6 := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/datetime`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file7, // bootstrap3/datetime/date.html
			file8, // bootstrap3/datetime/datetime.html
			file9, // bootstrap3/datetime/time.html

		},
	}
	dire := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/number`,
		DirModTime: time.Unix(1489735114, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filef, // bootstrap3/number/number.html
			fileg, // bootstrap3/number/range.html

		},
	}
	dirh := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/options`,
		DirModTime: time.Unix(1491998145, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filei, // bootstrap3/options/checkbox.html
			filej, // bootstrap3/options/radiobutton.html
			filek, // bootstrap3/options/select.html

		},
	}
	dirm := &embedded.EmbeddedDir{
		Filename:   `bootstrap3/text`,
		DirModTime: time.Unix(1491998145, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filen, // bootstrap3/text/passwordinput.html
			fileo, // bootstrap3/text/textareainput.html
			filep, // bootstrap3/text/textinput.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir3, // bootstrap3

	}
	dir3.ChildDirs = []*embedded.EmbeddedDir{
		dir6, // bootstrap3/datetime
		dire, // bootstrap3/number
		dirh, // bootstrap3/options
		dirm, // bootstrap3/text

	}
	dir6.ChildDirs = []*embedded.EmbeddedDir{}
	dire.ChildDirs = []*embedded.EmbeddedDir{}
	dirh.ChildDirs = []*embedded.EmbeddedDir{}
	dirm.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1489735114, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":                    dir1,
			"bootstrap3":          dir3,
			"bootstrap3/datetime": dir6,
			"bootstrap3/number":   dire,
			"bootstrap3/options":  dirh,
			"bootstrap3/text":     dirm,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"baseform.html":                       file2,
			"bootstrap3/button.html":              file4,
			"bootstrap3/cron.html":                file5,
			"bootstrap3/datetime/date.html":       file7,
			"bootstrap3/datetime/datetime.html":   file8,
			"bootstrap3/datetime/time.html":       file9,
			"bootstrap3/generic.tmpl":             filea,
			"bootstrap3/hidden.html":              fileb,
			"bootstrap3/input.html":               filec,
			"bootstrap3/mapinput.html":            filed,
			"bootstrap3/number/number.html":       filef,
			"bootstrap3/number/range.html":        fileg,
			"bootstrap3/options/checkbox.html":    filei,
			"bootstrap3/options/radiobutton.html": filej,
			"bootstrap3/options/select.html":      filek,
			"bootstrap3/static.html":              filel,
			"bootstrap3/text/passwordinput.html":  filen,
			"bootstrap3/text/textareainput.html":  fileo,
			"bootstrap3/text/textinput.html":      filep,
			"bootstrapform.html":                  fileq,
			"fieldset.html":                       filer,
		},
	})
}
