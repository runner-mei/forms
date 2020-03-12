package forms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"strings"
)

// InputChoice - Value pair used to define an option for select and redio input fields.
type InputChoice struct {
	Value string `json:"value" xorm:"value"`
	Label string `json:"label" xorm:"label"`
}

type HierarchyChoice struct {
	Label    string        `json:"label,omitempty"`
	Children []InputChoice `json:"children,omitempty"`
}

func AppendHierarchyChoices(allList, parts []HierarchyChoice) []HierarchyChoice {
	for pidx := range parts {
		foundIdx := -1
		for idx := range allList {
			if allList[idx].Label == parts[pidx].Label {
				foundIdx = idx
				break
			}
		}

		if foundIdx >= 0 {
			if len(parts[pidx].Children) > 0 {
				allList[foundIdx].Children = append(allList[foundIdx].Children, parts[pidx].Children...)
			}
		} else {
			allList = append(allList, parts[pidx])
		}
	}
	return allList
}

// RadioField creates a default radio button input field with the provided name and list of choices.
func RadioField(ctx interface{}, name, label string, choices interface{}) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, RADIO)
	ret.additionalData["choices"] = []InputChoice{}
	ret.SetRadioChoices(readChoices(name, choices))
	return ret
}

// SelectField creates a default select input field with the provided name and map of choices. Choices for SelectField are grouped
// by name (if <optgroup> is needed); "" group is the default one and does not trigger a <optgroup></optgroup> rendering.
func SelectField(ctx interface{}, name, label string, choices interface{}) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, SELECT)
	ret.additionalData["choices"] = map[string][]InputChoice{}
	ret.additionalData["multValues"] = map[string]struct{}{}
	ret.SetSelectChoices(readChoiceGroups(name, choices))
	return ret
}

// MultSourceSelectField creates a default select input field with the provided name and map of choices. Choices for SelectField are grouped
// by name (if <optgroup> is needed); "" group is the default one and does not trigger a <optgroup></optgroup> rendering.
func MultSourceSelectField(ctx interface{}, label string) *Field {
	return FieldWithTypeWithCtx(ctx, "", label, MULI_SOURCE_SELECT)
}

func (f *Field) AddSource(name, label string, hasNew bool, choices interface{}) *Field {
	var sources []map[string]interface{}
	o := f.additionalData["sources"]
	if o != nil {
		sources, _ = o.([]map[string]interface{})
	}

	sources = append(sources, map[string]interface{}{
		"name":    name,
		"label":   label,
		"choices": readChoices(name, choices),
		"hasNew":  hasNew,
	})
	f.additionalData["sources"] = sources
	return f
}

func validateChoices(results []InputChoice) {
	count := 0
	for idx := range results {
		if results[idx].Value == "" {
			count++
		}

		if results[idx].Label == "" {
			results[idx].Label = results[idx].Value
		}
	}

	if count > 1 {
		panic(errors.New("Value of InputChoice is empty"))
	}
}

func tryReadChoices(v interface{}) []InputChoice {
	if v == nil {
		return []InputChoice{}
	}

	if strArray, ok := v.([]string); ok {
		choices := []InputChoice{}
		for _, s := range strArray {
			vls := strings.Split(s, ";")
			if len(vls) == 2 {
				choices = append(choices, InputChoice{vls[1], vls[0]})
			} else {
				choices = append(choices, InputChoice{s, s})
			}
		}
		return choices
	}

	if strArray, ok := v.([][2]string); ok {
		choices := []InputChoice{}
		for _, sa := range strArray {
			choices = append(choices, InputChoice{sa[0], sa[1]})
		}
		return choices
	}

	if choices, ok := v.([]InputChoice); ok {
		if choices == nil {
			return []InputChoice{}
		}
		return choices
	}

	if strMap, ok := v.(map[string]interface{}); ok {
		choices := []InputChoice{}
		for k, v := range strMap {
			choices = append(choices, InputChoice{k, fmt.Sprint(v)})
		}
		return choices
	}

	if objectArray, ok := v.([]map[string]interface{}); ok {
		choices := []InputChoice{}
		for _, values := range objectArray {
			id := values["value"]
			if id == nil {
				id = values["id"]
			}
			label := values["label"]
			if label == nil {
				label = values["text"]
			}
			choices = append(choices, InputChoice{fmt.Sprint(id), fmt.Sprint(label)})
		}
		return choices
	}

	if strMap, ok := v.(map[int64]string); ok {
		choices := []InputChoice{}
		for k, v := range strMap {
			choices = append(choices, InputChoice{strconv.FormatInt(k, 10), v})
		}
		return choices
	}

	if strMap, ok := v.(map[int]string); ok {
		choices := []InputChoice{}
		for k, v := range strMap {
			choices = append(choices, InputChoice{strconv.Itoa(k), v})
		}
		return choices
	}
	return nil
}

func readChoices(name string, v interface{}) []InputChoice {
	opts := tryReadChoices(v)
	if opts != nil {
		return opts
	}

	if s, ok := v.(string); ok {
		if strings.HasPrefix(s, "[") {
			var results []InputChoice
			if err := json.Unmarshal([]byte(s), &results); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice, " + err.Error()))
			}
			validateChoices(results)
			return results
		}

		panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice."))
	}
	panic(fmt.Errorf("Choices arguments of "+name+" must be []InputChoice - [%T]%#v", v, v))
}

func ToChoices(name string, v interface{}) []InputChoice {
	return readChoices(name, v)
}

func readChoiceGroups(name string, v interface{}) []HierarchyChoice {
	opts := tryReadChoices(v)
	if opts != nil {
		if len(opts) == 0 {
			return []HierarchyChoice{}
		}
		return []HierarchyChoice{{Children: opts}}
	}

	if choices, ok := v.(map[string][]InputChoice); ok {
		var results []HierarchyChoice
		for k, v := range choices {
			results = append(results, HierarchyChoice{
				Label:    k,
				Children: v,
			})
		}
		return results
	}
	if choices, ok := v.([]HierarchyChoice); ok {
		return choices
	}
	if bs, ok := v.([]byte); ok {
		if bytes.HasPrefix(bs, []byte("{")) {
			var values map[string][]InputChoice
			if err := json.Unmarshal(bs, &values); err != nil {
				panic(errors.New("failed to unmarshal `" + string(bs) + "` to map[string][]InputChoice, " + err.Error()))
			}
			for _, choices := range values {
				validateChoices(choices)
			}

			var results []HierarchyChoice
			for k, choices := range values {
				results = append(results, HierarchyChoice{
					Label:    k,
					Children: choices,
				})
			}
			return results
		}

		if bytes.HasPrefix(bs, []byte("[")) {
			var choices []InputChoice
			if err := json.Unmarshal(bs, &choices); err == nil {
				validateChoices(choices)
				return []HierarchyChoice{{Children: choices}}
			}

			var hierarchyChoices []HierarchyChoice
			if err := json.Unmarshal(bs, &hierarchyChoices); err != nil {
				panic(errors.New("failed to unmarshal `" + string(bs) + "` to []InputChoice or []HierarchyChoice, " + err.Error()))
			}
			for idx := range hierarchyChoices {
				validateChoices(hierarchyChoices[idx].Children)
			}
			return hierarchyChoices
		}

		panic(errors.New("failed to unmarshal `" + string(bs) + "` to []InputChoice or []HierarchyChoice."))
	}
	if s, ok := v.(string); ok {
		if strings.HasPrefix(s, "{") {
			var values map[string][]InputChoice
			if err := json.Unmarshal([]byte(s), &values); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to map[string][]InputChoice, " + err.Error()))
			}
			for _, choices := range values {
				validateChoices(choices)
			}

			var results []HierarchyChoice
			for k, choices := range values {
				results = append(results, HierarchyChoice{
					Label:    k,
					Children: choices,
				})
			}
			return results
		}

		if strings.HasPrefix(s, "[") {
			var choices []InputChoice
			if err := json.Unmarshal([]byte(s), &choices); err == nil {
				validateChoices(choices)
				return []HierarchyChoice{{Children: choices}}
			}

			var hierarchyChoices []HierarchyChoice
			if err := json.Unmarshal([]byte(s), &hierarchyChoices); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice or []HierarchyChoice, " + err.Error()))
			}
			for idx := range hierarchyChoices {
				validateChoices(hierarchyChoices[idx].Children)
			}
			return hierarchyChoices
		}

		panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice or []HierarchyChoice."))
	}
	panic(fmt.Errorf("Choices argumentsof "+name+" must is map[string][]InputChoice or []InputChoice - [%T]%#v", v, v))
}

// Checkbox creates a default checkbox field with the provided name. It also makes it checked by default based
// on the checked parameter.
func Checkbox(ctx interface{}, name, label string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, CHECKBOX)
	if s := strings.ToUpper(ret.value); s == "TRUE" || s == "enabled" || s == "yes" || s == "1" || s == "ON" {
		ret.AddTag("checked")
	}
	return ret
}
