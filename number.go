package forms

// RangeField creates a default range field with the provided name. Min, max and step parameters define the expected behavior
// of the HTML field.
func RangeField(ctx interface{}, name, label string, min, max, step int) *Field {
	ret := FieldWithTypeWithCtx(ctx, name, label, RANGE)
	ret.SetParam("min", string(min))
	ret.SetParam("max", string(max))
	ret.SetParam("step", string(step))
	return ret
}

// NumberField craetes a default number field with the provided name.
func NumberField(ctx interface{}, name, label string) *Field {
	return FieldWithTypeWithCtx(ctx, name, label, NUMBER)
}

func init() {
	FieldFuncs["range_field"] = RangeField
	FieldFuncs["number_field"] = NumberField
}
