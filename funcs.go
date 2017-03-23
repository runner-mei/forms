package forms

func init() {
	FieldFuncs["radio_field"] = RadioField
	FieldFuncs["select_field"] = SelectField
	FieldFuncs["checkbox_field"] = Checkbox
	FieldFuncs["datetime_field"] = DatetimeField
	FieldFuncs["date_field"] = DateField
	FieldFuncs["time_field"] = TimeField
	FieldFuncs["cron_field"] = CronField
	FieldFuncs["submit"] = SubmitButton
	FieldFuncs["reset"] = ResetButton
	FieldFuncs["button"] = Button
	FieldFuncs["text_field"] = TextField
	FieldFuncs["password_field"] = PasswordField
	FieldFuncs["textarea_field"] = TextAreaField
	FieldFuncs["hidden_field"] = HiddenField
	FieldFuncs["map_field"] = MapField
	FieldFuncs["static_field"] = StaticField
	FieldFuncs["range_field"] = RangeField
	FieldFuncs["number_field"] = NumberField
}
