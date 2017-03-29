package forms

import (
	"encoding/json"
	"errors"
	"fmt"

	"strings"
)

// InputChoice - Value pair used to define an option for select and redio input fields.
type InputChoice struct {
	ID  string `json:"id"`
	Val string `json:"val"`
}

// RadioField creates a default radio button input field with the provided name and list of choices.
func RadioField(ctx interface{}, name, label string, choices interface{}) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, RADIO)
	ret.additionalData["choices"] = []InputChoice{}
	ret.SetRadioChoices(readChoices(choices))
	return ret
}

// SelectField creates a default select input field with the provided name and map of choices. Choices for SelectField are grouped
// by name (if <optgroup> is needed); "" group is the default one and does not trigger a <optgroup></optgroup> rendering.
func SelectField(ctx interface{}, name, label string, choices interface{}) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, SELECT)
	ret.additionalData["choices"] = map[string][]InputChoice{}
	ret.additionalData["multValues"] = map[string]struct{}{}
	ret.SetSelectChoices(readChoiceGroups(choices))
	return ret
}

func readChoices(v interface{}) []InputChoice {
	if choices, ok := v.([]InputChoice); ok {
		return choices
	}
	if s, ok := v.(string); ok {
		if strings.HasPrefix(s, "[") {
			var results []InputChoice
			if err := json.Unmarshal([]byte(s), &results); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice, " + err.Error()))
			}
			return results
		}

		panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice."))
	}
	panic(fmt.Errorf("Choices arguments must is []InputChoice - [%T]%#V", v, v))
}

func readChoiceGroups(v interface{}) map[string][]InputChoice {
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
			return results
		}

		if strings.HasPrefix(s, "[") {
			var results []InputChoice
			if err := json.Unmarshal([]byte(s), &results); err != nil {
				panic(errors.New("failed to unmarshal `" + s + "` to []InputChoice, " + err.Error()))
			}
			return map[string][]InputChoice{"": results}
		}

		panic(errors.New("failed to unmarshal `" + s + "` to map[string][]InputChoice."))
	}
	panic(fmt.Errorf("Choices arguments must is map[string][]InputChoice or []InputChoice - [%T]%#V", v, v))
}

// Checkbox creates a default checkbox field with the provided name. It also makes it checked by default based
// on the checked parameter.
func Checkbox(ctx interface{}, name, label string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, CHECKBOX)
	fmt.Println("ret.value:", ret.value)
	if s := strings.ToUpper(ret.value); s == "TRUE" || s == "enabled" || s == "yes" || s == "1" || s == "ON" {
		ret.AddTag("checked")
	}
	return ret
}
