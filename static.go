package forms

// StaticField returns a static field with the provided name and content
func StaticField(name, content string) *Field {
	ret := FieldWithType(name, STATIC)
	ret.SetText(content)
	return ret
}

func init() {
	FieldFuncs["static_field"] = StaticField
}
