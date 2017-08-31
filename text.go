package forms

// TextField creates a default text input field based on the provided name.
func TextField(ctx interface{}, name, label string) *Field {
	return FieldWithTypeWithCtx(ctx, name, label, TEXT)
}

// PasswordField creates a default password text input field based on the provided name.
func PasswordField(ctx interface{}, name, label string) *Field {
	return FieldWithTypeWithCtx(ctx, name, label, PASSWORD)
}

// =========== TEXT AREA

// TextAreaField creates a default textarea input field based on the provided name and dimensions.
func TextAreaField(ctx interface{}, name, label string, rows, cols int) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, TEXTAREA)
	if rows > 0 {
		ret.SetIntParam("rows", rows)
	}
	if cols > 0 {
		ret.SetIntParam("cols", cols)
	}
	return ret
}

// MapField creates a default textarea input field based on the provided name and dimensions.
func MapField(ctx interface{}, name, label string, rows, cols int) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, MAP)
	if rows > 0 {
		ret.SetIntParam("rows", rows)
	}
	if cols > 0 {
		ret.SetIntParam("cols", cols)
	}
	return ret
}

// ========================

// HiddenField creates a default hidden input field based on the provided name.
func HiddenField(ctx interface{}, name string) *Field {
	return FieldWithTypeWithCtx(ctx, name, "", HIDDEN)
}
