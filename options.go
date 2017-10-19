package forms

import (
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
	for idx := range results {
		if results[idx].Value == "" {
			panic(errors.New("Value of InputChoice is empty"))
		}

		if results[idx].Label == "" {
			results[idx].Label = results[idx].Value
		}
	}
}
func readChoices(name string, v interface{}) []InputChoice {
	if v == nil {
		return []InputChoice{}
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

func readChoiceGroups(name string, v interface{}) map[string][]InputChoice {
	if v == nil {
		return map[string][]InputChoice{}
	}
	if strMap, ok := v.(map[string]interface{}); ok {
		choices := []InputChoice{}
		for k, v := range strMap {
			choices = append(choices, InputChoice{k, fmt.Sprint(v)})
		}

		return map[string][]InputChoice{"": choices}
	}
	if choices, ok := v.(map[string][]InputChoice); ok {
		return choices
	}
	if choices, ok := v.([]InputChoice); ok {
		return map[string][]InputChoice{"": choices}
	}
	if s, ok := v.(string); ok {
		if strings.HasPrefix(s, "{") {
			var results map[string][]InputChoice
			if err := json.Unmarshal([]byte(s), &results); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to map[string][]InputChoice, " + err.Error()))
			}
			for _, choices := range results {
				validateChoices(choices)
			}
			return results
		}

		if strings.HasPrefix(s, "[") {
			var results []InputChoice
			if err := json.Unmarshal([]byte(s), &results); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice, " + err.Error()))
			}
			validateChoices(results)
			return map[string][]InputChoice{"": results}
		}

		panic(errors.New("failed to unmarshal `" + s + "` to map[string][]InputChoice."))
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
