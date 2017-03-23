package forms

// CronField creates a default text input field based on the provided name.
func CronField(ctx interface{}, name, label string) *Field {
	return FieldWithTypeWithCtx(ctx, name, label, CRON)
}
