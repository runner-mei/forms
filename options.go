package forms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"strings"
)

// InputChoice - Value pair used to define an option for select and redio input fields.
type InputChoice struct {
	Value string `json:"value"`
	Label string `json:"label"`
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
func readChoices(name string, v interface{}) []InputChoice {
	if v == nil {
		return []InputChoice{}
	}

	if strArray, ok := v.([]string); ok {
		choices := []InputChoice{}
		for _, s := range strArray {
			choices = append(choices, InputChoice{s, s})
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
		return choices
	}

	if strMap, ok := v.(map[string]interface{}); ok {
		choices := []InputChoice{}
		for k, v := range strMap {
			choices = append(choices, InputChoice{k, fmt.Sprint(v)})
		}
		return choices
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

func readChoiceGroups(name string, v interface{}) []HierarchyChoice {
	if v == nil {
		return []HierarchyChoice{}
	}

	if strArray, ok := v.([]string); ok {
		choices := []InputChoice{}
		for _, s := range strArray {
			choices = append(choices, InputChoice{s, s})
		}
		return []HierarchyChoice{{Children: choices}}
	}

	if strArray, ok := v.([][2]string); ok {
		choices := []InputChoice{}
		for _, sa := range strArray {
			choices = append(choices, InputChoice{sa[0], sa[1]})
		}

		return []HierarchyChoice{{Children: choices}}
	}

	if strMap, ok := v.(map[string]interface{}); ok {
		choices := []InputChoice{}
		for k, v := range strMap {
			choices = append(choices, InputChoice{k, fmt.Sprint(v)})
		}

		return []HierarchyChoice{{Children: choices}}
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
	if choices, ok := v.([]InputChoice); ok {
		return []HierarchyChoice{{Children: choices}}
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
