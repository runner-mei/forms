package forms

// SubmitButton creates a default button with the provided name and text.
func SubmitButton(ctx interface{}, name string, text string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, "", SUBMIT)
	ret.SetText(text)
	return ret
}

// ResetButton creates a default reset button with the provided name and text.
func ResetButton(ctx interface{}, name string, text string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, "", RESET)
	ret.SetText(text)
	return ret
}

// Button creates a default generic button
func Button(ctx interface{}, name string, text string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, "", BUTTON)
	ret.SetText(text)
	return ret
}
