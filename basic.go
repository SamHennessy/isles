package isles

import "github.com/SamHennessy/hlive"

func Div(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class(class))...)
}
