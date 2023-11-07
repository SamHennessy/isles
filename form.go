package isles

import "github.com/SamHennessy/hlive"

func AttrDisabled() *hlive.Attribute {
	return hlive.NewAttribute("disabled", "")
}

func AttrSelected() *hlive.Attribute {
	return hlive.NewAttribute("selected", "")
}

func FormControl(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class("form-control "+class))...)
}

func Label(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("label", append(elements, hlive.Class("label "+class))...)
}

func LabelText(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("span", append(elements, hlive.Class("label-text "+class))...)
}

func LabelTextAlt(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("span", append(elements, hlive.Class("label-text-alt "+class))...)
}

func Select(class string, onChanage hlive.EventHandler, elements ...any) *hlive.ComponentMountable {
	c := hlive.NewComponentMountable("select", append(elements, hlive.Class("select "+class))...)

	// TODO: fix hlive to allow nil
	if onChanage != nil {
		c.Add(hlive.On("change", onChanage))
	}

	return c
}

func Option(class string, value *hlive.LockBox[string], elements ...any) *hlive.ComponentMountable {
	c := hlive.NewComponentMountable("option", append(elements, hlive.Class(class))...)

	if value != nil {
		c.Add(hlive.AttrsLockBox{"value": value})
	}

	return c
}

func Button(class string, onClick hlive.EventHandler, elements ...any) *hlive.ComponentMountable {
	c := hlive.NewComponentMountable(
		"button",
		hlive.Attrs{"type": "submit"},
		hlive.Class("btn "+class), elements)

	if onClick != nil {
		c.Add(hlive.On("click", onClick))
	}

	return c
}

func InputText(class, placeholder string, onInput hlive.EventHandler, elements ...any) *hlive.ComponentMountable {
	c := hlive.NewComponentMountable(
		"input",
		hlive.Attrs{"type": "text", "placeholder": placeholder},
		hlive.Class("input "+class),
		elements,
	)

	if onInput != nil {
		c.Add(hlive.On("input", onInput))
	}

	return c
}

func Form(class string, onSubmit hlive.EventHandler, elements ...any) *hlive.ComponentMountable {
	c := hlive.NewComponentMountable(
		"form",
		append(elements, hlive.PreventDefault(), hlive.Class(class))...,
	)

	if onSubmit != nil {
		c.Add(hlive.On("submit", onSubmit))
	}

	return c
}

func Join(elements ...any) *hlive.Tag {
	return hlive.NewTag("div", hlive.Class("join"), elements)
}
