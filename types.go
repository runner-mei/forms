package forms

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Input field types
const (
	BUTTON         = "button"
	CHECKBOX       = "checkbox"
	COLOR          = "color" // Not yet implemented
	DATE           = "date"
	DATETIME       = "datetime"
	DATETIME_LOCAL = "datetime-local"
	EMAIL          = "email" // Not yet implemented
	FILE           = "file"  // Not yet implemented
	HIDDEN         = "hidden"
	IMAGE          = "image" // Not yet implemented
	MONTH          = "month" // Not yet implemented
	NUMBER         = "number"
	PASSWORD       = "password"
	RADIO          = "radio"
	RANGE          = "range"
	RESET          = "reset"
	SEARCH         = "search" // Not yet implemented
	SUBMIT         = "submit"
	TEL            = "tel" // Not yet implemented
	TEXT           = "text"
	TIME           = "time"
	URL            = "url"  // Not yet implemented
	WEEK           = "week" // Not yet implemented
	TEXTAREA       = "textarea"
	SELECT         = "select"
	STATIC         = "static"
	CRON           = "cron"
)

// Available form styles
const (
	BASE      = "base"
	BOOTSTRAP = "bootstrap3"
)

// CreateURL creates the complete url of the desired widget template
func CreateURL(widget string) string {
	if _, err := os.Stat(widget); os.IsNotExist(err) {
		for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
			widgetPath := path.Join(p, "mc/app/libs/forms", widget)
			if _, err := os.Stat(widgetPath); os.IsNotExist(err) {
				fmt.Println(widgetPath)
				continue
			} else {
				return widgetPath
			}
		}
	}
	return widget
}

type stringSet []string

func (set stringSet) add(value string) stringSet {
	for _, s := range set {
		if s == value {
			return set
		}
	}

	return append(set, value)
}

func (set stringSet) remove(value string) stringSet {
	offset := 0
	for idx, s := range set {
		if s == value {
			continue
		}
		if offset != idx {
			set[offset] = s
		}
		offset++
	}

	return set[:offset]
}

func (set stringSet) has(value string) bool {
	for _, s := range set {
		if s == value {
			return true
		}
	}

	return false
}
