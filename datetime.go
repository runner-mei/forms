package forms

// Datetime format string to convert from time.Time objects to HTML fields and viceversa.
const (
	DATETIME_FORMAT = "2006-01-02T15:05"
	DATE_FORMAT     = "2006-01-02"
	TIME_FORMAT     = "15:05"
)

// DatetimeField creates a default datetime input field with the given name.
func DatetimeField(ctx interface{}, name, label string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, DATETIME)
	return ret
}

// DateField creates a default date input field with the given name.
func DateField(ctx interface{}, name, label string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, DATE)
	return ret
}

// TimeField creates a default time input field with the given name.
func TimeField(ctx interface{}, name, label string) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, TIME)
	return ret
}
