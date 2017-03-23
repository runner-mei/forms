package forms

import (
	"testing"
	"time"
)

const (
	style = BOOTSTRAP
)

var (
	txt, psw, btn FieldInterface
)

func makeCtx() interface{} {
	return nil
}

func TestFieldRender(t *testing.T) {
	field := TextField(makeCtx(), "test", "test")
	field.AddClass("test").AddClass("class").SetID("testId").SetParam("param1", "val1").AddCss("css1", "val1").SetStyle(style).Disabled()
	field.AddLabelClass("LABEL")
	field.SetLabel("This is a label")
	field.AddError("ERROR")
	t.Log("Rendered field:", field.Render())
	txt = field
}

func TestPasswordRender(t *testing.T) {
	field := PasswordField(makeCtx(), "test", "test")
	field.AddClass("test")
	field.AddClass("class")
	field.SetID("testId")
	field.SetStyle(style)
	field.SetValue("asd")
	t.Log("Rendered field:", field.Render())
	psw = field
}

func TestButtonRender(t *testing.T) {
	field := SubmitButton(makeCtx(), "btn", "Click me!")
	field.SetStyle(style)
	t.Log("Rendered button:", field.Render())
	btn = field
}

func TestRadioButtonRender(t *testing.T) {
	field := RadioField(makeCtx(), "radio", "select one", []InputChoice{
		InputChoice{ID: "choice1", Val: "value1"},
		InputChoice{ID: "choice2", Val: "value2"},
	})
	field.SetStyle(style)
	t.Log("Rendered radio:", field.Render())
}

func TestSelectRender(t *testing.T) {
	field := SelectField(makeCtx(), "select", "select one", map[string][]InputChoice{
		"": []InputChoice{
			InputChoice{"choice1", "value1"},
			InputChoice{"choice2", "value2"},
		},
	}).MultipleChoice().SetLabel("asd").AddSelected("choice1", "choice2")
	field.SetStyle(style).SetValue("choice1")
	t.Log("Rendered select:", field.Render())
}

func TestHiddenRender(t *testing.T) {
	field := HiddenField(makeCtx(), "hidden")
	field.SetStyle(style)
	t.Log("Rendered hidden:", field.Render())
}

func TestNumberRender(t *testing.T) {
	field := NumberField(makeCtx(), "number", "input number")
	field.SetStyle(style)
	t.Log("Rendered number:", field.Render())
}

func TestFormRender(t *testing.T) {
	form := BootstrapForm(POST, "")
	form.Elements(&FieldSetType{}, txt, psw, btn)
	// form.AddField(psw)
	// form.AddField(btn)
	t.Log("Rendered form:", form.Render())
}

func TestFormFromSimpleModel(t *testing.T) {
	type User struct {
		Username  string
		Password1 string `form_widget:"password" form_label:"Password 1"`
		Password2 string `form_widget:"password" form_label:"Password 2"`
		SkipThis  int    `form_options:"skip"`
	}

	u := User{}

	form := BaseFormFromModel(u, POST, "/action.html")
	t.Log("Rendered form:", form.Render())
}

func TestFormFromModel(t *testing.T) {
	type Model struct {
		User      string    `form_label:"User label test"`
		password  string    `form_widget:"password"`
		ID        int       `form_min:"0" form_max:"5"`
		Ts        time.Time `form_min:"2013-04-22T15:00"`
		RadioTest string    `form_widget:"select" form_choices:"|A|Option A|G1|B|Option B" form_value:"A"`
		BoolTest  bool      //`form_options:"checked"`
	}

	form := BaseFormFromModel(Model{"asd", "lol", 20, time.Now(), "B", false}, POST, "")
	t.Log("Rendered form:", form.Render())
}

func TestBSFormFromModel(t *testing.T) {
	type Model struct {
		User      string    `form_label:"User label test"`
		password  string    `form_widget:"password"`
		ID        int       `form_min:"0" form_max:"5"`
		Ts        time.Time `form_min:"2013-04-22T15:00"`
		RadioTest string    `form_widget:"select" form_choices:"|A|Option A|G1|B|Option B" form_value:"A"`
		BoolTest  bool      //`form_options:"checked"`
	}

	form := BootstrapFormFromModel(Model{"asd", "lol", 20, time.Now(), "B", false}, POST, "")
	t.Log("Rendered form:", form.Render())
}

func TestInlineCreation(t *testing.T) {
	form := BaseForm(POST, "/action.html").Elements(
		TextField(makeCtx(), "text_field", "Username"),
		FieldSet("psw_fieldset",
			PasswordField(makeCtx(), "psw1", "Password 1").AddClass("password_class"),
			PasswordField(makeCtx(), "psw2", "Password 2").AddClass("password_class"),
		),
		SubmitButton(makeCtx(), "btn1", "Submit"),
	)
	t.Log("Rendered form:", form.Render())
}

func TestPizzaCreation(t *testing.T) {

	type Ingredient struct {
		IngredientID int    `db_key:"true" db_autoincrement:"true"`
		Name         string `db_size:"30"`
	}

	type Pizza struct {
		ID   int    `db_key:"true" db_autoincrement:"true"`
		Name string `db_size:"30"`
		// Ingredients []int
		Price float32 `form_widget:"number" form_min:"0"`

		// Transient
		Ingrs []Ingredient `db_transient:"true" form_options:"skip"`
	}

	form := BootstrapFormFromModel(Pizza{Price: 2.2}, POST, "")
	t.Log("Rendered form:", form.Render())
}
