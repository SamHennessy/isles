package isles

import "github.com/SamHennessy/hlive"

func Navbar(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class("navbar "+class))...)
}

func NavbarStart(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class("navbar-start "+class))...)

}

func NavbarCenter(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class("navbar-center "+class))...)

}

func NavbarEnd(class string, elements ...any) *hlive.ComponentMountable {
	return hlive.NewComponentMountable("div", append(elements, hlive.Class("navbar-end "+class))...)

}
