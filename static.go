package forms

// StaticField returns a static field with the provided name and content
func StaticField(name, content string) *Field {
	ret := FieldWithType(nil, name, STATIC)
	ret.SetText(content)
	return ret
}
