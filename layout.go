package isles

import "github.com/SamHennessy/hlive"

func Container(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class("container "+class))...)
}
